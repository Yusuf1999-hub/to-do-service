package grpc_client

import (
	"github.com/Yusuf1999-hub/to-do-service/config"
)

// IGrpcClient ...
type IGrpcClient interface{}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	return &GrpcClient{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}
