package main

import (
	"github.com/blackironj/bookchive-server/oauth2/google"
	"github.com/blackironj/bookchive-server/router"
)

const ( //FIXME: read from config file
	googleCallbackURL = "http://localhost:3000/auth/signin/google/callback"
	credFile          = "credfile.json"
)

func init() {
	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.email",
		// You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
	}

	google.SetupOAuth(googleCallbackURL, credFile, scopes)
}

func main() {
	r := router.InitRouter()

	r.Run(":3000")
}
