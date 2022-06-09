package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/denisbrodbeck/machineid"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ntsc.ac.cn/ta-registry/pkg/pb"
	"ntsc.ac.cn/ta-registry/pkg/rpc"
)

func tc(t *testing.T) pb.RegistryServiceClient {
	certPath, err := filepath.Abs("../../etc/certs")
	if err != nil {
		t.Fatal(err)
	}
	caPath := certPath + "/trusted.crt"
	serverCertPath := certPath + "/client.crt"
	privKeyPath := certPath + "/client.key"
	trusted, err := ioutil.ReadFile(caPath)
	if err != nil {
		t.Fatal(err)
	}
	cert, err := ioutil.ReadFile(serverCertPath)
	if err != nil {
		t.Fatal(err)
	}
	privKey, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		t.Fatal(err)
	}
	tlsConf, err := rpc.NewClientTLSConfig(&rpc.ClientTLSConfig{
		CACert:     trusted,
		Cert:       cert,
		PrivKey:    privKey,
		ServerName: "s1.restry.ta.ntsc.ac.cn",
	})
	if err != nil {
		t.Fatal(err)
	}
	conn, err := rpc.DialRPCConn(&rpc.DialOptions{
		RemoteAddr: "tcp://127.0.0.1:1358",
		TLSConfig:  tlsConf,
	})
	if err != nil {
		t.Fatal(err)
	}
	return pb.NewRegistryServiceClient(conn)
}

func TestRegistRouter(t *testing.T) {
	c := tc(t)
	id, err := machineid.ID()
	if err != nil {
		t.Fatal(err)
	}
	conf, err := c.RegistRouter(context.Background(), &pb.RegistRouterRequest{
		SysTime:   timestamppb.Now(),
		MachineID: id,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(conf)
}

func TestETCD(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()
}
