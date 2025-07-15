package clientrpc

import (
	"context"
	"time"

	"github.com/axilock/axilock-backend/internal/axierr"

	"github.com/axilock/axilock-backend/internal/service/commitsvc"
	clientpb "github.com/axilock/axilock-protos/client"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ClientService) SendCommitData(ctx context.Context, req *clientpb.SendCommitDataRequest) (*clientpb.SendCommitDataResponse, error) {
	user, payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, axierr.UnauthenticatedError(err)
	}
	dbArr := make([]commitsvc.CreateCommmitSvcDBParamas, 0, len(req.GetCommits()))
	for _, val := range req.GetCommits() {
		dbArr = append(dbArr, commitsvc.CreateCommmitSvcDBParamas{
			CommitID:   val.GetCommitId(),
			AuthorName: val.GetCommitAuthor(),
			CommitTime: val.GetCommitTime().AsTime().Format(time.DateTime),
			PushTime:   req.GetPushTime().AsTime().Format(time.DateTime),
		})
	}
	args := commitsvc.CreateCommitGrpcReq{
		RepoURL:    req.GetRepoUrl(),
		CommitData: dbArr,
		Org:        user.Org,
		UserID:     user.ID,
	}

	err = s.Services.CommitSvc.CreateCommitsForSession(ctx, args, payload.ID.String())
	if err != nil {
		if axierr.IsUniqueViolation(err) {
			return nil, status.Errorf(codes.AlreadyExists, "commit already exists")
		}
		return nil, status.Errorf(codes.Internal, "cannot add commits %s", err)
	}
	resp := &clientpb.SendCommitDataResponse{
		Status:    true,
		CreatedAt: timestamppb.Now(),
	}
	return resp, nil
}
