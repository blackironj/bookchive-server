package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blackironj/bookchive-server/service"
)

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("book_id")
	book, err := service.GetBook(bookID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}
