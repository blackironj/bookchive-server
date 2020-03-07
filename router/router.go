package router

import (
	"github.com/gin-gonic/gin"

	"github.com/blackironj/bookchive-server/handler"
	"github.com/blackironj/bookchive-server/oauth2/google"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authPath := r.Group("/auth")
	{
		signinPath := authPath.Group("/signin/google")
		{
			signinPath.Use(google.Session())
			signinPath.GET("", handler.SigninWithGoogle)
			signinPath.GET("/callback", google.Auth(), handler.SigninWithGoogleCallback)
		}
	}

	return r
}
