package bookctl

import (
	"github.com/gin-gonic/gin"
	"template/middleware/request"
	"template/middleware/response"
	"template/models/book"
	"template/service/bookService"
)

func AddBook(c *gin.Context) {
	in := &book.AddBookRequest{}
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
	in := &book.ListBookRequest{}
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
