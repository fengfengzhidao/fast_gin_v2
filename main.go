package main

import (
	"fast_gin/core"
	"fmt"
)

func main() {
	cfg := core.ReadConfig()
	fmt.Println(cfg.DB)
}
