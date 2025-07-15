package gh

import (
	"context"

	"github.com/axilock/axilock-backend/pkg/util"
	"github.com/google/go-github/v72/github"

	"golang.org/x/oauth2"
	oauthgithub "golang.org/x/oauth2/github"
)

type GithubUserClient struct {
	Client *github.Client
}

func GetGithubAuthEndpoint(config util.Config) string {
	oauthConfig := &oauth2.Config{
		ClientID:     config.GithubClientID,
		ClientSecret: config.GithubClientSecret,
		Endpoint:     oauthgithub.Endpoint,
		RedirectURL:  "https://app.axilock.ai/auth/github/callback",
	}
	return oauthConfig.AuthCodeURL("state")
}

func GetClientWithCode(ctx context.Context, config util.Config, code string) (GithubUserClient, error) {
	oauthConfig := &oauth2.Config{
		ClientID:     config.GithubClientID,
		ClientSecret: config.GithubClientSecret,
		Endpoint:     oauthgithub.Endpoint,
	}

	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return GithubUserClient{}, err
	}
	client := github.NewClient(nil).WithAuthToken(token.AccessToken)
	return GithubUserClient{
		Client: client,
	}, nil
}
