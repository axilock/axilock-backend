package clientrpc

import (
	"github.com/axilock/axilock-backend/internal/gapi"
	clientpb "github.com/axilock/axilock-protos/client"
)

type ClientService struct {
	*gapi.Server
	clientpb.UnimplementedSessionServiceServer
	clientpb.UnimplementedHealthServiceServer
	clientpb.UnimplementedMetadataServiceServer
	clientpb.UnimplementedRegexServiceServer
	clientpb.UnimplementedCommitDataServiceServer
	clientpb.UnimplementedAlertServiceServer
}
