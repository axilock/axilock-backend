package tokensvc

import (
	"context"
	"fmt"
	"time"

	"github.com/axilock/axilock-backend/internal/db/redis"
	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/rs/zerolog/log"
)

type TokenServiceInterface interface {
	SetCliToken(ctx context.Context, token string, val string) error
	GetSuccessForToken(ctx context.Context, token string) (string, error)
	DeleteCliSessionToken(ctx context.Context, token string) error
	UpdateToken(ctx context.Context, token string, val string) error
	SetClientUpdateCache(ctx context.Context, token string, val string) error
	GetClientUpdateCache(ctx context.Context, token string) (string, error)
	CreateUserDbCache(ctx context.Context, token string, val string) error
	GetUserDbCache(ctx context.Context, token string) (string, error)
}

type TokenService struct {
	store  db.Store
	memory redis.RedisStore
}

func NewTokenService(store db.Store, memory redis.RedisStore) TokenServiceInterface {
	return &TokenService{
		store:  store,
		memory: memory,
	}
}

// func (s *TokenService) CreateTokenForUser(ctx context.Context, req CreateTokenRequest) error {
// 	args := db.CreateTokenForUserParams{
// 		User: req.User,
// 		Org: pgtype.Int8{
// 			Int64: req.Org,
// 			Valid: true,
// 		},
// 		Version:    req.Version,
// 		TokenType:  req.TokenType,
// 		TokenValue: req.TokenValue,
// 	}
// 	err := s.store.CreateTokenForUser(ctx, args)
// 	if err != nil {
// 		return fmt.Errorf("cannot create token")
// 	}
// 	return nil
// }

// func (s *TokenService) GetUserByToken(ctx context.Context, tokenValue string) (db.User, error) {
// 	user, err := s.store.GetUserByToken(ctx, tokenValue)
// 	return user, err
// }

func (s *TokenService) SetCliToken(ctx context.Context, token string, val string) error {
	key := fmt.Sprintf("cli_token:%s", token)
	err := s.memory.SetToken(ctx, key, val, 5*time.Minute)
	if err != nil {
		return fmt.Errorf("cannot create token")
	}
	return nil
}

func (s *TokenService) UpdateToken(ctx context.Context, token string, val string) error {
	key := fmt.Sprintf("cli_token:%s", token)
	v, err := s.memory.GetToken(ctx, key)
	if err != nil {
		return fmt.Errorf("token storage error occoured")
	}
	if v == nil {
		return fmt.Errorf("token not found")
	}
	err = s.memory.SetToken(ctx, key, val, 5*time.Minute)
	if err != nil {
		return fmt.Errorf("cannot create token")
	}
	return nil
}

func (s *TokenService) GetSuccessForToken(ctx context.Context, token string) (string, error) {
	key := fmt.Sprintf("cli_token:%s", token)
	val, err := s.memory.GetToken(ctx, key)
	if err != nil {
		return "", err
	}
	if val.(string) == "" {
		return "", fmt.Errorf("token not found")
	}
	return val.(string), nil
}

func (s *TokenService) DeleteCliSessionToken(ctx context.Context, token string) error {
	key := fmt.Sprintf("cli_token:%s", token)
	err := s.memory.DeleteToken(ctx, key)
	if err != nil {
		return fmt.Errorf("cannot delete token")
	}
	return nil
}

func (s *TokenService) SetClientUpdateCache(ctx context.Context, token string, val string) error {
	key := fmt.Sprintf("client_update:%s", token)
	err := s.memory.SetToken(ctx, key, val, 10*time.Minute)
	if err != nil {
		return fmt.Errorf("cannot create token")
	}
	return nil
}

func (s *TokenService) GetClientUpdateCache(ctx context.Context, token string) (string, error) {
	key := fmt.Sprintf("client_update:%s", token)
	val, err := s.memory.GetToken(ctx, key)
	if err != nil {
		return "", err
	}
	if val == nil {
		return "", nil
	}
	log.Info().Msgf("client update cache hit for token %s", token)
	return val.(string), nil
}

func (s *TokenService) CreateUserDbCache(ctx context.Context, token string, val string) error {
	key := fmt.Sprintf("user_db_cache:%s", token)
	err := s.memory.SetToken(ctx, key, val, 30*time.Minute)
	if err != nil {
		return fmt.Errorf("cannot create token")
	}
	return nil
}

func (s *TokenService) GetUserDbCache(ctx context.Context, token string) (string, error) {
	key := fmt.Sprintf("user_db_cache:%s", token)
	val, err := s.memory.GetToken(ctx, key)
	if err != nil {
		return "", err
	}
	if val == nil {
		return "", nil
	}
	log.Info().Msgf("user db cache hit for token %s", token)
	return val.(string), nil
}
