package server

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"ntsc.ac.cn/ta-registry/pkg/pb"
	"ntsc.ac.cn/ta-registry/pkg/rpc"
)

type RegistryServer struct {
	conf      *RegistryServerConfig
	rpcServer *rpc.Server
	discovery *rpc.EtcdDiscovery
}

// NewRegistryServer create registry server
func NewRegistryServer(conf *RegistryServerConfig) (*RegistryServer, error) {
	if conf == nil {
		return nil, fmt.Errorf("server config not define")
	}
	if err := conf.Check(); err != nil {
		return nil, fmt.Errorf("check server config failed: %s", err.Error())
	}
	rpcConf, err := conf.RPCConfig()
	if err != nil {
		return nil, fmt.Errorf("generate rpc config failed: %s", err.Error())
	}
	server := RegistryServer{
		conf: conf,
	}
	rpcServ, err := rpc.NewServer(rpcConf, []grpc.ServerOption{
		grpc.StreamInterceptor(
			rpc.StreamServerInterceptor(server.certCheckFunc)),
		grpc.UnaryInterceptor(
			rpc.UnaryServerInterceptor(server.certCheckFunc)),
	}, func(g *grpc.Server) {
		pb.RegisterRegistryServiceServer(g, &server)
	})
	if err != nil {
		return nil, fmt.Errorf("create grpc server failed: %s", err.Error())
	}
	server.rpcServer = rpcServ
	discovery, err := rpc.NewEtcdDiscovery(conf.GetEtcdEndpoints()...)
	if err != nil {
		return nil, fmt.Errorf("create etcd discovert failed: %s", err.Error())
	}
	server.discovery = discovery
	logrus.WithField("prefix", "server").
		Debugf("create etcd discovery with [%s] success", conf.EtcdEndpoints)
	return &server, nil
}

// Start start registry server
func (s *RegistryServer) Start() chan error {
	errChan := make(chan error, 1)
	go func() {
		err := <-s.rpcServer.Start()
		errChan <- err
	}()
	return errChan
}
