package workers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/axilock/axilock-backend/pkg/workers/ghevents"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func (d *RedisTaskDistributer) DistributeGithubTask(ctx context.Context,
	payload []byte,
	eventType string,
	access_token string,
	opts ...asynq.Option,
) error {
	taskID := fmt.Sprintf("task:github:%s", eventType)
	body, err := json.Marshal(ghevents.GHMeta{
		Access_token: access_token,
		Data:         payload,
	})
	if err != nil {
		return fmt.Errorf("failed to create meta task")
	}
	task := asynq.NewTask(taskID, body, opts...)
	taskInfo, err := d.client.EnqueueContext(ctx, task, opts...)
	if err != nil {
		return fmt.Errorf("failed to enqueue task")
	}
	log.Info().Str("enqued task", taskInfo.ID).Str("type", taskInfo.Type).Msg("enqueued task")
	return nil
}

func (p *RedisTaskProcessor) ProcessGithubTask(ctx context.Context, task *asynq.Task) error {
	taskID := task.Type()
	taskType := p.datastore[taskID]
	if taskType == nil {
		return fmt.Errorf("cannot find task name %s in map", taskID)
	}
	log.Info().Str("processing task", task.ResultWriter().
		TaskID()).Str("type", taskID).Msg("processing")
	err := taskType.ProcessWebhook(ctx, task.Payload())
	if err != nil {
		return fmt.Errorf("cannot process error")
	}
	return nil
}
