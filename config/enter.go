package config

type Config struct {
	DB     DB     `yaml:"db"`
	Redis  Redis  `yaml:"redis"`
	System System `yaml:"system"`
	Jwt    Jwt    `yaml:"jwt"`
}
