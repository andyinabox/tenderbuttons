package server

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

type RunMode string

const (
	RunModeHTTP             RunMode = "http"
	RunModeHTTPSSelfSigned  RunMode = "https-ss"
	RunModeHTTPSLetsEncrypt RunMode = "https-le"
)

const (
	SelfSignedCertFilePath = ".cert/localhost.crt"
	SelfSignedKeyFilePath  = ".cert/localhost.key"
)

var ports map[RunMode]int

func init() {
	// set default ports
	ports = make(map[RunMode]int)
	ports[RunModeHTTP] = 8080
	ports[RunModeHTTPSSelfSigned] = 4443
	ports[RunModeHTTPSLetsEncrypt] = 443
}

type Config struct {
	Port    int
	Handler http.Handler
	RunMode RunMode
}

type Server struct {
	srv *http.Server
	cnf *Config
}

func New(cnf *Config) *Server {

	if cnf.RunMode != RunModeHTTP && cnf.RunMode != RunModeHTTPSSelfSigned && cnf.RunMode != RunModeHTTPSLetsEncrypt {
		log.Fatalf("invalid server run mode: %s", cnf.RunMode)
	}

	if cnf.Port == 0 {
		cnf.Port = ports[cnf.RunMode]
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cnf.Port),
		Handler: cnf.Handler,
	}

	// configure for self-signed TLS
	if cnf.RunMode == RunModeHTTPSSelfSigned {
		log.Debug("configure for self-signed tls")
		serverTLSCert, err := tls.LoadX509KeyPair(SelfSignedCertFilePath, SelfSignedKeyFilePath)
		if err != nil {
			log.Fatalf("Error loading certificate and key file: %v", err)
		}

		srv.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{serverTLSCert},
		}

	}

	// configure for lets encrypt TLS
	if cnf.RunMode == RunModeHTTPSLetsEncrypt {
		log.Fatal("lets-encrypt run mode not implemented yet")
	}

	return &Server{
		srv: srv,
		cnf: cnf,
	}
}

func (s *Server) Start() error {

	log.Info("starting server", "port", s.cnf.Port, "mode", s.cnf.RunMode)

	if s.cnf.RunMode == RunModeHTTPSSelfSigned || s.cnf.RunMode == RunModeHTTPSLetsEncrypt {
		return s.srv.ListenAndServeTLS("", "")
	}

	return s.srv.ListenAndServe()
}

func (s *Server) Close() {
	s.srv.Close()
}
