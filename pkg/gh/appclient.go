package gh

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v72/github"
	"golang.org/x/oauth2"
)

type GithubAppInterface interface {
	FindInstallationByInstallID(ctx context.Context, installID int64) (*github.Installation, error)
	GetInstallAccessToken(ctx context.Context, installid int64) (string, error)
	FindInstallationByName(ctx context.Context, name string) (*github.Installation, error)
	GetInstallationClient(token string) *github.Client
	GetClientForOrgname(ctx context.Context, orgname string) (*github.Client, error)
}

type GithubAppClient struct {
	client      *github.Client
	appTokensrc oauth2.TokenSource
}

func NewAppClientwithID(appidstr string, appPrivkey string) (GithubAppInterface, error) {
	appid, err := strconv.ParseInt(appidstr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse appid")
	}
	appTokenSource, err := NewApplicationTokenSource(appid, []byte(appPrivkey))
	if err != nil {
		return nil, fmt.Errorf("cannot parse appid")
	}
	httpClient := oauth2.NewClient(context.Background(), appTokenSource)
	client := github.NewClient(httpClient)
	return &GithubAppClient{
		client:      client,
		appTokensrc: appTokenSource,
	}, nil
}

func (c *GithubAppClient) FindInstallationByInstallID(ctx context.Context, installID int64) (*github.Installation, error) {
	installations, _, err := c.client.Apps.ListInstallations(ctx, &github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get installations for app")
	}
	for _, v := range installations {
		if installID == v.GetID() {
			return v, nil
		}
	}
	return nil, fmt.Errorf("cannot find installation for this install id")
}

func (c *GithubAppClient) GetInstallAccessToken(ctx context.Context, installid int64) (string, error) {
	token, _, err := c.client.Apps.CreateInstallationToken(ctx, installid, &github.InstallationTokenOptions{})
	if err != nil {
		return "", fmt.Errorf("cannot get installations for app")
	}
	return token.GetToken(), nil
}

func (c *GithubAppClient) FindInstallationByName(ctx context.Context, name string) (*github.Installation, error) {
	i, _, err := c.client.Apps.FindOrganizationInstallation(ctx, name)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (c *GithubAppClient) GetInstallationClient(token string) *github.Client {
	return NewInstallationClient(token)
}

func (c *GithubAppClient) GetInstallUsers(ctx context.Context, orgname string) ([]*github.User, error) {
	users, resp, err := c.client.Organizations.ListMembers(ctx, orgname, &github.ListMembersOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
		PublicOnly: false,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get users for org %s", orgname)
	}
	for resp.NextPage != 0 {
		users, resp, err = c.client.Organizations.ListMembers(ctx, orgname, &github.ListMembersOptions{
			ListOptions: github.ListOptions{
				PerPage: 100,
				Page:    resp.NextPage,
			},
			PublicOnly: false,
		})
		if err != nil {
			return nil, fmt.Errorf("cannot get users for org %s", orgname)
		}
		users = append(users, users...)
	}
	return users, nil
}

func (c *GithubAppClient) GetClientForOrgname(ctx context.Context, orgname string) (*github.Client, error) {
	install, err := c.FindInstallationByName(ctx, orgname)
	if err != nil {
		return nil, err
	}
	token, err := c.GetInstallAccessToken(ctx, install.GetID())
	if err != nil {
		return nil, err
	}
	return NewInstallationClient(token), nil
}
