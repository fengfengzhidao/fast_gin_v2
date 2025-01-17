package config

type SiteLogin struct {
	Captcha bool `yaml:"captcha"`
}

type Site struct {
	Login SiteLogin `yaml:"login"`
}
