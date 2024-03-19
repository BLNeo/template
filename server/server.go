package server

import (
	signPb "github.com/BLNeo/protobuf-grpc-file/sign"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"template/conf"
	"template/models"
	"template/router"
	"template/service"
	"template/tool/grpc_server"
	"template/tool/log"
	"template/tool/mysql"
	"template/tool/util"
)

type Server struct {
	gin        *gin.Engine       // 路由服务
	db         *gorm.DB          // db数据库服务
	signClient signPb.SignClient // signRpc客户端
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
	// jwt
	util.InitJwtSecret()
	// gin
	gin.SetMode(conf.Http.Mode)
	s.gin = gin.Default()
	router.InitRouter(s.gin)

	// signGrpc
	signGrpcConn, err := grpc_server.InitSignGrpcConn(conf.SignGrpc.Host)
	if err != nil {
		return err
	}
	s.signClient = signPb.NewSignClient(signGrpcConn)
	service.InitSignRpcClient(s.signClient)

	log.Logger.Info("server init success")
	return nil
}

func (s *Server) Run() error {
	return s.gin.Run(conf.Conf.Http.Port)
}
