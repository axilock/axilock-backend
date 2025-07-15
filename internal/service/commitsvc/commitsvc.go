package commitsvc

import (
	"context"
	"fmt"
	"time"

	"github.com/axilock/axilock-backend/internal/axierr"
	"github.com/axilock/axilock-backend/internal/constants"
	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5/pgtype"
)

type CommitServiceInterface interface {
	CreateCommitsForSession(ctx context.Context, req CreateCommitGrpcReq, sessionID string) error
	GetCommitByCommitID(ctx context.Context, commitID string) (int64, error)
	CreateCommitsFromGithub(ctx context.Context, args CreateCommitGithubReq) error
}

type CommitService struct {
	store db.Store
}

func NewCommitService(store db.Store) CommitServiceInterface {
	return &CommitService{
		store: store,
	}
}

func (s *CommitService) CreateCommitsForSession(ctx context.Context, req CreateCommitGrpcReq, sessionID string) error {

	commitargs := make([]db.CreateCommmitsFromRPCParams, 0, len(req.CommitData))
	for _, val := range req.CommitData {
		cTime, err := time.Parse(time.DateTime, val.CommitTime)
		if err != nil {
			return fmt.Errorf("cannot parse commit time format  %w", err)
		}
		pTime, err := time.Parse(time.DateTime, val.PushTime)
		if err != nil {
			return fmt.Errorf("cannot parse push time forma %w", err)
		}
		commitargs = append(commitargs, db.CreateCommmitsFromRPCParams{
			Sessionid: sessionID,
			Author:    pgtype.Text{String: val.AuthorName, Valid: true},
			CommitTime: pgtype.Timestamptz{
				Time:  cTime,
				Valid: !cTime.IsZero(),
			},
			PushTime: pgtype.Timestamptz{
				Time:  pTime,
				Valid: !pTime.IsZero(),
			},
			CommitID:    val.CommitID,
			Source:      constants.SOURCE_AXI_CLI,
			Org:         req.Org,
			UserID:      req.UserID,
			UserRepoUrl: req.RepoURL,
		})
	}

	_, err := s.store.CreateCommmitsFromRPC(ctx, commitargs)
	if err != nil {
		if axierr.IsUniqueViolation(err) {
			return err
		}
		return fmt.Errorf("cannot insert commit data to db %w", err)
	}
	return nil
}

// func (s *CommitService) GetCommitsHealth(ctx context.Context, org int64) (db.GetProtectedCommitsRow, error) {
// 	return s.store.GetProtectedCommits(ctx, org)
// }

func (s *CommitService) GetCommitByCommitID(ctx context.Context, commitID string) (int64, error) {
	return s.store.GetCommitByCommitID(ctx, commitID)
}

func (s *CommitService) CreateCommitsFromGithub(ctx context.Context, args CreateCommitGithubReq) error {
	commitargs := make([]db.CreateVCSCommitParams, 0, len(args.CommitData))
	for _, val := range args.CommitData {
		cTime, err := time.Parse(time.DateTime, val.CommitTime)
		if err != nil {
			return fmt.Errorf("cannot parse commit time format  %w", err)
		}
		if val.ScannedByCli {
			err := s.store.LinkRepoInCommitCli(ctx, db.LinkRepoInCommitCliParams{
				CommitID: val.CommitID,
				Org:      args.Org,
				Repo:     pgtype.Int8{Valid: true, Int64: args.RepoID}})
			if err != nil {
				log.Error().Err(err).Msg("cannot link repo in commit cli")
			}
		}
		commitargs = append(commitargs, db.CreateVCSCommitParams{
			CommitID:    val.CommitID,
			AuthorName:  val.AuthorName,
			AuthorEmail: val.AuthorEmail,
			CommitTime: pgtype.Timestamptz{
				Time:  cTime,
				Valid: !cTime.IsZero(),
			},
			Org:          args.Org,
			Provider:     constants.SOURCE_GITHUB,
			ScannedByCli: pgtype.Bool{Valid: true, Bool: val.ScannedByCli},
			ScannedByRunner: pgtype.Bool{
				Valid: true,
				Bool:  false,
			},
			RepoID: args.RepoID,
		})
	}
	_, err := s.store.CreateVCSCommit(ctx, commitargs)
	if err != nil {
		return fmt.Errorf("cannot insert commit data to db %w", err)
	}
	return nil
}
