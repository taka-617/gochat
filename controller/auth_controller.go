package controller

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"

	"github.com/stretchr/gomniauth/providers/google"

	"github.com/taka-617/gochat/service"
)

var provider string

func init() {
	// setup gomniauth
	gomniauth.SetSecurityKey("zrgwvdvotd0fsehvrrerqcgom9ds06c1ib66")
	gomniauth.WithProviders(
		// github.New("3d1e6ba69036e0624b61", "7e8938928d802e7582908a5eadaaaf22d64babf1", "http://localhost:8080/auth/callback/github"),
		google.New("793646803752-kmeaf4kde7nhvr3qkrmsf7ahkpkmlt3k.apps.googleusercontent.com", "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", "http://localhost/auth/callback/google"),
		// facebook.New("537611606322077", "f9f4d77b3d3f4f5775369f5c9f88f65e", "http://localhost:8080/auth/callback/facebook"),
	)
}

func AuthLogin(c *gin.Context) {
	provider = c.Param("provider")

	provider, err := gomniauth.Provider(provider)
	if err != nil {
		c.Redirect(http.StatusOK, "/")
		return
	}

	loginURL, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		c.Redirect(http.StatusOK, "/")
		return
	}

	c.Request.Header.Set("Location", loginURL)
	c.Redirect(http.StatusTemporaryRedirect, loginURL)
}

func AuthCallback(c *gin.Context) {
	prov, err := gomniauth.Provider(provider)
	if err != nil {
		c.Redirect(http.StatusOK, "/")
		return
	}

	// get the credentials
	creds, err := prov.CompleteAuth(objx.MustFromURLQuery(c.Request.URL.RawQuery))
	if err != nil {
		c.Redirect(http.StatusOK, "/")
		return
	}

	// get the user
	user, err := prov.GetUser(creds)
	if err != nil {
		c.Redirect(http.StatusOK, "/")
		return
	}

	fmt.Println(user.Data()["id"])
	fmt.Println(user.Name())

	var auth_service service.AuthService

	auth_service.LoginOauth(provider, user.Data()["id"].(string), user.Name(), user.Email())

	c.HTML(200, "index.html", gin.H{})
}
