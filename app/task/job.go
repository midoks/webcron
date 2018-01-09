package task

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/midoks/webcron/app/models"
	"golang.org/x/crypto/ssh"
	// "io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Job struct {
	id         int                                               // 任务ID
	logId      int64                                             // 日志记录ID
	name       string                                            // 任务名称
	task       *models.AppCron                                   // 任务对象
	runFunc    func(time.Duration) (string, string, error, bool) // 执行函数
	status     int                                               // 任务状态，大于0表示正在执行中
	Concurrent bool                                              // 同一个任务是否允许并行执行
}

func NewJobFromTask(cron *models.AppCron) (*Job, error) {
	if cron.Id < 1 {
		return nil, fmt.Errorf("ToJob: 缺少id")
	}
	job := NewCommandJob(cron)
	job.task = cron
	job.Concurrent = cron.Concurrent == 1
	return job, nil
}

func IsWin() bool {
	// fmt.Println("sys:", runtime.GOOS)
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

func ConnectByUser(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func ConnectByRsa() {

}

func ServerCmd(item *models.AppItem) {
	fmt.Println("ServerCmd Item", item)

	server, _ := models.ServerGetById(item.ServerId)

	if server.Type == 0 {
		session, err := ConnectByUser(server.User, server.Pwd, server.Ip, server.Port)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()

		session.Stdout = os.Stdout
		session.Stderr = os.Stderr

		session.Run("ls")
	}
	fmt.Println("ServerCmd Server", server)
}

func NewCommandJob(cron *models.AppCron) *Job {
	job := &Job{
		id:   cron.Id,
		name: cron.Name,
	}
	job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
		bufOut := new(bytes.Buffer)
		bufErr := new(bytes.Buffer)

		item, _ := models.ItemGetById(cron.ItemId)

		var cmd *exec.Cmd

		if item.Type == 0 {
			ServerCmd(item)
		} else {

		}

		fmt.Println(cron.ItemId)

		if IsWin() {
			cmd = exec.Command("cmd", "/c", cron.Cmd)
		} else {
			cmd = exec.Command("/bin/bash", "-c", cron.Cmd)
		}

		cmd.Stdout = bufOut
		cmd.Stderr = bufErr
		cmd.Start()
		err, isTimeout := runCmdWithTimeout(cmd, timeout)

		fmt.Println(bufOut)

		return bufOut.String(), bufErr.String(), err, isTimeout
	}
	return job
}

func (j Job) Run() {

	// fmt.Println(j.task)
	// return

	timeout := time.Duration(time.Hour * 24)
	if j.task.Timeout > 0 {
		timeout = time.Second * time.Duration(j.task.Timeout)
	}

	j.status++
	defer func() {
		j.status--
	}()

	// fmt.Println(timeout)
	t := time.Now()
	cmdOut, cmdErr, err, isTimeout := j.runFunc(timeout)

	ut := time.Now().Sub(t) / time.Millisecond

	log := new(models.AppCronLog)
	log.CronId = j.id
	log.Output = ConvertToString(cmdOut, "gbk", "utf8")
	log.Error = cmdErr
	log.ProcessTime = int(ut)
	log.CreateTime = t.Unix()

	if isTimeout {
		log.Status = models.CRON_TIMEOUT
		log.Error = fmt.Sprintf("任务执行超过 %d 秒\n----------------------\n%s\n", int(timeout/time.Second), cmdErr)
	} else if err != nil {
		log.Status = models.CRON_ERROR
		log.Error = err.Error() + ":" + cmdErr
	}

	j.logId, _ = models.CronLogAdd(log)

	// 更新上次执行时间
	j.task.PrevTime = t.Unix()
	j.task.ExecNum++
	j.task.Update("PrevTime", "ExecNum")
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
