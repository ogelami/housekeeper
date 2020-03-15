package main

import (
//	"encoding/json"
	"net/http"
	"net"
	"errors"
	"io/ioutil"
	"strings"
	"os"
//	"crypto/tls"
	"../housekeeper"
)

type s_configuration struct {
	Webserver struct {
		Protocol string `json:"protocol"`
		Listen string `json:"listen"`
		Index string `json:"index"`
		WebPath string `json:"web_path"`
		Certificate string `json:"certificate"`
		CertificateKey string `json:"certificate_key"`
		LogRequests bool `json:"log_requests"`
	} `json:"webserver"`
}

var configuration s_configuration

func PongServer(w http.ResponseWriter, req *http.Request) {
	ip, port, err := net.SplitHostPort(req.RemoteAddr)

	if err != nil {
		housekeeper.SharedInformation.Logger.Error(err)
	}

	if configuration.Webserver.LogRequests {
		housekeeper.SharedInformation.Logger.Infof("%s:%s made an ip request", ip, port)
	}

	file := configuration.Webserver.WebPath + "/" + req.URL.Path[strings.LastIndex(req.URL.Path, "/") + 1:]

	housekeeper.SharedInformation.Logger.Infof("> %s", file)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	_, err = os.Stat(file)

	if err != nil {
		housekeeper.SharedInformation.Logger.Warningf("%s does not exist!", file)
		body, _ := ioutil.ReadFile(configuration.Webserver.WebPath + "/" + configuration.Webserver.Index)
		w.Write([]byte(body))
	} else {
		body, _ := ioutil.ReadFile(file)
		w.Write([]byte(body))
	}
}

func validateConfiguration() error {
	if configuration.Webserver.Protocol == "" {
		return errors.New("Webserver.Protocol missing from configuration")
	}

	if configuration.Webserver.Listen == "" {
		return errors.New("Webserver.Listen missing from configuration")
	}

/*	if configuration.Webserver.Certificate == "" {
		return errors.New("Webserver.Certificate missing from configuration")
	}

	if configuration.Webserver.CertificateKey == "" {
		return errors.New("Webserver.CertificateKey missing from configuration")
	}*/

	return nil
}

func Startup() error {
	err := housekeeper.ParseConfig(&configuration)

	if err != nil {
		return err
	}

	err = validateConfiguration()

	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", PongServer)

/*	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{
			tls.CurveP521,
			tls.CurveP384,
			tls.CurveP256,
		},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}*/

	srv := &http.Server{
		Addr: configuration.Webserver.Listen,
		Handler: mux,
/*		TLSConfig: cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),*/
	}
	
	if err == nil {
		housekeeper.SharedInformation.Logger.Info("Serving webserver!")
	}

	err = srv.ListenAndServe()
//	err = srv.ListenAndServeTLS(configuration.Webserver.Certificate, configuration.Webserver.CertificateKey)

	return err
}
