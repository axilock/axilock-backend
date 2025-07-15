package githubsvc

import (
	"time"

	"github.com/google/go-github/v72/github"
)

const (
	AuthTokenType = "auth_token_type"
)

type (
	GhInstallationPayload = github.InstallationEvent
	GhPushEvent           = github.PushEvent
)

type GhTokenData struct {
	GithubID string `json:"github_id,omitempty"`
	OrgName  string `json:"org_name,omitempty"`
	OrgID    int32  `json:"org_id,omitempty"`
}

type CreateGithubOrgRequest struct {
	GhTokenData
	Org int32
	Exp time.Time
}

type GetGithubOrgRequest struct {
	Org int64
}

type GithubUserCoverage struct {
	Username  string `json:"username"`
	Onboarded bool   `json:"onboarded"`
}
