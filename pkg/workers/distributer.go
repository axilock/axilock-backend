package workers

import (
	"context"

	"github.com/axilock/axilock-backend/internal/service"
	"github.com/axilock/axilock-backend/pkg/workers/ghevents"
	"github.com/hibiken/asynq"
)

type TaskDistributerInterface interface {
	DistributeGithubTask(ctx context.Context,
		payload []byte,
		eventType string,
		access_token string,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributer struct {
	client    *asynq.Client
	datastore map[string]ghevents.GHEvents
}

func NewRedisTaskDistributer(redisOptn asynq.RedisClientOpt, svc service.Services) TaskDistributerInterface {
	client := asynq.NewClient(redisOptn)
	store := ghevents.NewEventDataStore(ghevents.NewEventClient(svc))
	return &RedisTaskDistributer{
		client:    client,
		datastore: store,
	}
}
