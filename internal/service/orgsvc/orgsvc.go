package orgsvc

import (
	"context"

	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type OrgServiceInterface interface {
	GetOrgByID(ctx context.Context, id int64) (db.Organisation, error)
	GetReposForOrg(ctx context.Context, org int64) (int64, error)
	GetOrgbyGithubOrgID(ctx context.Context, orgID int64) (db.Organisation, error)
}

type OrgService struct {
	store db.Store
}

func NewOrgService(db db.Store) OrgServiceInterface {
	return &OrgService{
		store: db,
	}
}

// func (s *OrgService) CreateRegexForOrg(ctx context.Context, req CreateRegexReq) error {
// 	args := db.CreateRegexForOrgParams{
// 		Org: pgtype.Int8{
// 			Int64: req.Org,
// 			Valid: true,
// 		},
// 		Regstring: req.RegexStr,
// 		Description: pgtype.Text{
// 			String: req.Desc,
// 			Valid:  true,
// 		},
// 		Version: 1,
// 		Name:    fmt.Sprintf("CUSTOM_%s", req.Name),
// 	}

// 	err := s.store.CreateRegexForOrg(ctx, args)
// 	if err != nil {
// 		return fmt.Errorf("cannot create regex for org %w", err)
// 	}
// 	return nil
// }

// func (s *OrgService) GetRegexForOrg(ctx context.Context, orgID int64) ([]db.Regex, error) {
// 	return s.store.GetRegexesForOrg(ctx,
// 		pgtype.Int8{
// 			Int64: orgID,
// 			Valid: true,
// 		})
// }

func (s *OrgService) GetOrgByID(ctx context.Context, id int64) (db.Organisation, error) {
	return s.store.GetOrgByID(ctx, pgtype.Int8{
		Int64: id,
		Valid: true,
	})
}

func (s *OrgService) GetReposForOrg(ctx context.Context, org int64) (int64, error) {
	return s.store.GetRepoCountForOrg(ctx, org)
}

func (s *OrgService) GetOrgbyGithubOrgID(ctx context.Context, orgID int64) (db.Organisation, error) {
	return s.store.GetOrgByEntity(ctx, db.GetOrgByEntityParams{
		GithubOrgID: pgtype.Int8{
			Int64: orgID,
			Valid: true,
		},
	})
}
