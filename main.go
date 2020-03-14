package main

import (
	"github.com/blackironj/bookchive-server/da"
	"github.com/blackironj/bookchive-server/env"
	"github.com/blackironj/bookchive-server/oauth2/google"
	"github.com/blackironj/bookchive-server/router"
)

func init() {
	env.Setup()

	google.SetupOAuth(
		env.Conf.Oauth.Google.CallbackURL,
		env.Conf.Oauth.Google.CredFilePath,
		env.Conf.Oauth.Google.Scopes)

	da.InitDB()
}

func main() {
	r := router.InitRouter()

	r.Run(":" + env.Conf.Server.Port)
}
