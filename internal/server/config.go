package server

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/url"

	"ntsc.ac.cn/ta-registry/pkg/rpc"
)

const (
	TRUSTED_CERT_CHAIN_NAME = "trusted.crt"
	SERVER_CERT_NAME        = "server.crt"
	SERVER_PRIVATE_KEY_NAME = "server.key"
)

// RegistryServerConfig registry server config
type RegistryServerConfig struct {
	Listener string
	CertPath string
}

// Check check registry server config
func (c *RegistryServerConfig) Check() error {
	if c.Listener == "" {
		return fmt.Errorf("listener url not define")
	}
	uri, err := url.Parse(c.Listener)
	if err != nil {
		return fmt.Errorf("parse listener url failed: %s", err.Error())
	}
	if uri.Scheme != "tcp" {
		return fmt.Errorf("unsupport scheme [%s]", uri.Scheme)
	}
	if uri.Port() == "" {
		return fmt.Errorf("listener port not define")
	}
	if _, err := net.ResolveIPAddr(uri.Scheme, uri.Host); err != nil {
		return fmt.Errorf("parse listener addr failed: %s", err.Error())
	}
	return nil
}

// RPCConfig get grpc config
func (c *RegistryServerConfig) RPCConfig() (*rpc.ServerConfig, error) {
	trustedPath := c.CertPath + "/" + TRUSTED_CERT_CHAIN_NAME
	certPath := c.CertPath + "/" + SERVER_CERT_NAME
	privKeyPath := c.CertPath + "/" + SERVER_PRIVATE_KEY_NAME
	var trusted, cert, privKey []byte
	var err error
	if trusted, err = ioutil.ReadFile(trustedPath); err != nil {
		return nil, fmt.Errorf("read trusted certificate chain failed: %s", err.Error())
	}
	if cert, err = ioutil.ReadFile(certPath); err != nil {
		return nil, fmt.Errorf("read server certificate failed: %s", err.Error())
	}
	if privKey, err = ioutil.ReadFile(privKeyPath); err != nil {
		return nil, fmt.Errorf("read server private key failed: %s", err.Error())
	}
	uri, _ := url.Parse(c.Listener)
	return &rpc.ServerConfig{
		TrustedCert:      trusted,
		ServerCert:       cert,
		ServerPrivKey:    privKey,
		RequireAndVerify: true,
		BindAddr:         uri.Host,
	}, nil
}
