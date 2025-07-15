package reposvc

import (
	"context"

	"github.com/axilock/axilock-backend/internal/constants"
	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type RepoService struct {
	store db.Store
}

func NewRepoService(db db.Store) RepoService {
	return RepoService{
		store: db,
	}
}

// func (s RepoService) CreateRepoBulk(ctx context.Context, repos []CreateRepoReq) error {
// 	for _, repo := range repos {
// 		data, err := ExpludeRepoURL(repo.Repourl)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s.store.CreateRepo(ctx, db.CreateRepoParams{
// 			Name:    repo.Name,
// 			Repourl: data.Repopath,
// 			Org:     repo.Org,
// 			Provider: pgtype.Text{
// 				String: repo.Provider,
// 				Valid:  true,
// 			},
// 		})
// 		if err != nil {
// 			if !axierr.Is(err, axierr.ErrUniqueViolation) {
// 				log.Error().Err(err).Msg("cannot create repo")
// 			}
// 		}
// 	}
// 	return nil
// }

// func (s RepoService) GetRepoStats(ctx context.Context, orgID int64) (db.GetRepoCoverageRow, error) {
// 	repos, err := s.store.GetRepoCoverage(ctx, orgID)
// 	if err != nil {
// 		return db.GetRepoCoverageRow{}, err
// 	}
// 	return repos, nil
// }

func (s RepoService) CreateGithubRepo(ctx context.Context, repos []CreateRepoReq) (int64, error) {
	args := make([]db.CreateRepoWithProviderParams, 0, len(repos))
	for _, repo := range repos {
		args = append(args, db.CreateRepoWithProviderParams{
			Name:      repo.Name,
			Repourl:   repo.Repourl,
			Org:       repo.Org,
			Provider:  repo.Provider,
			VcsRepoID: repo.VcsRepoID,
		})
	}
	_, err := s.store.CreateRepoWithProvider(ctx, args)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (s RepoService) GetGithubRepo(ctx context.Context, githubRepoID int64) (int64, error) {
	repos, err := s.store.GetRepoByEntity(ctx, db.GetRepoByEntityParams{
		VcsRepoID: pgtype.Int8{
			Int64: githubRepoID,
			Valid: true,
		},
		Provider: pgtype.Text{
			String: constants.SOURCE_GITHUB,
			Valid:  true,
		},
	})
	if err != nil {
		return 0, err
	}
	return repos.ID, nil
}
