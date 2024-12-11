// package tlsserver provides a webserver that can be run in https mode.
// it is only really tested for local development with a self-signed cert.
package tlsserver

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

type RunMode string

const (
	RunModeHTTP  RunMode = "http"
	RunModeHTTPS RunMode = "https"
)

var ports map[RunMode]int

func init() {
	// set default ports
	ports = make(map[RunMode]int)
	ports[RunModeHTTP] = 8080
	ports[RunModeHTTPS] = 4443
}

type Config struct {
	Port         int
	Handler      http.Handler
	RunMode      RunMode
	CertFilePath string
	KeyFilePath  string
}

type Server struct {
	srv *http.Server
	cnf *Config
}

func New(cnf *Config) *Server {

	if cnf.RunMode != RunModeHTTP && cnf.RunMode != RunModeHTTPS {
		log.Fatalf("invalid server run mode: %s", cnf.RunMode)
	}

	if cnf.Port == 0 {
		cnf.Port = ports[cnf.RunMode]
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cnf.Port),
		Handler: cnf.Handler,
	}

	s := &Server{
		srv: srv,
		cnf: cnf,
	}

	// configure for self-signed TLS
	if cnf.RunMode == RunModeHTTPS {
		log.Debug("configure for self-signed tls")
		serverTLSCert, err := tls.LoadX509KeyPair(cnf.CertFilePath, cnf.KeyFilePath)
		if err != nil {
			log.Fatalf("Error loading certificate and key file: %v", err)
		}

		srv.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{serverTLSCert},
		}

	}

	return s
}

func (s *Server) Start() error {

	log.Info("starting server", "port", s.cnf.Port, "mode", s.cnf.RunMode)

	if s.cnf.RunMode == RunModeHTTPS {
		return s.srv.ListenAndServeTLS("", "")
	}

	return s.srv.ListenAndServe()
}

func (s *Server) Close() {
	s.srv.Close()
}
