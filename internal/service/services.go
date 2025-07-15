package service

import (
	"github.com/axilock/axilock-backend/internal/db/redis"
	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/axilock/axilock-backend/internal/service/alertsvc"
	"github.com/axilock/axilock-backend/internal/service/commitsvc"
	"github.com/axilock/axilock-backend/internal/service/githubsvc"
	"github.com/axilock/axilock-backend/internal/service/metadata"
	"github.com/axilock/axilock-backend/internal/service/orgsvc"
	"github.com/axilock/axilock-backend/internal/service/reposvc"
	"github.com/axilock/axilock-backend/internal/service/tokensvc"

	"github.com/axilock/axilock-backend/internal/service/usersvc"
	"github.com/axilock/axilock-backend/pkg/util"
)

type Services struct {
	Usersvc   usersvc.UserServiceInterface
	Orgsvc    orgsvc.OrgServiceInterface
	Tokensvc  tokensvc.TokenServiceInterface
	Metasvc   metadata.MetaDataServiceInterface
	CommitSvc commitsvc.CommitServiceInterface
	GithubSvc githubsvc.GithubService
	AlertSvc  alertsvc.AlertServiceInterface
	RepoSvc   reposvc.RepoService
}

func NewServiceInit(store db.Store, config util.Config, redis redis.RedisStore) *Services {
	usersvc := usersvc.NewUserService(store, config)
	orgsvc := orgsvc.NewOrgService(store)
	tokensvc := tokensvc.NewTokenService(store, redis)
	metasvc := metadata.NewMetaDataService(store)
	commitsvc := commitsvc.NewCommitService(store)
	githubsvc := githubsvc.NewGithubService(store, config)
	alertsvc := alertsvc.NewAlertService(store)
	reposvc := reposvc.NewRepoService(store)

	return &Services{
		Usersvc:   usersvc,
		Orgsvc:    orgsvc,
		Tokensvc:  tokensvc,
		Metasvc:   metasvc,
		CommitSvc: commitsvc,
		GithubSvc: githubsvc,
		AlertSvc:  alertsvc,
		RepoSvc:   reposvc,
	}
}
