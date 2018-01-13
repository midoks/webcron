package task

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"github.com/midoks/webcron/app/models"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
	
	defer session.Close()

	return session, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func ConnectByRsa(user string, host string, port int) (*ssh.Session, error) {

	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	rsaContent, rsaErr := ioutil.ReadFile(fmt.Sprintf("conf/%s", beego.AppConfig.String("local.id_rsa")))
	if rsaErr != nil {
		beego.Warn(beego.AppConfig.String("local.id_rsa"), rsaErr)
		return session, nil
	}

	rsaValue := []byte(rsaContent)
	pKeys, pErr := ssh.ParseRawPrivateKey(rsaValue)
	if pErr != nil {
		beego.Warn(fmt.Sprintf("Unable to parse test key %s: %v", pKeys, pErr))
		return nil, pErr
	}

	signer, serr := ssh.NewSignerFromKey(pKeys)
	if serr != nil {
		beego.Warn(fmt.Sprintf("NewSignerFromKey:", serr))
		return nil, serr
	}

	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.PublicKeys(signer))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	defer session.Close()

	return session, nil
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

		var err error
		isTimeout := false

		if item.Type == 0 {
			server, _ := models.ServerGetById(item.ServerId)

			var (
				session *ssh.Session
			)

			if server.Type == 0 {
				session, err = ConnectByUser(server.User, server.Pwd, server.Ip, server.Port)

				if err != nil {
				} else {
					
					session.Stdout = bufOut
					session.Stderr = bufErr
					session.Run(cron.Cmd)
				}

			} else {
				session, err = ConnectByRsa(server.User, server.Ip, server.Port)

				if err != nil {
				} else {
					// defer session.Close()
					session.Stdout = bufOut
					session.Stderr = bufErr
					session.Run(cron.Cmd)
				}

			}
			isTimeout = false
			if err != nil {
				beego.Debug(err)
				return bufOut.String(), bufErr.String(), err, isTimeout
			}

		} else {
			var cmd *exec.Cmd
			if IsWin() {
				cmd = exec.Command("cmd", "/c", cron.Cmd)
			} else {
				cmd = exec.Command("/bin/bash", "-c", cron.Cmd)
			}

			cmd.Stdout = bufOut
			cmd.Stderr = bufErr
			cmd.Start()

			err, isTimeout := runCmdWithTimeout(cmd, timeout)
			if err != nil {
				beego.Warn("runCmdWithTimeout:", err, isTimeout)
			}
		}

		return bufOut.String(), bufErr.String(), err, isTimeout
	}
	return job
}

func (j *Job) Status() int {
	return j.status
}

func (j *Job) GetName() string {
	return j.name
}

func (j *Job) GetId() int {
	return j.id
}

func (j *Job) GetLogId() int64 {
	return j.logId
}

func (j *Job) Run() {



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
