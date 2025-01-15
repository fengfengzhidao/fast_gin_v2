package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fmt"
)

func main() {
	core.InitLogger()
	flags.Parse()
	global.Config = core.ReadConfig()
	global.DB = core.InitGorm()
	global.Redis = core.InitRedis()

	if flags.Run() {
		return
	}

	fmt.Println("web服务")
}
