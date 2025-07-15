package ghevents

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/axilock/axilock-backend/internal/axierr"
	"github.com/axilock/axilock-backend/internal/constants"
	"github.com/axilock/axilock-backend/internal/service/reposvc"

	"github.com/google/go-github/v72/github"
	"github.com/rs/zerolog/log"
)

type PrOpen struct {
	*EventClient
}

func (e *PrOpen) GetTaskID() string {
	return "task:github:installation"
}

func (e *PrOpen) ProcessWebhook(ctx context.Context, data []byte) error {
	var event GHMeta
	err := json.Unmarshal(data, &event)
	if err != nil {
		return fmt.Errorf("cannot process event %v", err)
	}
	var ghevent github.InstallationEvent
	err = json.Unmarshal(event.Data, &ghevent)
	if err != nil {
		return fmt.Errorf("cannot process event %v", err)
	}
	client := e.GithubSvc.GetInstallationClient(event.Access_token)
	switch ghevent.GetAction() {
	case "created":
		err = e.processInstallCreated(ctx, ghevent, client)
	default:
		return nil
	}
	if err != nil {
		return err
	}
	log.Info().Msg("installation and repos created")
	return nil
}

func (e PrOpen) processInstallCreated(ctx context.Context, ghevent github.InstallationEvent, client *github.Client) error {
	org, err := e.GithubSvc.GetVCSInstallationID(ctx, ghevent.GetInstallation().GetTargetID())
	if err != nil {
		if axierr.Is(err, axierr.ErrRecordNotFound) {
			log.Error().Err(err).Msg("cannot find org")
			return nil
		}
		log.Error().Err(err).Msg("cannot find org")
		return nil
	}
	repos, resp, err := client.Repositories.ListByOrg(ctx, ghevent.GetInstallation().GetAccount().GetLogin(), nil)
	if err != nil {
		return fmt.Errorf("cannot get repos for org %s", ghevent.GetInstallation().GetAccount().GetLogin())
	}
	for resp.NextPage != 0 {
		r, rsp, err := client.Repositories.ListByOrg(ctx, ghevent.GetInstallation().GetAccount().GetLogin(), &github.RepositoryListByOrgOptions{
			ListOptions: github.ListOptions{
				Page:    resp.NextPage,
				PerPage: 100,
			},
		})
		if err != nil {
			return fmt.Errorf("cannot get repos for org %s", ghevent.GetInstallation().GetAccount().GetLogin())
		}
		resp = rsp
		repos = append(repos, r...)
	}

	args := make([]reposvc.CreateRepoReq, 0, len(repos))
	for _, r := range repos {

		_, err := e.RepoSvc.GetGithubRepo(ctx, r.GetID())
		if err != nil {
			if axierr.Is(err, axierr.ErrRecordNotFound) {
				log.Info().Err(err).Msg("repo cannot be found creating new repo")
				args = append(args, reposvc.CreateRepoReq{
					Name:      r.GetName(),
					Repourl:   r.GetSSHURL(),
					Org:       org.OrgID,
					Provider:  constants.SOURCE_GITHUB,
					VcsRepoID: r.GetID(),
				})
			}
		}
	}
	_, err = e.RepoSvc.CreateGithubRepo(ctx, args)
	if err != nil {
		log.Error().Err(err).Msg("cannot create repos")
		return fmt.Errorf("cannot create repos %v", err)
	}
	return nil
}
