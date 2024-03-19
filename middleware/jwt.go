package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template/service"
	"template/tool/e"
	"template/tool/log"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		// rpc校验token
		token := c.GetHeader("token")
		if token == "" {
			token = c.Query("token")
		}
		err := service.NewSignRpcService().VerifyToken(token)
		if err != nil {
			log.Logger.Error(err.Error())
			code = e.ErrorAuth
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
