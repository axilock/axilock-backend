package clientrpc

import (
	"context"
	"time"

	clientpb "github.com/axilock/axilock-protos/client"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ClientService) CreateAuthSession(ctx context.Context, req *clientpb.CreateAuthSessionRequest) (*clientpb.CreateAuthSessionResponse, error) {
	err := s.Services.Tokensvc.SetCliToken(ctx, req.GetInitToken(), "")
	if err != nil {
		log.Error().Err(err).Msg("failed to set cli session token")
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	ctx2, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Minute))
	defer cancel()
	for {
		select {
		case <-ctx2.Done():
			return nil, status.Errorf(codes.DeadlineExceeded, "deadline exceeded")
		default:
			userId, err := s.Services.Tokensvc.GetSuccessForToken(ctx2, req.GetInitToken())
			if err != nil {
				time.Sleep(500 * time.Millisecond)
				continue
			}
			err = s.Services.Tokensvc.DeleteCliSessionToken(ctx2, req.GetInitToken())
			if err != nil {
				log.Error().Err(err).Msg("failed to delete cli session token")
				return nil, status.Errorf(codes.Internal, "internal error")
			}
			cliAccessToken, cliAccessPayload, err := s.TokenMaker.CreateToken(userId, 0)
			if err != nil {
				log.Error().Err(err).Msg("failed to create cli access token")
				return nil, status.Errorf(codes.Internal, "internal error")
			}
			resp := &clientpb.CreateAuthSessionResponse{
				CliAuthToken:       cliAccessToken,
				CliauthTokenExpiry: timestamppb.New(cliAccessPayload.ExpiredAt),
			}
			return resp, nil
		}
	}
}
