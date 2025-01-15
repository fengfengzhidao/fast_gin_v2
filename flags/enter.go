package flags

import (
	"fast_gin/global"
	"flag"
	"fmt"
)

type FlagOptions struct {
	File    string
	Version bool
	DB      bool
}

var Options FlagOptions

func Parse() {
	flag.StringVar(&Options.File, "f", "settings.yaml", "配置文件路径")
	flag.BoolVar(&Options.Version, "v", false, "打印当前版本")
	flag.BoolVar(&Options.DB, "db", false, "迁移表结构")
	flag.Parse()
}

func Run() (ok bool) {
	if Options.DB {
		fmt.Println("表结构迁移")
		return true
	}
	if Options.Version {
		fmt.Println("当前后端版本", global.Version)
		return true
	}
	return false
}
