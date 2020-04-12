package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blackironj/bookchive-server/model"
	"github.com/blackironj/bookchive-server/service"
)

func AddDiary(ctx *gin.Context) {
	var diary model.Diary
	if bindErr := ctx.BindJSON(&diary); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, bindErr.Error())
		return
	}

	if addErr := service.AddDiary(&diary); addErr != nil {
		ctx.JSON(http.StatusNotFound, addErr.Error())
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
