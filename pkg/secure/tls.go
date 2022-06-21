package secure

import (
	"crypto/tls"
	"crypto/x509"
)

// NewTLSCerts create TLS x509 certificates
func NewTLSCerts(certCA, cert, privKey []byte) (
	tls.Certificate, *x509.CertPool, error) {
	certificate, err := tls.X509KeyPair(cert, privKey)
	if err != nil {
		return tls.Certificate{}, nil, err
	}
	certPool, err := DecodeCertificateChainPool(certCA)
	if err != nil {
		return tls.Certificate{}, nil, err
	}
	return certificate, certPool, nil
}

// NewTLSConfig create tls config
func NewTLSConfig(trusted, cert, privKey []byte, rav bool) (*tls.Config, error) {
	certificate, certPool, err := NewTLSCerts(
		trusted, cert, privKey)
	if err != nil {
		return nil, err
	}
	tlsConf := &tls.Config{
		MinVersion:               tls.VersionTLS11,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	}
	if rav {
		tlsConf.ClientAuth = tls.RequireAndVerifyClientCert
	}
	return tlsConf, nil
}
