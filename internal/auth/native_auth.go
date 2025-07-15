package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/axilock/axilock-backend/internal/db/redis"
	"github.com/axilock/axilock-backend/pkg/util"
)

type NativeAuth struct {
	store redis.RedisStore
}

func NewNativeAuth(store redis.RedisStore) Maker {
	return NativeAuth{
		store: store,
	}
}

func (a NativeAuth) CreateToken(userID string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", nil, err
	}
	val, err := json.Marshal(payload)
	if err != nil {
		return "", nil, err
	}
	acc_token := util.RandomUUID()
	key := fmt.Sprintf("access_token:%s", acc_token)
	err = a.store.SetToken(context.Background(), key, val, duration)
	if err != nil {
		return "", nil, err
	}
	return acc_token, payload, nil
}

func (a NativeAuth) VerifyToken(token string) (*Payload, error) {
	key := fmt.Sprintf("access_token:%s", token)
	val, err := a.store.GetToken(context.Background(), key)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, fmt.Errorf("invalid token")
	}
	var payload Payload
	err = json.Unmarshal([]byte(val.(string)), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
