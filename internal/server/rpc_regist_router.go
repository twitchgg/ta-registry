package server

import (
	"context"

	"github.com/sirupsen/logrus"
	"ntsc.ac.cn/tas/tas-commons/pkg/pb"
	"ntsc.ac.cn/tas/tas-commons/pkg/rpc"
)

// RegistRouter regist wireguard router
func (s *RegistryServer) RegistRouter(ctx context.Context,
	req *pb.RegistRouterRequest) (*pb.RouterConfig, error) {
	if err := rpc.CheckMachineID(ctx, req.MachineID); err != nil {
		return nil, err
	}
	if logrus.IsLevelEnabled(logrus.TraceLevel) {
		return s._routerConfigMock(req.MachineID), nil
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
				// Addresses: []string{
				// 	"172.28.113.172/20",
				// 	"172.28.118.1/24",
				// },
				// Gateway: "172.28.112.1",
				DhcpClient: "dhclient",
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
							PubKey:   "h4r0T8MLH/gUxYzmQ5vc1R4cNGXqX6yAJBTsyaX1kiw=",
							PsKey:    "nTDIlBmzlY2k2JkPRMFp+4+lYgX7vYtv4wu6hiCIbjY=",
							PeerAddr: "10.200.201.1/32",
							AllowIPs: []string{
								"10.200.201.0/24",
							},
							Keepalive: 0,
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
