package conf

import (
	"github.com/spf13/viper"
	"template/tool/email"
	"template/tool/http"
	"template/tool/mysql"
	"template/tool/redis"
)

var Conf *Config

type App struct {
	Mode        string `yaml:"mode"`
	ServerName  string `yaml:"serverName"`
	JwtSecret   string `yaml:"jwtSecret"`
	TokenIssuer string `yaml:"tokenIssuer"`
}

type Grpc struct {
	Port string `yaml:"port"`
}

type Config struct {
	App       *App          `yaml:"app"`
	Mysql     *mysql.Config `yaml:"mysql"`
	Http      *http.Config  `yaml:"http"`
	Grpc      *Grpc         `yaml:"grpc"`
	Redis     *redis.Config `yaml:"redis"`
	EmailNTES *email.Config `yaml:"emailNTES"`
}

func LoadConfig() error {
	v := viper.New()
	v.AddConfigPath("conf")

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(&Conf)
	return err
}
