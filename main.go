package main

import (
	"template/conf"
	"template/server"
)

// @title go项目模板
// @version 1.0
func main() {
	if err := conf.LoadConfig(); err != nil {
		panic(err)
	}

	s := server.NewServer()
	if err := s.Init(conf.Conf); err != nil {
		panic(err)
	}
	err := s.Run()
	if err != nil {
		panic(err)
	}
}
