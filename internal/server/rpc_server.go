package server

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
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

func (s *RegistryServer) certCheckFunc(ctx context.Context) (context.Context, error) {
	pr, _ := peer.FromContext(ctx)
	if !s.rpcServer.IsSSL() {
		return nil, fmt.Errorf("not tls service")
	}
	logrus.WithField("prefix", "server").
		Debugf("client address: %s", pr.Addr.String())
	cert, err := rpc.GetClientCertificate(pr)
	if err != nil {
		return nil, err
	}
	logrus.WithField("prefix", "server").
		Debugf("client common name [%s],issuer common name[%s]",
			cert.Subject.CommonName, cert.Issuer.CommonName)
	return ctx, nil
}

func (s *RegistryServer) getCertExtValue(ctx context.Context, key string) (string, error) {
	pr, _ := peer.FromContext(ctx)
	cert, err := rpc.GetClientCertificate(pr)
	if err != nil {
		return "", fmt.Errorf("get client certificate failed: %s", err.Error())
	}
	for _, v := range cert.Extensions {
		if v.Id.String() == key {
			return string(v.Value), nil
		}
	}
	return "", fmt.Errorf("not found extension [%s]", key)
}
