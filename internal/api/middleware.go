package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/axilock/axilock-backend/internal/auth"
	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/gofiber/fiber/v3"
)

const (
	authheaderkey           string = "authorization"
	authorizationTypeBearer string = "bearer"
	authorizationPayloadKey string = "authPayload"
)

func (s *Server) getUserFromContext(ctx fiber.Ctx) (db.User, error) {
	authpayload := ctx.Locals(authorizationPayloadKey).(*auth.Payload)

	cache, _ := s.services.Tokensvc.GetUserDbCache(ctx.Context(), authpayload.ID.String())
	if cache != "" {
		user := db.User{}
		json.Unmarshal([]byte(cache), &user)
		return user, nil
	}
	user, err := s.services.Usersvc.GetUserByID(ctx.Context(), authpayload.UserID)
	if err != nil {
		return db.User{}, fmt.Errorf("cannot get user")
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		return db.User{}, fmt.Errorf("cannot marshal user")
	}
	err = s.services.Tokensvc.CreateUserDbCache(ctx.Context(), authpayload.ID.String(), string(userJson))
	if err != nil {
		return db.User{}, fmt.Errorf("cannot create user db cache")
	}
	return user, nil
}

func authMiddleWare(tokenMaker auth.Maker) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		authHeaderVal := ctx.Get(authheaderkey)

		if len(authHeaderVal) == 0 {
			err := errors.New("authorization not provided")
			return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
		}

		fields := strings.Fields(authHeaderVal)
		if len(fields) < 2 {
			err := errors.New("invalid authorization format")
			return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))

		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("bearer type authorization only supported")
			return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))

		}
		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
		}
		ctx.Locals(authorizationPayloadKey, payload)
		return ctx.Next()
	}
}
