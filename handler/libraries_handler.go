package handler

import (
	"fmt"
	"net/http"

	"github.com/blackironj/bookchive-server/middleware/jwt"
	"github.com/blackironj/bookchive-server/model"
	"github.com/blackironj/bookchive-server/service"
	"github.com/gin-gonic/gin"
)

func AddLib(ctx *gin.Context) {
	var books []*model.Book
	if bindErr := ctx.ShouldBindJSON(&books); bindErr != nil {
		fmt.Println(bindErr)
		ctx.JSON(http.StatusBadRequest, bindErr.Error())
		return
	}

	data, _ := ctx.Get(jwt.UUID_KEY)
	uuid := data.(string)

	if err := service.AddLib(uuid, books); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
