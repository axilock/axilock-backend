package gh

import (
	"context"

	"github.com/google/go-github/v72/github"
	"golang.org/x/oauth2"
)

func NewInstallationClient(acc_t string) *github.Client {
	tksrc := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: acc_t})
	httpClient := oauth2.NewClient(context.Background(), tksrc)
	client := github.NewClient(httpClient)
	return client
}
