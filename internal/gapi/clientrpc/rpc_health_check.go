package clientrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	clientpb "github.com/axilock/axilock-protos/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type datastore struct {
	backendVer      string
	latestClientVer string
	systemstatus    string
}

func (s *ClientService) DoHealthCheck(_ context.Context, _ *clientpb.HealthRequest) (*clientpb.HealthResponse, error) {
	data := datastore{
		backendVer:      "0.2.3",
		latestClientVer: "0.12.3",
		systemstatus:    "ok",
	}
	resp := &clientpb.HealthResponse{
		BackendVer:      data.backendVer,
		Systemstatus:    data.systemstatus,
		LatestClientver: data.latestClientVer,
		CreatedAt:       timestamppb.Now(),
	}
	return resp, nil
}

type ApiResponse struct {
	Version   string `json:"version"`
	BinaryUrl string `json:"binary_url"`
}

func (s *ClientService) ClientUpdateRpc(ctx context.Context, req *clientpb.ClientUpdateRequest) (*clientpb.ClientUpdateResponse, error) {
	token, _ := s.Services.Tokensvc.GetClientUpdateCache(ctx, fmt.Sprintf("%x", fmt.Appendf(nil, "%s-%s-%s", req.Os.String(), req.Arch.String(), req.Environment.String())))
	if token != "" {
		var apiresp ApiResponse
		err := json.Unmarshal([]byte(token), &apiresp)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot get updated client ver")
		}
		if req.GetClientVer() == apiresp.Version {
			return &clientpb.ClientUpdateResponse{
				ToUpdate: false,
			}, nil
		}
		return &clientpb.ClientUpdateResponse{
			ToUpdate:         true,
			LatestClientver:  apiresp.Version,
			ClientUpdatePath: apiresp.BinaryUrl,
		}, nil
	}

	apiGwUrl := fmt.Sprintf("https://6ep6wyubv1.execute-api.ap-south-1.amazonaws.com/default/get-latestbinary-from-s3?os=%s&arch=%s&env=%s",
		req.Os.String(), req.Arch.String(), req.Environment.String())
	resp, err := http.Get(apiGwUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot exec lambda")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.Internal, "status code not 200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot read body")
	}
	var apiresp ApiResponse
	err = json.Unmarshal(body, &apiresp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get updated client ver")
	}
	_ = s.Services.Tokensvc.SetClientUpdateCache(ctx, fmt.Sprintf("%x", fmt.Appendf(nil, "%s-%s-%s", req.Os.String(), req.Arch.String(), req.Environment.String())), string(body))
	response := &clientpb.ClientUpdateResponse{}
	if req.GetClientVer() == apiresp.Version {
		response.ToUpdate = false
		return response, nil
	}
	response.ToUpdate = true
	response.LatestClientver = apiresp.Version
	response.ClientUpdatePath = apiresp.BinaryUrl
	return response, nil
}
