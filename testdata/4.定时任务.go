package main

import "fast_gin/service/cron_ser"

func main() {
	cron_ser.CronInit()

	select {}
}
