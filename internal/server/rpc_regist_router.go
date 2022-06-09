package server

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"ntsc.ac.cn/ta-registry/pkg/pb"
	"ntsc.ac.cn/ta-registry/pkg/rpc"
)

const (
	CERT_EXT_KEY_MACHINE_ID = "1.1.1.1.1.1"
)

// RegistRouter regist wireguard router
func (s *RegistryServer) RegistRouter(ctx context.Context,
	req *pb.RegistRouterRequest) (*pb.RouterConfig, error) {
	machineID, err := s.getCertExtValue(ctx, CERT_EXT_KEY_MACHINE_ID)
	if err != nil {
		return nil, rpc.GenerateError(codes.InvalidArgument,
			fmt.Errorf("not found machine id from certificate"))
	}
	if machineID != req.MachineID {
		return nil, rpc.GenerateError(codes.InvalidArgument,
			fmt.Errorf("machine id does not match"))
	}
	if logrus.IsLevelEnabled(logrus.TraceLevel) {
		return s._routerConfigMock(), nil
	}
	return &pb.RouterConfig{}, nil
}

func (s *RegistryServer) _routerConfigMock() *pb.RouterConfig {
	logrus.WithField("prefix", "server.mock").
		Tracef("_router_config_mock")
	return &pb.RouterConfig{}
}
