package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"template/router/apis/bookctl"
	"template/router/apis/testctl"
)

// InitRouter initialize routing information
func InitRouter(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.Use(middleware.CORSMiddleware(), middleware.JWT())
	testRoute(r.Group("/test"))
	bookRoute(r.Group("/book"))
}

func testRoute(rg *gin.RouterGroup) {
	rg.GET("", testctl.GetTest)
}

func bookRoute(rg *gin.RouterGroup) {
	rg.POST("/add", bookctl.AddBook)
	rg.GET("/list", bookctl.ListBook)
}
