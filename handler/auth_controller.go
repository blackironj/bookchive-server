package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blackironj/bookchive-server/da"
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

	var name sql.NullString
	signinName := signinData.(google.User).Name
	if signinName != "" {
		name.String = signinName
		name.Valid = true
	}

	user := &da.Users{
		Email: signinData.(google.User).Email,
		Name:  name,
	}

	if err := service.Signin(user); err != nil {
		ctx.JSON(http.StatusBadRequest, "fail to signin")
		return
	}

	//TODO: Generate a JWT and then return jwt to user
	ctx.JSON(http.StatusOK, "SUCCESS")
}
