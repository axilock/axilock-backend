package workers

import (
	"github.com/axilock/axilock-backend/internal/service"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func InitProcessor(optn asynq.RedisClientOpt, svc service.Services) {
	taskProcessor := NewRedisTaskProcessor(optn, svc)
	// log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start processor")
	}
}
