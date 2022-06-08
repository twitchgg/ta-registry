package server

import (
	"context"

	"ntsc.ac.cn/ta-registry/pkg/pb"
)

// RegistRouter regist wireguard router
func (s *RegistryServer) RegistRouter(context.Context, *pb.RegistRouterRequest) (*pb.RouterConfig, error) {
	return &pb.RouterConfig{}, nil
}
