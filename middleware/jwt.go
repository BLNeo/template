package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"template/conf"
	"template/tool/response"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte
var issuer = "" // 发行人
const TokenExpiresAt = time.Hour * 12

var once sync.Once

func InitJwtSecret() {
	once.Do(func() {
		jwtSecret = []byte(conf.Conf.App.JwtSecret)
		issuer = conf.Conf.App.TokenIssuer
	})
}

type Claims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int64, userName string) (string, error) {
	claims := Claims{
		userId,
		userName,

		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiresAt)),
			Issuer:    issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = response.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = response.Unauthorized
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = response.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
				code = response.ErrorAuthCheckTokenTimeout
			}
		}

		if code != response.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  response.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
