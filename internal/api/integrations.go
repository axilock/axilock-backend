package api

import (
	"github.com/axilock/axilock-backend/internal/axierr"
	"github.com/gofiber/fiber/v3"
)

type integrations struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	ActionURL string `json:"action_url"`
}

type getIntegrationsResponse struct {
	Integrations []integrations `json:"integrations"`
}

func (s *Server) GetIntegrations(ctx fiber.Ctx) error {
	user, err := s.getUserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	_, err = s.services.GithubSvc.GetGithubOrgByOrgID(ctx.Context(), user.Org)
	if err != nil {
		if axierr.Is(err, axierr.ErrRecordNotFound) {
			resp := getIntegrationsResponse{
				Integrations: []integrations{
					{
						Name: "github", Status: "inactive",
						ActionURL: "https://github.com/apps/axilock-runner/installations/new",
					},
				},
			}

			return ctx.JSON(resp)
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	resp := getIntegrationsResponse{
		Integrations: []integrations{
			{Name: "github", Status: "active", ActionURL: ""},
		},
	}
	return ctx.JSON(resp)
}
