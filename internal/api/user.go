package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axilock/axilock-backend/internal/axierr"
	"github.com/axilock/axilock-backend/pkg/util"
	"github.com/gofiber/fiber/v3"
)

type createUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type createUserReponse struct {
	Status bool `json:"success"`
}

// func (s *Server) CreateUser(ctx fiber.Ctx) error {
// 	var req createUserRequest
// 	if err := ctx.Bind().Body(&req); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(errResponse(err))
// 	}
// 	usersvcreq := usersvc.CreateUserRequest{
// 		Email:    req.Email,
// 		Password: req.Password,
// 	}
// 	_, _, err := s.services.Usersvc.CreateUser(ctx.Context(), usersvcreq)
// 	if err != nil {
// 		if axierr.ErrorCode(err) == axierr.UniqueViolation {
// 			return ctx.Status(fiber.StatusForbidden).JSON(errResponse(err))
// 		}
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
// 	}
// 	resp := createUserReponse{
// 		Status: true,
// 	}
// 	return ctx.JSON(resp)
// }

type loginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string `json:"accesstoken"`
}

func (s *Server) LoginUser(ctx fiber.Ctx) error {
	var req loginUserRequest

	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errResponse(err))
	}
	user, err := s.services.Usersvc.GetUserByEmail(ctx.Context(), req.Email)
	if err != nil {
		if axierr.Is(err, axierr.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(errResponse(err))
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	err = util.CheckPassword(user.HashPassword, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
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

type orgstats struct {
	Repocount          int64  `json:"repocount"`
	OrgName            string `json:"org_name"`
	CustomPatternCount int64  `json:"custom_pattern_count"`
	IsGithubOnboarded  bool   `json:"is_github_onboarded"`
}

type getUserResponse struct {
	Username string   `json:"username"`
	Orgstats orgstats `json:"orgstats"`
}

func (s *Server) GetUserDetails(ctx fiber.Ctx) error {
	// var req getUserRequest
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errReponse(err))
	// 	return
	// }
	user, err := s.getUserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	org, err := s.services.Orgsvc.GetOrgByID(ctx.Context(), user.Org)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	// regex, err := s.services.Orgsvc.GetRegexForOrg(ctx.Context(), user.Org)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	// }
	reposcount, err := s.services.Orgsvc.GetReposForOrg(ctx.Context(), org.ID)
	if err != nil && !axierr.Is(err, axierr.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	isGithubOnboarded := true
	githubInstallation, err := s.services.GithubSvc.GetGithubOrgByOrgID(ctx.Context(), org.ID)
	if err != nil && axierr.Is(err, axierr.ErrRecordNotFound) {
		isGithubOnboarded = false
		resp := getUserResponse{
			Username: user.GithubUserName.String,
			Orgstats: orgstats{
				OrgName:            org.Name,
				Repocount:          reposcount,
				CustomPatternCount: int64(0),
				IsGithubOnboarded:  isGithubOnboarded,
			},
		}
		return ctx.JSON(resp)
	}
	resp := getUserResponse{
		Username: user.GithubUserName.String,
		Orgstats: orgstats{
			OrgName:            githubInstallation.OrgName,
			Repocount:          reposcount,
			CustomPatternCount: int64(0),
			IsGithubOnboarded:  true,
		},
	}
	return ctx.JSON(resp)
}

type cliAuthRequest struct {
	Clitoken string `json:"clitoken" validate:"required"`
	Provider string `json:"provider" validate:"required"`
	Code     string `json:"code" validate:"required"`
}

func (s *Server) InitCliAuth(ctx fiber.Ctx) error {
	var req cliAuthRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errResponse(err))
	}
	user, err := s.services.Usersvc.CreateUserWithGithub(ctx.Context(), req.Code, req.Provider)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	err = s.services.Tokensvc.UpdateToken(ctx.Context(), req.Clitoken, user.Uuid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	return ctx.JSON(successResponse())
}

type updateReq struct {
	Environment string `json:"environment"`
	Os          string `json:"os"`
	Arch        string `json:"arch"`
	ClientVer   string `json:"client_ver"`
}

type updateApiResp struct {
	Version   string `json:"version"`
	BinaryUrl string `json:"binary_url"`
}

func (s *Server) UpdateClient(ctx fiber.Ctx) error {
	var req updateReq
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errResponse(err))
	}
	apiGwUrl := fmt.Sprintf("https://6ep6wyubv1.execute-api.ap-south-1.amazonaws.com/default/get-latestbinary-from-s3?os=%s&arch=%s&env=%s",
		req.Os, req.Arch, req.Environment)
	resp, err := http.Get(apiGwUrl)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	var apiresp updateApiResp
	err = json.Unmarshal(body, &apiresp)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	return ctx.Redirect().To(apiresp.BinaryUrl)
}

type inboundReq struct {
	Email string `json:"email" validate:"required,email"`
}

type DiscordWebhook struct {
	Content string `json:"content"`
}

func (s *Server) Inbound(ctx fiber.Ctx) error {
	var req inboundReq
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errResponse(err))
	}
	webhookData := DiscordWebhook{
		Content: fmt.Sprintf("New user intent to onboard %s", req.Email),
	}
	jsonBody, err := json.Marshal(webhookData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	resp, err := http.Post(s.config.DiscordWebhook, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	defer resp.Body.Close()
	return ctx.Redirect().To("https://cal.com/axilock/support?email=" + req.Email)
}

func (s *Server) GetUserCoverage(ctx fiber.Ctx) error {
	user, err := s.getUserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse(err))
	}
	data, err := s.services.GithubSvc.GetUserCoverage(ctx.Context(), user.Org)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
	resp := aletReponse{
		Alerts: data,
	}
	return ctx.JSON(resp)
}
