package http

type Config struct {
	Port         string `yaml:"port"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
}
