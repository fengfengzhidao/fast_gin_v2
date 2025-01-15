package core

import (
	"fast_gin/config"
	"fast_gin/flags"
	"fast_gin/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile(flags.Options.File)
	if err != nil {
		fmt.Printf("配置文件读取错误 %s", err)
		return
	}
	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		fmt.Printf("配置文件格式错误 %s", err)
		return
	}
	return
}

func DumpConfig() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		fmt.Printf("配置文件转换错误 %s", err)
		return
	}
	err = os.WriteFile(flags.Options.File, byteData, 0666)
	if err != nil {
		fmt.Printf("配置文件写入错误 %s", err)
		return
	}
	fmt.Println("配置文件写入成功")
}
