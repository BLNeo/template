package bookctl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template/models/book"
	"template/service/book_service"
	"template/tool/e"
)

func AddBook(c *gin.Context) {
	appG := e.Gin{C: c}
	in := &book.AddBookRequest{}
	if err := c.ShouldBindJSON(in); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	err := book_service.NewBookService().Create(in)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func ListBook(c *gin.Context) {
	appG := e.Gin{C: c}
	in := &book.ListBookRequest{}
	if err := c.ShouldBindQuery(in); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	date, err := book_service.NewBookService().List(in)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, date)
}
