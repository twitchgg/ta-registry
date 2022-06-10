package server

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"ntsc.ac.cn/ta-registry/pkg/pb"
	"ntsc.ac.cn/ta-registry/pkg/rpc"
	"ntsc.ac.cn/ta-registry/pkg/secure"
)

const (
	CERT_EXT_KEY_MACHINE_ID = "1.1.1.1.1.1"
)

// RegistRouter regist wireguard router
func (s *RegistryServer) RegistRouter(ctx context.Context,
	req *pb.RegistRouterRequest) (*pb.RouterConfig, error) {
	pr, _ := peer.FromContext(ctx)
	cert, err := rpc.GetClientCertificate(pr)
	if err != nil {
		return nil, fmt.Errorf("get client certificate failed: %s", err.Error())
	}
	machineID, err := secure.GetCertExtValue(cert, CERT_EXT_KEY_MACHINE_ID)
	if err != nil {
		return nil, rpc.GenerateError(codes.InvalidArgument,
			fmt.Errorf("not found machine id from certificate"))
	}
	if machineID != req.MachineID {
		return nil, rpc.GenerateError(codes.InvalidArgument,
			fmt.Errorf("machine id does not match"))
	}
	if logrus.IsLevelEnabled(logrus.TraceLevel) {
		return s._routerConfigMock(machineID), nil
	}
	return &pb.RouterConfig{}, nil
}

func (s *RegistryServer) _routerConfigMock(machineID string) *pb.RouterConfig {
	logrus.WithField("prefix", "server.mock").
		Tracef("_router_config_mock")
	switch machineID {
	case "db152eb167d74e39acaf388fe4c7bf70":
		return &pb.RouterConfig{
			WanInfo: &pb.EthernetCard{
				Name: "eth0",
				Addresses: []string{
					"172.28.113.172/20",
					"172.28.118.1/24",
				},
				Gateway: "172.28.112.1",
			},
			WgConfig: []*pb.WireguardConfig{
				{
					Name: "wg0",
					InterfaceDef: &pb.WireguardInterface{
						Address: "172.16.0.1/32",
						Port:    51820,
						PrivKey: "GG9Qp95vtkkuWlfQwS1+4KxfcsnTqzt8A6/SJ2IrK3U=",
					},
					Peers: []*pb.WireguardPeer{
						{
							Name:     "dev peer",
							Desc:     "dev peer",
							PubKey:   "jPB9cY+cJvZnanF3O8OHO6adCBse0bUSJHWxUznkLXU=",
							PsKey:    "nTDIlBmzlY2k2JkPRMFp+4+lYgX7vYtv4wu6hiCIbjY=",
							PeerAddr: "10.200.200.4/32",
							AllowIPs: []string{
								"172.28.113.0/24",
							},
							Keepalive: 25,
						},
					},
				},
			},
			IptablesRules: &pb.IPTablesRules{
				WhiteList:            true,
				MasqueraedInterfaces: []string{"eth0"},
			},
			DnsServer: []string{"192.168.2.1", "114.114.114.114"},
		}
	default:
		return &pb.RouterConfig{}
	}
}
