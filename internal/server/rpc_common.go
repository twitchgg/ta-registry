package server

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"ntsc.ac.cn/ta-registry/pkg/rpc"
	"ntsc.ac.cn/ta-registry/pkg/secure"
)

const (
	CERT_EXT_KEY_MACHINE_ID = "1.1.1.1.1.1"
)

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

func (s *RegistryServer) checkMachineID(ctx context.Context, reqMachineID string) error {
	pr, _ := peer.FromContext(ctx)
	cert, err := rpc.GetClientCertificate(pr)
	if err != nil {
		return fmt.Errorf("get client certificate failed: %s", err.Error())
	}
	machineID, err := secure.GetCertExtValue(cert, CERT_EXT_KEY_MACHINE_ID)
	if err != nil {
		return rpc.GenerateError(codes.InvalidArgument,
			fmt.Errorf("not found machine id from certificate"))
	}
	if machineID != reqMachineID {
		return rpc.GenerateError(codes.InvalidArgument,
			fmt.Errorf("machine id does not match"))
	}
	return nil
}
