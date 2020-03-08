package main

import (
	"github.com/blackironj/bookchive-server/env"
	"github.com/blackironj/bookchive-server/oauth2/google"
	"github.com/blackironj/bookchive-server/router"
)

const ( //FIXME: read from config file
	googleCallbackURL = "http://localhost:3000/auth/signin/google/callback"
	credFile          = "credfile.json"
)

func init() {
	env.Setup()

	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}

	google.SetupOAuth(googleCallbackURL, credFile, scopes)
}

func main() {
	r := router.InitRouter()

	r.Run(":3000")
}
