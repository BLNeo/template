package testctl

import (
	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})

}
