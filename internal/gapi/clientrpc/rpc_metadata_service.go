package clientrpc

import (
	"context"

	"github.com/axilock/axilock-backend/internal/axierr"

	"github.com/axilock/axilock-backend/internal/service/metadata"
	clientpb "github.com/axilock/axilock-protos/client"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ClientService) InitMetadata(ctx context.Context, req *clientpb.InstallerInitRequest) (*clientpb.InstallerInitResponse, error) {
	user, _, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, axierr.UnauthenticatedError(err)
	}

	args := metadata.CreatetMetaDataReq{
		UserID:   user.ID,
		OrgID:    user.Org,
		MetaType: req.GetStatus().String(),
		MetaData: req.GetMetadata(),
	}
	err = s.Services.Metasvc.CreateMetadata(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create metadata")
	}
	resp := &clientpb.InstallerInitResponse{
		Status:    true,
		CreatedAt: timestamppb.Now(),
	}
	return resp, nil
}

func (s *ClientService) RepoMetadata(ctx context.Context, req *clientpb.MetadataRepoRequest) (*clientpb.MetadataRepoResponse, error) {
	user, _, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, axierr.UnauthenticatedError(err)
	}
	args := metadata.CreatetMetaDataReq{
		UserID:   user.ID,
		OrgID:    user.Org,
		MetaType: metadata.RepoMeta,
		MetaData: req.GetMetadata(),
	}
	err = s.Services.Metasvc.CreateMetadata(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create metadata%s", err)
	}
	resp := &clientpb.MetadataRepoResponse{
		Status:    true,
		CreatedAt: timestamppb.Now(),
	}
	return resp, nil
}
