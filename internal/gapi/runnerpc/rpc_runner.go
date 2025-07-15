package runnerpc

import (
	"context"

	runnerpb "github.com/axilock/axilock-protos/runner"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *RunnerService) StartRepoScan(context.Context, *runnerpb.RunnerScanRequest) (*runnerpb.RunnerScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
