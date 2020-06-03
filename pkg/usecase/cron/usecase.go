package cron

type Usecase interface {
	CronPakan(waktu string) error
	CronNotifGuideline() error
	InitCron()
}
