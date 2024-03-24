package middleware

import (
	"github.com/gin-gonic/gin"
	"template/service"
	"template/tool/response"
	"template/tool/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = response.SUCCESS
		// rpc校验token
		token := c.GetHeader("token")
		if token == "" {
			token = c.Query("token")
		}
		data, err := service.NewSignRpcService().VerifyToken(token)
		if err != nil {
			code = response.Unauthorized
		}

		if code != response.SUCCESS {
			response.UnauthorizedError(c, err.Error())
			c.Abort()
			return
		}
		// 用户id保存上下文中
		c.Set(util.CtxUserId, data.UserId)
		c.Next()
	}
}
