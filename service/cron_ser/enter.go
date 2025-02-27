package cron_ser

import (
	"github.com/robfig/cron/v3"
	"time"
)

func CronInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	crontab := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))

	//crontab.AddFunc("0 0 23 * * *", HelloCron)
	crontab.Start()
}
