package githubsvc

import (
	"context"
	"strconv"

	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/google/go-github/v72/github"

	"github.com/axilock/axilock-backend/pkg/gh"
	"github.com/axilock/axilock-backend/pkg/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type GithubService struct {
	config util.Config
	store  db.Store
	client gh.GithubAppInterface
}

func NewGithubService(store db.Store, config util.Config) GithubService {
	if config.RunningEnv == "test" {
		return GithubService{}
	}
	client, err := gh.NewAppClientwithID(config.GithubAppID, config.GithubAppPrivateKey)
	if err != nil {
		log.Fatal().Str("error", "cannot init github app client").Msg("error")
		return GithubService{}
	}
	return GithubService{
		config: config,
		store:  store,
		client: client,
	}
}

func (s GithubService) CreateGithubInstallation(ctx context.Context, installID string, userorg int64) error {
	installIDInt, err := strconv.ParseInt(installID, 10, 64)
	if err != nil {
		return err
	}
	ghinstall, err := s.client.FindInstallationByInstallID(ctx, installIDInt)
	if err != nil {
		return err
	}
	args := db.CreateGithubInstallationParams{
		OrgID:   userorg,
		OrgName: ghinstall.GetAccount().GetLogin(),
		Type:    "github",
		Token: pgtype.Text{
			String: "",
			Valid:  true,
		},
		InstallID: ghinstall.GetID(),
		VcsOrgID:  ghinstall.GetTargetID(),
	}
	err = s.store.CreateGithubInstallation(ctx, args)
	if err != nil {
		return err
	}
	return nil
}

func (s GithubService) GetGithubOrgByOrgID(ctx context.Context, id int64) (db.GitProvider, error) {
	return s.store.GetGithubInstallation(ctx, db.GetGithubInstallationParams{
		OrgID: pgtype.Int8{
			Int64: id,
			Valid: true,
		},
	})
}

func (s GithubService) GetVCSInstallationID(ctx context.Context, id int64) (db.GitProvider, error) {
	return s.store.GetGithubInstallation(ctx, db.GetGithubInstallationParams{
		VcsOrgID: pgtype.Int8{
			Int64: id,
			Valid: true,
		},
	})
}

func (s GithubService) GetINstallationAccessToken(ctx context.Context, id int64) (string, error) {
	return s.client.GetInstallAccessToken(ctx, id)
}

func (s GithubService) GetInstallationClient(acc_t string) *github.Client {
	return s.client.GetInstallationClient(acc_t)
}

// func (s GithubService) GetInstallationIDByOrgID(ctx context.Context, orgID int64) (int64, error) {
// 	return s.store.GetGithubInstallationForOrg(ctx, pgtype.Int8{
// 		Int64: orgID,
// 		Valid: true,
// 	})
// }

func (s GithubService) GetUserCoverage(ctx context.Context, orgID int64) ([]GithubUserCoverage, error) {
	org, err := s.store.GetGithubInstallation(ctx, db.GetGithubInstallationParams{
		OrgID: pgtype.Int8{
			Int64: orgID,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}
	client, err := s.client.GetClientForOrgname(ctx, org.OrgName)
	if err != nil {
		return nil, err
	}
	users, resp, err := client.Organizations.ListMembers(ctx, org.OrgName, &github.ListMembersOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
		PublicOnly: false,
	})
	for resp.NextPage != 0 {
		users, resp, err = client.Organizations.ListMembers(ctx, org.OrgName, &github.ListMembersOptions{
			ListOptions: github.ListOptions{
				PerPage: 100,
				Page:    resp.NextPage,
			},
			PublicOnly: false,
		})
		users = append(users, users...)
	}
	if err != nil {
		return nil, err
	}
	ghusers := make(map[string]bool)
	for _, user := range users {
		ghusers[user.GetLogin()] = true
	}
	commitusers, err := s.store.GetUniqueCommitUsernames(ctx, orgID)
	if err != nil {
		return nil, err
	}
	result := []GithubUserCoverage{}
	for _, user := range commitusers {
		if _, ok := ghusers[user]; ok {
			result = append(result, GithubUserCoverage{
				Username:  user,
				Onboarded: true,
			})
		} else {
			result = append(result, GithubUserCoverage{
				Username:  user,
				Onboarded: false,
			})
		}
	}
	return result, nil
}
