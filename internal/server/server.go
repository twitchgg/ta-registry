package server

import (
	"fmt"

	"ntsc.ac.cn/ta-registry/pkg/rpc"
)

type RegistryServer struct {
	conf      *RegistryServerConfig
	rpcServer *rpc.Server
}

// NewRegistryServer create registry server
func NewRegistryServer(conf *RegistryServerConfig) (*RegistryServer, error) {
	if conf == nil {
		return nil, fmt.Errorf("server config not define")
	}
	server := RegistryServer{}
	return &server, nil
}

func (s *RegistryServer) Start() chan error {
	return nil
}
