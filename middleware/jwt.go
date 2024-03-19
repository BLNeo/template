package middleware

import (
	"context"
	signPb "github.com/BLNeo/protobuf-grpc-file/sign"
	"github.com/gin-gonic/gin"
	"net/http"
	"template/service"
	"template/tool/e"
	"template/tool/log"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		// rpc校验token
		token := c.GetHeader("token")
		if token == "" {
			code = e.InvalidParams
		} else {
			ctx, f := context.WithTimeout(context.Background(), time.Duration(3*int(time.Second)))
			defer f()
			in := &signPb.VerifyTokenRequest{
				Token: token,
			}
			result, err := service.GetSignRpcClient().VerifyToken(ctx, in)
			if err != nil || !result.Enabled {
				log.Logger.Error(err.Error())
				code = e.ErrorAuthCheckTokenFail
			}
			// todo 数据补充 个人信息保存上下文中
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
	//return func(c *gin.Context) {
	//	var code int
	//	var data interface{}
	//	code = e.SUCCESS
	//	token := c.Query("token")
	//	if token == "" {
	//		code = e.InvalidParams
	//	} else {
	//		claims, err := util.ParseToken(token)
	//		if err != nil {
	//			code = e.ErrorAuthCheckTokenFail
	//		} else if time.Now().Unix() > claims.ExpiresAt {
	//			code = e.ErrorAuthCheckTokenTimeout
	//		}
	//	}
	//	if code != e.SUCCESS {
	//		c.JSON(http.StatusUnauthorized, gin.H{
	//			"code": code,
	//			"msg":  e.GetMsg(code),
	//			"data": data,
	//		})
	//		c.Abort()
	//		return
	//	}
	//	c.Next()
	//}

}
