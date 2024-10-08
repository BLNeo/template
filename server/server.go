package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"template/conf"
	"template/middleware"
	"template/models"
	"template/router"
	"template/tool/log"
	"template/tool/mysql"
	"template/tool/redis"
	"template/tool/request"
)

type Server struct {
	gin *gin.Engine // 路由服务
	db  *gorm.DB    // db数据库服务
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(conf *conf.Config) error {
	var err error
	log.Init(conf.App.ServerName)

	s.db, err = mysql.InitEngine(conf.Mysql)
	if err != nil {
		return err
	}

	// 赋值到models包的db
	models.InitDb(s.db)
	// 表同步
	models.AutoMigrate()

	// redis
	err = redis.InitClient(conf.Redis)
	if err != nil {
		return err
	}

	// jwt
	middleware.InitJwtSecret()
	// gin
	gin.SetMode(conf.App.Mode)
	s.gin = gin.Default()
	router.InitRouter(s.gin)
	// validator // 入参校验翻译器
	err = request.InitTrans()
	if err != nil {
		return err
	}

	log.Logger.Info("server init success")
	return nil
}

func (s *Server) Run() error {
	return s.gin.Run(":" + conf.Conf.Http.Port)
}
