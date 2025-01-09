package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"os"
	"product_api/helpers"
)

func OAuthConfig() *oauth2.Config {
	helpers.LoadEnv()
	oauthConf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return oauthConf
}
