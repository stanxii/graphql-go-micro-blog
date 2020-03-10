package config

import "github.com/jinzhu/configor"

var Config = struct {
	Mysql struct {
		Type     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Host     string `required:"true"`
		Port     string `default:"3306"`
		Table     string `require:"true"`
	}
}{}

func init() {
	configor.Load(&Config, "./config.yml")
}