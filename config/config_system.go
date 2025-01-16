package config

import "fmt"

type System struct {
	Mode string `yaml:"mode"`
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.IP, s.Port)
}
