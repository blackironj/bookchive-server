package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/blackironj/bookchive-server/middleware/jwt"
	"github.com/blackironj/bookchive-server/model"
	"github.com/blackironj/bookchive-server/service"
)

func GetDiaries(ctx *gin.Context) {
	userUUID := ctx.Param("user_uuid")
	selfUUID, _ := ctx.Get(jwt.UUID_KEY)

	if userUUID == "me" {
		userUUID = selfUUID.(string)
	}

	if userUUID != selfUUID.(string) {
		ctx.JSON(http.StatusForbidden, "cannot see other user's diaries")
		return
	}

	libUKstr := ctx.Param("library_uk")
	libUK, convErr := strconv.Atoi(libUKstr)
	if convErr != nil {
		ctx.JSON(http.StatusBadRequest, convErr.Error())
		return
	}

	diaries, getErr := service.GetDiaries(userUUID, libUK)
	if getErr != nil {
		ctx.JSON(http.StatusNotFound, getErr.Error())
		return
	}

	ctx.JSON(http.StatusOK, diaries)
}

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
