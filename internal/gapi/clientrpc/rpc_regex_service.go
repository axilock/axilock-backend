package clientrpc

import (
	"context"

	clientpb "github.com/axilock/axilock-protos/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ClientService) CreateRegex(ctx context.Context, req *clientpb.CreateRegexRequest) (*clientpb.CreateRegexResponse, error) {
	// authPayload, err := s.AuthorizeUser(ctx)
	// if err != nil {
	// 	return nil, axierr.UnauthenticatedError(err)
	// }
	// user, err := s.Services.Usersvc.GetUserByID(ctx, authPayload.UserID)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "user not mapped to org")
	// }
	// args := orgsvc.CreateRegexReq{
	// 	RegexStr:  req.GetRegexStr(),
	// 	RegexType: req.GetType(),
	// 	Desc:      req.GetDesc(),
	// 	Org:       user.Org,
	// }
	// err = s.Services.Orgsvc.CreateRegexForOrg(ctx, args)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "failed to store regex")
	// }
	// resp := &clientpb.CreateRegexResponse{
	// 	Success:   true,
	// 	CreatedAt: timestamppb.Now(),
	// }
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegex not implemented")
}

func (s *ClientService) SyncRegex(ctx context.Context, _ *clientpb.SyncRegexDataRequest) (*clientpb.SyncRegexDataResponse, error) {
	// authPayload, err := s.AuthorizeUser(ctx)
	// if err != nil {
	// 	return nil, axierr.UnauthenticatedError(err)
	// }
	// user, err := s.Services.Usersvc.GetUserByID(ctx, authPayload.UserID)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "user not mapped to org")
	// }
	// regexs, err := s.Services.Orgsvc.GetRegexForOrg(ctx, user.Org)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "failed to get regex %s", err)
	// }
	// resp := &clientpb.SyncRegexDataResponse{
	// 	IsChanged: true,
	// 	Items:     nil,
	// }
	// for _, val := range regexs {
	// 	regex := clientpb.RegexData{
	// 		Desc:     val.Description.String,
	// 		RegexStr: val.Regstring,
	// 		Version:  int32(val.Version),
	// 	}
	// 	resp.Items = append(resp.Items, &regex)
	// }
	return nil, status.Errorf(codes.Unimplemented, "method SyncRegex not implemented")
}
