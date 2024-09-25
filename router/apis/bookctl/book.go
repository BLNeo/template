package bookctl

import (
	"github.com/gin-gonic/gin"
	"template/service/bookService"
	"template/tool/request"
	"template/tool/response"
)

func AddBook(c *gin.Context) {
	in := &bookService.AddBookRequest{}
	if err := request.ValidParams(c, in); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := bookService.NewBookService().Create(in)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func ListBook(c *gin.Context) {
	in := &bookService.ListBookRequest{}
	if err := request.ValidParams(c, in); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	date, count, err := bookService.NewBookService().List(in)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, map[string]interface{}{
		"list":  date,
		"count": count,
	})
}
