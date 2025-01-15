package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"github.com/sirupsen/logrus"
)

func main() {
	core.InitLogger()
	flags.Run()
	global.Config = core.ReadConfig()
	global.DB = core.InitGorm()
	global.Redis = core.InitRedis()

	logrus.Debugf("你好")
	logrus.Infof("你好")
	logrus.Warnf("你好")
	logrus.Errorf("你好")
}
