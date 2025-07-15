package clientrpc

import (
	"context"

	"github.com/axilock/axilock-backend/internal/axierr"
	"github.com/axilock/axilock-backend/internal/constants"
	"github.com/axilock/axilock-backend/internal/service/alertsvc"
	clientpb "github.com/axilock/axilock-protos/client"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ClientService) SecretAlert(ctx context.Context, req *clientpb.SecretAlertRequest) (*clientpb.SecretAlertResponse, error) {
	user, _, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, axierr.UnauthenticatedError(err)
	}
	args := alertsvc.CreateSecretAlertReq{
		Filename:   req.GetFileName(),
		Repo:       req.GetRepo(),
		Org:        user.Org,
		Lineno:     req.GetLineNumber(),
		Commitid:   req.GetCommitId(),
		Source:     constants.SOURCE_AXI_CLI,
		Secrettype: req.GetSecretType(),
		IsVerified: req.GetIsVerified(),
		Filepath:   "testtest",
	}
	err = s.Services.AlertSvc.CreateSecretAlert(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create alert for user %v", err)
	}
	return &clientpb.SecretAlertResponse{
		Success: true,
	}, nil
}
