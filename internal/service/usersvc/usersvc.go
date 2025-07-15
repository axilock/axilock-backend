package usersvc

import (
	"context"
	"fmt"

	"github.com/axilock/axilock-backend/internal/axierr"
	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/axilock/axilock-backend/pkg/gh"
	"github.com/axilock/axilock-backend/pkg/util"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserServiceInterface interface {
	// CreateUser(ctx context.Context, req CreateUserRequest) (db.User, db.Organisation, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	GetUserByID(ctx context.Context, userID string) (db.User, error)
	CreateUserWithGithub(ctx context.Context, code string, provider string) (db.User, error)
	GetUserByGithubID(ctx context.Context, githubID int64) (db.User, error)
}

type UserService struct {
	store  db.Store
	config util.Config
}

func NewUserService(store db.Store, c util.Config) UserServiceInterface {
	return &UserService{
		store:  store,
		config: c,
	}
}

// func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (db.User, db.Organisation, error) {
// 	var emptyUser db.User
// 	var emptyOrg db.Organisation
// 	domain, err := util.GetDomain(req.Email)
// 	if err != nil {
// 		return emptyUser, emptyOrg, fmt.Errorf("cannot get domain")
// 	}
// 	var org db.Organisation
// 	org, err = s.store.GetOrgByEntity(ctx, db.GetOrgByEntityParams{
// 		Domain: pgtype.Text{
// 			String: domain,
// 			Valid:  true,
// 		},
// 	})
// 	if err != nil {
// 		if axierr.Is(err, axierr.ErrRecordNotFound) {
// 			orgArgs := db.CreateOrganiationParams{
// 				Domain: domain,
// 				Name:   strings.Split(domain, ".")[0],
// 			}
// 			org, err = s.store.CreateOrganiation(ctx, orgArgs)
// 			if err != nil {
// 				return emptyUser, emptyOrg, fmt.Errorf("failed to create organisation")
// 			}
// 		} else {
// 			return emptyUser, emptyOrg, fmt.Errorf("internal server error")
// 		}
// 	}
// 	hasPass, err := util.NewPassword(req.Password)
// 	if err != nil {
// 		return emptyUser, emptyOrg, fmt.Errorf("failed to hash password")
// 	}

// 	arg := db.CreateUserParams{
// 		Email:        req.Email,
// 		Org:          org.ID,
// 		HashPassword: hasPass,
// 	}
// 	user, err := s.store.CreateUser(ctx, arg)
// 	if err != nil {
// 		if axierr.Is(err, axierr.ErrUniqueViolation) {
// 			return emptyUser, emptyOrg, err
// 		}
// 		return emptyUser, emptyOrg, err
// 	}
// 	return user, org, nil
// }

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return s.store.GetUserByEmail(ctx, email)
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (db.User, error) {
	return s.store.GetUserByEntityId(ctx, userID)
}

func (s *UserService) GetUserByGithubID(ctx context.Context, githubID int64) (db.User, error) {
	return s.store.GetUserByGithubId(ctx, pgtype.Int8{
		Int64: githubID,
		Valid: true,
	})
}

func (s *UserService) CreateUserWithGithub(ctx context.Context, code string, provider string) (db.User, error) {
	ghuser, err := gh.GetClientWithCode(ctx, s.config, code)
	if err != nil {
		log.Error().Err(err).Msg("failed to get client with code")
		return db.User{}, err
	}
	ghuserdetails, _, err := ghuser.Client.Users.Get(ctx, "")
	if err != nil {
		log.Error().Err(err).Msg("failed to get user details")
		return db.User{}, err
	}
	var useremail string
	emails, _, err := ghuser.Client.Users.ListEmails(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to get emails")
		return db.User{}, err
	}
	for _, email := range emails {
		if email.GetVerified() {
			useremail = email.GetEmail()
			break
		}
	}
	if useremail == "" {
		log.Error().Msg("failed to get user details")
		return db.User{}, fmt.Errorf("failed to get user details")
	}
	userdb, err := s.store.GetUserByGithubId(ctx, pgtype.Int8{
		Int64: ghuserdetails.GetID(),
		Valid: true,
	})
	if err != nil {
		if axierr.Is(err, axierr.ErrRecordNotFound) {
			// orgs, _, err := ghuser.Client.Organizations.List(ctx, "", nil)
			// if err != nil {
			// 	log.Error().Err(err).Msg("failed to get orgs")
			// 	return db.User{}, err
			// }
			// if len(orgs) == 0 {
			// 	log.Error().Msg("no orgs found")
			// 	return db.User{}, fmt.Errorf("no orgs found")
			// }
			// orgIndex := -1
			// for i, org := range orgs {
			// 	if verfieidOrgs(org.GetLogin()) {
			// 		orgIndex = i
			// 		break
			// 	}
			// }
			// if orgIndex == -1 {
			// 	log.Error().Msg("org is not verfieid to onboard")
			// 	return db.User{}, fmt.Errorf("org is not verfieid to onboard")
			// }
			org, err := s.store.GetOrgByEntity(ctx, db.GetOrgByEntityParams{
				Name: pgtype.Text{
					String: ghuserdetails.GetLogin(),
					Valid:  true,
				},
			})
			if err != nil {
				if axierr.Is(err, axierr.ErrRecordNotFound) {
					args := db.CreateOrganiationParams{
						Name: ghuserdetails.GetLogin(),
						Domain: pgtype.Text{
							String: util.GetDomain(ghuserdetails.GetEmail()),
							Valid:  true,
						},
						Provider: provider,
					}
					org, err = s.store.CreateOrganiation(ctx, args)
					if err != nil {
						log.Error().Err(err).Msg("failed to create organisation")
						return db.User{}, fmt.Errorf("failed to create organisation")
					}
				} else {
					log.Error().Err(err).Msg("failed to get organisation")
					return db.User{}, fmt.Errorf("failed to get organisation")
				}
			}
			hashPassword, err := util.NewPassword(useremail + ".axilock")
			if err != nil {
				log.Error().Err(err).Msg("failed to hash password")
				return db.User{}, fmt.Errorf("failed to hash password")
			}
			args := db.CreateUserForGithubParams{
				Org: org.ID,
				GithubUserID: pgtype.Int8{
					Int64: ghuserdetails.GetID(),
					Valid: true,
				},
				GithubUserName: pgtype.Text{
					String: ghuserdetails.GetLogin(),
					Valid:  true,
				},
				Email:        useremail,
				HashPassword: hashPassword,
				Provider:     provider,
			}
			_, err = s.store.CreateUserForGithub(ctx, []db.CreateUserForGithubParams{args})
			if err != nil {
				log.Error().Err(err).Msg("failed to create user")
				return db.User{}, fmt.Errorf("failed to create user")
			}
			userdb, err = s.store.GetUserByGithubId(ctx, pgtype.Int8{
				Int64: ghuserdetails.GetID(),
				Valid: true,
			})
			if err != nil {
				log.Error().Err(err).Msg("failed to get user")
				return db.User{}, fmt.Errorf("failed to get user")
			}
			return userdb, nil
		}
		log.Error().Err(err).Msg("failed to get user")
		return db.User{}, fmt.Errorf("failed to get user")
	}
	return userdb, nil
}

// func verfieidOrgs(org string) bool {
// 	return slices.Contains([]string{"axilock"}, org)
// }
