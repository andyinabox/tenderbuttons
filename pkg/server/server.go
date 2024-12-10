package server

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/acme/autocert"
)

type RunMode string

const (
	RunModeHTTP             RunMode = "http"
	RunModeHTTPSSelfSigned  RunMode = "https-ss"
	RunModeHTTPSLetsEncrypt RunMode = "https-le"
)

const (
	SelfSignedCertFilePath  = ".cert/localhost.crt"
	SelfSignedKeyFilePath   = ".cert/localhost.key"
	LetsEncryptCertCacheDir = "certs"
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
	Port         int
	Handler      http.Handler
	RunMode      RunMode
	AllowedHosts []string // only used for lets encrypt
}

type Server struct {
	srv *http.Server
	cm  *autocert.Manager
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

	s := &Server{
		srv: srv,
		cnf: cnf,
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
		s.cm = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(cnf.AllowedHosts...), //Your domain here
			Cache:      autocert.DirCache(LetsEncryptCertCacheDir),  //Folder for storing certificates
		}

		srv.TLSConfig = &tls.Config{
			GetCertificate: s.cm.GetCertificate,
			MinVersion:     tls.VersionTLS12, // improves cert reputation score at https://www.ssllabs.com/ssltest/
		}
	}

	return s
}

func (s *Server) Start() error {

	log.Info("starting server", "port", s.cnf.Port, "mode", s.cnf.RunMode)

	if s.cnf.RunMode == RunModeHTTPSLetsEncrypt {
		go func() {
			log.Debug("starting http server for lets encrypt")
			log.Fatal(http.ListenAndServe(":http", s.cm.HTTPHandler(nil)))
		}()
	}

	if s.cnf.RunMode == RunModeHTTPSSelfSigned || s.cnf.RunMode == RunModeHTTPSLetsEncrypt {
		return s.srv.ListenAndServeTLS("", "")
	}

	return s.srv.ListenAndServe()
}

func (s *Server) Close() {
	s.srv.Close()
}
