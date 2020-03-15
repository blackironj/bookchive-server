package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blackironj/bookchive-server/model"
	"github.com/blackironj/bookchive-server/oauth2/google"
	"github.com/blackironj/bookchive-server/service"
)

//SigninWithGoogle sign-in with google account
func SigninWithGoogle(ctx *gin.Context) {
	google.SigninHandler(ctx)
}

//SigninWithGoogleCallback is a Oauth2 callback handler
func SigninWithGoogleCallback(ctx *gin.Context) {
	signinData, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusBadRequest, "Cannot find a sign-in data")
		return
	}

	signinName := signinData.(google.User).Name

	user := &model.Users{
		Email: signinData.(google.User).Email,
		Name:  &signinName,
	}

	if err := service.Signin(user); err != nil {
		ctx.JSON(http.StatusBadRequest, "fail to signin")
		return
	}

	token, tokErr := service.GenJWT(user)
	if tokErr != nil {
		ctx.JSON(http.StatusInternalServerError, tokErr)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
