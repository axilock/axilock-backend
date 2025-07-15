package api

import (
	_ "embed"

	"github.com/axilock/axilock-backend/internal/auth"
	"github.com/axilock/axilock-backend/internal/service"
	"github.com/axilock/axilock-backend/pkg/util"
	"github.com/axilock/axilock-backend/pkg/workers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

// //go:embed static/axi.sh
// var axiSh string

type Server struct {
	router      *fiber.App
	tokenMaker  auth.Maker
	config      util.Config
	services    *service.Services
	distributer workers.TaskDistributerInterface
}

type ServerConfig struct {
	TokenMaker      auth.Maker
	Config          util.Config
	Services        *service.Services
	TaskDistributer workers.TaskDistributerInterface
}

func (s *Server) Start(address string) error {
	return s.router.Listen(address)
}

func NewServer(cfg ServerConfig) (*Server, error) {
	server := &Server{
		services:    cfg.Services,
		tokenMaker:  cfg.TokenMaker,
		config:      cfg.Config,
		distributer: cfg.TaskDistributer,
	}
	server.setupRouter()
	return server, nil
}

func errResponse(err error) fiber.Map {
	return fiber.Map{
		"success": false,
		"error":   err.Error(),
	}
}

func successResponse() fiber.Map {
	return fiber.Map{
		"success": true,
	}
}

func setupMiddlewares(router *fiber.App) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*.axilock.ai", "http://localhost:3000", "https://axilock.ai", "http://*.axilock.ai"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", authheaderkey},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// router.Get("/axi.sh", func(ctx fiber.Ctx) error {
	// 	return ctx.SendString(axiSh)
	// })
}

func (s *Server) setupRouter() {
	router := fiber.New(
		fiber.Config{
			StructValidator: &structValidator{validate: validator.New()},
		},
	)
	router.Use(recover.New())
	router.Use(DefaultStructuredLogger())
	setupMiddlewares(router)
	v1router := router.Group("/v1")
	// v1router.Post("/user/create", s.CreateUser)
	v1router.Post("/user/login", s.LoginUser)
	v1router.Post("/github/webhook", s.WebhookGithub)
	v1router.Get("/user/auth/github-url", s.GithubLoginURL)
	v1router.Post("/user/auth/github", s.GithubLogin)
	v1router.Post("/auth/cli-auth", s.InitCliAuth)
	v1router.Post("/client/update", s.UpdateClient)
	v1router.Post("/inbound", s.Inbound)

	authrouter := v1router.Group("/").Use(authMiddleWare(s.tokenMaker))
	{
		authrouter.Get("/user/details", s.GetUserDetails)
		authrouter.Post("/github-app/callback", s.GithubCallback)
		authrouter.Get("/integrations/all", s.GetIntegrations)
		authrouter.Get("/alerts/all", s.GetAlertStats)
		authrouter.Get("/alerts/repo", s.GetTop10RepoAlerts)
		authrouter.Get("/alerts/weekly", s.GetWeeklyStats)
		authrouter.Get("/alerts/protected/graph", s.GetProtectedSecretsGraphData)
		authrouter.Get("/repo/repostats", s.GetRepoStats)
		authrouter.Get("/commits/health", s.GetCommitsHealth)
		authrouter.Get("/alerts/secret/type", s.GetAlertSecretTypeCount)
		authrouter.Get("/users/coverage", s.GetUserCoverage)
	}
	router.Use(func(ctx fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusNotFound)
	})
	s.router = router
}
