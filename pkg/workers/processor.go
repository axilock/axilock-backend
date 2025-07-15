package workers

import (
	"context"

	"github.com/axilock/axilock-backend/internal/service"
	"github.com/axilock/axilock-backend/pkg/workers/ghevents"
	"github.com/hibiken/asynq"
)

type TaskProcessorInterface interface {
	Start() error
	ProcessGithubTask(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server    *asynq.Server
	datastore map[string]ghevents.GHEvents
}

func NewRedisTaskProcessor(redisOptn asynq.RedisClientOpt, svc service.Services) TaskProcessorInterface {
	server := asynq.NewServer(
		redisOptn,
		asynq.Config{},
	)
	store := ghevents.NewEventDataStore(ghevents.NewEventClient(svc))
	return &RedisTaskProcessor{
		server:    server,
		datastore: store,
	}
}

func (p *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	for _, v := range p.datastore {
		mux.HandleFunc(v.GetTaskID(), p.ProcessGithubTask)
	}
	return p.server.Start(mux)
}
