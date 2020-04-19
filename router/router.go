package router

import (
	"github.com/gin-gonic/gin"

	"github.com/blackironj/bookchive-server/handler"
	"github.com/blackironj/bookchive-server/middleware/jwt"
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

	v1Path := r.Group("/v1")
	v1Path.Use(jwt.CheckToken())
	{
		v1Path.GET("/books/:book_id", handler.GetBook)
		v1Path.POST("/libraries", handler.AddLib)

		v1Path.POST("/diaries", handler.AddDiary)

		v1Path.GET("/users/:user_uuid/libraries", handler.GetBooksInLibrary)
		v1Path.GET("/users/:user_uuid/diaries/:library_uk", handler.GetDiaries)
	}

	return r
}
