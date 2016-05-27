package core

import (
	"time"

	"gopkg.in/robfig/cron.v2"
)

//此类主要是针对session 的状态进行管理
var mainCron *cron.Cron

//StartSessionJob 定时任务
func StartSessionJob() {
	mainCron = cron.New()
	mainCron.Start()
}

//AddJob 添加一个定时任务
func AddSessionJob(cron string, cmd func()) (cron.EntryID, error) {

	return mainCron.AddFunc(cron, cmd)
}

//AddSessionJobForDurationt 添加一个定时任务
func AddSessionJobForDuration(duration time.Duration, cmd func()) (cron.EntryID, error) {
	return mainCron.Schedule(cron.Every(duration), newJobs(cmd)), nil
}

// StopSessionAllJob 停止所有的定时任务
func StopSessionAllJob() {
	mainCron.Stop()
}

// Remove 移除一个定时任务
func RemoveSessionJob(id cron.EntryID) {
	mainCron.Remove(id)
}

type jobs struct {
	cmd func()
}

func (j *jobs) Run() {
	j.cmd()
}
func newJobs(cmd func()) *jobs {
	return &jobs{cmd}
}
