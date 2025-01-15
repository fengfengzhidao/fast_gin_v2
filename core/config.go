package core

import (
	"fast_gin/config"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile("settings.yaml")
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
