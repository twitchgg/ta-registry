package server

import (
	"fmt"
	"net"
	"net/url"
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
