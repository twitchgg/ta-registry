package server

import (
	"context"

	"github.com/sirupsen/logrus"
	"ntsc.ac.cn/tas/tas-commons/pkg/pb"
	"ntsc.ac.cn/tas/tas-commons/pkg/rpc"
)

func (s *RegistryServer) RegistServer(ctx context.Context,
	req *pb.RegistServerRequest) (*pb.ServerConfig, error) {
	if err := rpc.CheckMachineID(ctx, req.MachineID); err != nil {
		return nil, err
	}
	if logrus.IsLevelEnabled(logrus.TraceLevel) {
		return s._serverConfigMock(req.MachineID), nil
	}
	return &pb.ServerConfig{}, nil
}

func (s *RegistryServer) _serverConfigMock(machineID string) *pb.ServerConfig {
	switch machineID {
	case "db152eb167d74e39acaf388fe4c7bf70":
		return &pb.ServerConfig{
			ServiceMode:    []*pb.ServiceMode{},
			MiddleWareMode: []*pb.MiddleWareMode{},
		}
	default:
		return &pb.ServerConfig{}
	}
}
