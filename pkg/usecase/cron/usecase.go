package cron

type Usecase interface {
	CronPakan(waktu string) error
	InitCron()
}
