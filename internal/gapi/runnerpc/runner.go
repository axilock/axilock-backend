package runnerpc

import (
	"github.com/axilock/axilock-backend/internal/gapi"
	runnerpb "github.com/axilock/axilock-protos/runner"
)

type RunnerService struct {
	*gapi.Server
	runnerpb.UnimplementedScanServiceServer
}
