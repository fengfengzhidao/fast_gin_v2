package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fmt"
)

func main() {
	flags.Run()
	global.Config = core.ReadConfig()
	global.Config.DB.Port = 3307
	fmt.Println(global.Config.DB)
}
