package config

type Jwt struct {
	Expires int    `yaml:"expires"`
	Issuer  string `yaml:"issuer"`
	Key     string `yaml:"key"`
}
