package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blackironj/bookchive-server/middleware/jwt"
	"github.com/blackironj/bookchive-server/service"
)

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("book_id")
	book, err := service.GetBook(bookID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func GetBooksInLibrary(ctx *gin.Context) {
	userUUID := ctx.Param("user_uuid")
	selfUUID, _ := ctx.Get(jwt.UUID_KEY)

	if userUUID == "me" {
		userUUID = selfUUID.(string)
	}

	if userUUID != selfUUID.(string) {
		ctx.JSON(http.StatusForbidden, "you cannot see other user's book list")
		return
	}

	books, err := service.GetBooksInLibrary(userUUID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, books)
}
