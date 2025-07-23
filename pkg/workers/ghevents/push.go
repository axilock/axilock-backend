package ghevents

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/axilock/axilock-backend/internal/axierr"
	"github.com/axilock/axilock-backend/internal/constants"
	"github.com/axilock/axilock-backend/internal/service/commitsvc"
	"github.com/axilock/axilock-backend/internal/service/reposvc"
	"github.com/google/go-github/v72/github"
	"github.com/rs/zerolog/log"
)

type PrComment struct {
	*EventClient
}

func (c *PrComment) GetTaskID() string {
	return "task:github:push"
}

func (c *PrComment) ProcessWebhook(ctx context.Context, data []byte) error {
	var event GHMeta
	err := json.Unmarshal(data, &event)
	if err != nil {
		return fmt.Errorf("cannot process event %v", err)
	}
	var ghevent github.PushEvent
	// ccreate the repo in db if not there

	err = json.Unmarshal(event.Data, &ghevent)
	if err != nil {
		return fmt.Errorf("cannot process event %v", err)
	}
	repoID, orgID, err := c.createRepoandOrg(ctx, ghevent)
	if err != nil {
		return err
	}

	err = c.addCommits(ctx, ghevent, repoID, orgID)
	if err != nil {
		return err
	}
	log.Info().Msg("github commits created")
	return nil
}

func (c *PrComment) createRepoandOrg(ctx context.Context, ghevent github.PushEvent) (int64, int64, error) {
	repostruct, err := reposvc.ExpludeRepoURL(ghevent.GetRepo().GetSSHURL())
	if err != nil {
		return 0, 0, fmt.Errorf("cannot process event %v", err)
	}
	var orgID int64

	if ghevent.GetRepo().GetOwner().GetType() == "Organization" {
		orgID = ghevent.GetOrganization().GetID()
	} else if ghevent.GetRepo().GetOwner().GetType() == "User" {
		orgID = ghevent.GetRepo().GetOwner().GetID()
	}
	org, err := c.GithubSvc.GetVCSInstallationID(ctx, orgID)
	if err != nil {
		if axierr.Is(err, axierr.ErrRecordNotFound) {
			log.Error().Err(err).Msg("cannot find org")
		}
		return 0, 0, fmt.Errorf("cannot process event %v", err)
	}
	var repoID int64
	repoID, err = c.RepoSvc.GetGithubRepo(ctx, ghevent.GetRepo().GetID())
	if err != nil {
		if axierr.Is(err, axierr.ErrRecordNotFound) {
			log.Info().Err(err).Msg("repo cannot be found creating new repo")
			repoID, err = c.RepoSvc.CreateGithubRepo(ctx, []reposvc.CreateRepoReq{
				{
					Name:      ghevent.GetRepo().GetName(),
					Repourl:   repostruct.Repopath,
					Org:       org.OrgID,
					Provider:  constants.SOURCE_GITHUB,
					VcsRepoID: ghevent.GetRepo().GetID(),
				},
			})
			if err != nil {
				log.Error().Err(err).Msg("cannot create github repo")
				return 0, 0, fmt.Errorf("cannot create github repo %v", err)
			}
		} else {
			log.Error().Err(err).Msg("cannot find github repo")
			return 0, 0, fmt.Errorf("cannot find github repo %v", err)
		}
	}
	return repoID, org.OrgID, nil
}

func (c *PrComment) addCommits(ctx context.Context, ghevent github.PushEvent, repoID int64, orgID int64) error {
	args := make([]commitsvc.CreateCommmitSvcDBParamas, 0, len(ghevent.GetCommits()))
	var verified bool
	for _, commit := range ghevent.GetCommits() {
		verified = true
		_, err := c.Services.CommitSvc.GetCommitByCommitID(ctx, commit.GetID())
		if err != nil {
			if axierr.Is(err, axierr.ErrRecordNotFound) {
				log.Info().Err(err).Msg("shadow commit detected")
				verified = false
			}
		}
		args = append(args, commitsvc.CreateCommmitSvcDBParamas{
			CommitID:     commit.GetID(),
			CommitTime:   commit.GetTimestamp().Time.Format(time.DateTime),
			PushTime:     commit.GetTimestamp().Time.Format(time.DateTime),
			AuthorName:   commit.GetCommitter().GetLogin(),
			AuthorEmail:  commit.GetCommitter().GetEmail(),
			ScannedByCli: verified,
		})
	}
	err := c.Services.CommitSvc.CreateCommitsFromGithub(ctx, commitsvc.CreateCommitGithubReq{
		RepoID:     repoID,
		CommitData: args,
		Org:        orgID,
	})
	if err != nil {
		log.Error().Err(err).Msg("cannot create commits")
		return fmt.Errorf("cannot create commits %v", err)
	}
	return nil
}
