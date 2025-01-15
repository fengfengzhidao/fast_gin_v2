package config

type Config struct {
	DB    DB    `yaml:"db"`
	Redis Redis `yaml:"redis"`
}
