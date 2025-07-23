package api

import (
	"encoding/json"
	"time"

	"github.com/axilock/axilock-backend/internal/axierr"
	"github.com/axilock/axilock-backend/pkg/gh"
	"github.com/gofiber/fiber/v3"
	"github.com/google/go-github/v72/github"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func (s *Server) WebhookGithub(ctx fiber.Ctx) error {
	req := ctx.Request()
	defer req.CloseBodyStream()
	eventHdr := req.Header.Peek("x-github-event")
	if string(eventHdr) == "" {
		return nil
	}
	body, err := req.BodyUncompressed()
	if err != nil {
		log.Info().Msg("failed to get body")
		return nil
	}
	var ins github.InstallationEvent
	err = json.Unmarshal(body, &ins)
	if err != nil {
		log.Info().Msg("failed to parse install event")
		return nil
	}
	acc_token, err := s.services.GithubSvc.GetINstallationAccessToken(ctx.Context(), ins.Installation.GetID())
	if err != nil {
		log.Info().Msg("failed to get instlaation access token")
		return nil
	}
	if err := s.distributer.DistributeGithubTask(ctx.Context(), body, string(eventHdr), acc_token,
		asynq.Timeout(10*time.Minute), asynq.MaxRetry(1)); err != nil {
		log.Info().Msg("failed to queue github task")
	}
	return ctx.JSON(successResponse())
}

type GithubCallbackResp struct {
	Code           string `json:"code"`
	InstallationID string `json:"installation_id"`
	Action         string `json:"setup_action"`
	State          string `json:"state"`
}

func (s *Server) GithubCallback(ctx fiber.Ctx) error {
	var req GithubCallbackResp
	if err := ctx.Bind().Body(&req); err != nil {
		log.Err(err).Str("error", err.Error()).Msg("recieved error")
		return ctx.Status(fiber.StatusBadRequest).JSON(errResponse(err))
	}
	user, err := s.getUserFromContext(ctx)
	if err != nil {
		log.Err(err).Str("error", err.Error()).Msg("recieved error")
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
	}
	_, err = s.services.GithubSvc.GetGithubOrgByOrgID(ctx.Context(), user.Org)
	if err != nil {
		if axierr.Is(err, axierr.ErrRecordNotFound) {
			err = s.services.GithubSvc.CreateGithubInstallation(ctx.Context(), req.InstallationID, user.Org)
			if err != nil {
				log.Err(err).Str("error", err.Error()).Msg("recieved error")
				return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
			}
			return ctx.JSON(successResponse())
		}
		log.Err(err).Str("error", err.Error()).Msg("recieved error")
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	return ctx.JSON(successResponse())
}

func (s *Server) GithubLoginURL(ctx fiber.Ctx) error {
	return ctx.Redirect().To(gh.GetGithubAuthEndpoint(s.config))
}

// GithubLogin returns the github login url
// @Summary Github Login
// @Tags Github
// @Produce json
// @Success 200 {object} string
// @Router /api/v1/github/login [post]
func (s *Server) GithubLogin(ctx fiber.Ctx) error {
	var req GithubCallbackResp
	if err := ctx.Bind().Body(&req); err != nil {
		log.Err(err).Str("error", err.Error()).Msg("recieved error")
		return ctx.Status(fiber.StatusBadRequest).JSON(errResponse(err))
	}
	user, err := s.services.Usersvc.CreateUserWithGithub(ctx.Context(), req.Code, "github")
	if err != nil {
		log.Err(err).Str("error", err.Error()).Msg("recieved error")
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	accessToken, _, err := s.tokenMaker.CreateToken(user.Uuid, s.config.AccessTokenDuration)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	resp := loginUserResponse{
		AccessToken: accessToken,
	}
	return ctx.JSON(resp)
}
