package bookctl

import (
	"github.com/gin-gonic/gin"
	"template/middleware/request"
	"template/middleware/response"
	"template/models/book"
	"template/service/book_service"
)

func AddBook(c *gin.Context) {
	in := &book.AddBookRequest{}
	if err := request.ValidParams(c, in); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := book_service.NewBookService().Create(in)
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
	date, count, err := book_service.NewBookService().List(in)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, map[string]interface{}{
		"list":  date,
		"count": count,
	})
}
