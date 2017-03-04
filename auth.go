package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func NewOAuthClient() *http.Client {
	pat := os.Getenv("DIGIOCEAN_TOKEN")
	if len(pat) < 1 {
		fmt.Println("DIGIOCEAN_TOKEN environmet variable must not be empty.")
		os.Exit(1)
	}

	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	return oauthClient
}
