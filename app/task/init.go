package task

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/webcron/app/models"
	"os/exec"
	// "runtime"
	"time"
)

func Init() {
	fmt.Println("cron init")
	// runtime.GOMAXPROCS(runtime.NumCPU())
	list, _ := models.CronGetList(1, 1000000, "status", 1, "item_id__gt", 0)
	for _, task := range list {
		fmt.Println(task.CronSpec)
		job, err := NewJobFromTask(task)
		if err != nil {
			beego.Error("InitJobs:", err.Error())
			continue
		}
		AddJob(task.CronSpec, *job)
	}
	fmt.Println("cron end")
}

func runCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		beego.Warn(fmt.Sprintf("任务执行时间超过%d秒，进程将被强制杀掉: %d", int(timeout/time.Second), cmd.Process.Pid))
		go func() {
			<-done // 读出上面的goroutine数据，避免阻塞导致无法退出
		}()
		if err = cmd.Process.Kill(); err != nil {
			beego.Error(fmt.Sprintf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err))
		}
		return err, true
	case err = <-done:
		return err, false
	}
}