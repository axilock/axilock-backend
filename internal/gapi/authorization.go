package gapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/axilock/axilock-backend/internal/auth"
	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	auhtorizationHeader string = "authorization"
	authorizationScheme string = "bearer"
)

func (s *Server) AuthorizeUser(ctx context.Context) (db.User, *auth.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return db.User{}, nil, fmt.Errorf("missing metadata context")
	}
	values := md.Get(auhtorizationHeader)
	if len(values) == 0 {
		return db.User{}, nil, fmt.Errorf("missing header")
	}
	authHedear := values[0]

	fields := strings.Fields(authHedear)

	if len(fields) != 2 {
		return db.User{}, nil, fmt.Errorf("missing header")
	}
	authType := strings.ToLower(fields[0])

	if authType != authorizationScheme {
		return db.User{}, nil, fmt.Errorf("unsupported auth scheme")
	}

	accessToken := fields[1]
	payload, err := s.TokenMaker.VerifyToken(accessToken)
	if err != nil {
		return db.User{}, nil, fmt.Errorf("unauntheticated user")
	}

	cache, _ := s.Services.Tokensvc.GetUserDbCache(ctx, payload.ID.String())
	if cache != "" {
		user := db.User{}
		json.Unmarshal([]byte(cache), &user)
		return user, payload, nil
	}

	user, err := s.Services.Usersvc.GetUserByID(ctx, payload.UserID)
	if err != nil {
		return db.User{}, nil, status.Errorf(codes.Internal, "cannot find user")
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		return db.User{}, nil, fmt.Errorf("cannot marshal user")
	}
	err = s.Services.Tokensvc.CreateUserDbCache(ctx, payload.ID.String(), string(userJson))
	if err != nil {
		return db.User{}, nil, fmt.Errorf("cannot create user db cache")
	}
	return user, payload, nil
}
