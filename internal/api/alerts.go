package api

import (
	"github.com/gofiber/fiber/v3"
)

type getAlertsRequest struct {
	Count int32  `json:"count"`
	State string `json:"state"`
}

type getAlertsResponse struct {
	Alerts []alertdata `json:"alerts"`
}

type alertdata struct {
	AlertName  string `json:"secret_type"`
	State      string `json:"state"`
	Source     string `json:"source"`
	SourceRepo string `json:"source_repo"`
	Date       string `json:"date"`
	Path       string `json:"path"`
}

func (s *Server) GetAlertStats(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(successResponse())
	// var req getAlertsRequest
	// if err := ctx.Bind().Query(&req); err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(errResponse(err))
	// }
	// user, _ := s.getUserFromContext(ctx)
	// if req.Count == 0 {
	// 	req.Count = 10
	// }
	// if req.State == "" {
	// 	req.State = alertsvc.ALERT_OPEN
	// }
	// alerts, err := s.services.AlertSvc.GetAlertData(ctx.Context(), user.Org, req.Count, req.State)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	// }
	// respdata := make([]alertdata, 0, len(alerts))
	// for _, alert := range alerts {
	// 	respdata = append(respdata, alertdata{
	// 		AlertName:  alert.Name,
	// 		State:      alert.Status,
	// 		Source:     alert.Source.String,
	// 		Date:       alert.CreatedAt.Time.Format(time.DateTime),
	// 		Path:       alert.FilePath,
	// 		SourceRepo: "",
	// 	})
	// }
	// resp := getAlertsResponse{
	// 	Alerts: respdata,
	// }
	// return ctx.JSON(resp)
}

type aletReponse struct {
	Alerts any `json:"data"`
}

func (s *Server) GetTop10RepoAlerts(ctx fiber.Ctx) error {
	user, err := s.getUserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
	}
	data, err := s.services.AlertSvc.GetTop10Repo(ctx.Context(), user.Org)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	resp := aletReponse{
		Alerts: data,
	}
	return ctx.JSON(resp)
}

func (s *Server) GetWeeklyStats(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(successResponse())
	// user, err := s.getUserFromContext(ctx)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
	// }
	// data, err := s.services.AlertSvc.GetWeeklyStats(ctx.Context(), user.Org)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	// }
	// resp := aletReponse{
	// 	Alerts: data,
	// }
	// return ctx.JSON(resp)
}

func (s *Server) GetRepoStats(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(successResponse())
	// user, err := s.getUserFromContext(ctx)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
	// }
	// data, err := s.services.RepoSvc.GetRepoStats(ctx.Context(), user.Org)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	// }
	// resp := aletReponse{
	// 	Alerts: data,
	// }
	// return ctx.JSON(resp)
}

func (s *Server) GetProtectedSecretsGraphData(ctx fiber.Ctx) error {
	user, err := s.getUserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
	}
	data, err := s.services.AlertSvc.GetProtectedSecretsOverTime(ctx.Context(), user.Org)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	resp := aletReponse{
		Alerts: data,
	}
	return ctx.JSON(resp)
}

type commitResponse struct {
	Data any `json:"data"`
}

func (s *Server) GetCommitsHealth(ctx fiber.Ctx) error {
	user, err := s.getUserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
	}
	data, err := s.services.AlertSvc.GetCommitsHealth(ctx.Context(), user.Org)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	resp := commitResponse{
		Data: data,
	}
	return ctx.JSON(resp)
}

func (s *Server) GetAlertSecretTypeCount(ctx fiber.Ctx) error {
	user, err := s.getUserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
	}
	data, err := s.services.AlertSvc.GetAlertSecretTypeCount(ctx.Context(), user.Org)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	resp := aletReponse{
		Alerts: data,
	}
	return ctx.JSON(resp)
}
