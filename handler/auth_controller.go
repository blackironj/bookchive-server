package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/blackironj/bookchive-server/oauth2/google"
)

//SigninWithGoogle sign-in with google account
func SigninWithGoogle(ctx *gin.Context) {
	google.SigninHandler(ctx)
}

//SigninWithGoogleCallback is a Oauth2 callback handler
func SigninWithGoogleCallback(ctx *gin.Context) {
	//do something
	ctx.JSON(200, gin.H{"message": "Hello, Google account"})
}
