package handler

import (
	"fmt"
	"github.com/19700101000000/system-sample/api/env"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
)

var (
	// TODO: randomize it
	oauthStateString  = "pseudo-random"
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost/api/auth/google/callback",
		ClientID:     env.GoogleClientID,
		ClientSecret: env.GoogleClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
)

func AuthGoogleSignin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
func AuthGoogleCallback(c *gin.Context) {
	v := struct {
		State string `form:"state"`
		Code  string `form:"code"`
	}{}
	c.Bind(&v)

	content, err := getUserInfo(v.State, v.Code)
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": string(content),
	})
}
func getUserInfo(state, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
