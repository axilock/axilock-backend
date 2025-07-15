package ghevents

import (
	"context"

	"github.com/axilock/axilock-backend/internal/service"
)

type GHMeta struct {
	Access_token string `json:"access_token"`
	Data         []byte `json:"data"`
}

type GHEvents interface {
	ProcessWebhook(ctx context.Context, data []byte) error
	GetTaskID() string
}

type EventClient struct {
	service.Services
}

func NewEventClient(svc service.Services) *EventClient {
	return &EventClient{
		svc,
	}
}

func NewEventDataStore(c *EventClient) map[string]GHEvents {
	eventMap := make(map[string]GHEvents)
	for _, v := range GetGHEventProcessors(c) {
		eventMap[v.GetTaskID()] = v
	}
	return eventMap
}

func GetGHEventProcessors(c *EventClient) []GHEvents {
	return []GHEvents{
		&PrOpen{
			EventClient: c,
		},
		&PrComment{
			EventClient: c,
		},
	}
}
