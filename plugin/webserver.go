package main

import (
	"encoding/json"
	"net/http"
	//"net"
	"errors"
	//"io/ioutil"
//	"strings"
	//"os"
//	"crypto/tls"
	"../housekeeper"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
)

type s_configuration struct {
	Webserver struct {
		Protocol string `json:"protocol"`
		Listen string `json:"listen"`
		WebPath string `json:"web_path"`
		Certificate string `json:"certificate"`
		CertificateKey string `json:"certificate_key"`
//		LogRequests bool `json:"log_requests"`
	} `json:"webserver"`
}

type s_websocketResponse struct {
	Topic string `json:"topic"`
	Message string `json:"message"`
}

var configuration s_configuration

var upgrader = websocket.Upgrader{}

/*func PongServer(w http.ResponseWriter, req *http.Request) {
//	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	_, _, err := net.SplitHostPort(req.RemoteAddr)

//	housekeeper.SharedInformation.Logger.Error(req.URL.Path)
//	housekeeper.SharedInformation.Logger.Error(configuration.Webserver.WebPath)

	if err != nil {
		housekeeper.SharedInformation.Logger.Error(err)
	}

//	if configuration.Webserver.LogRequests {
//		housekeeper.SharedInformation.Logger.Infof("%s:%s made an ip request", ip, port)
//	}

//	file := configuration.Webserver.WebPath + "/" + req.URL.Path[strings.LastIndex(req.URL.Path, "/") + 1:]
	file := configuration.Webserver.WebPath + req.URL.Path

	housekeeper.SharedInformation.Logger.Infof("> %s", file)

//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	housekeeper.SharedInformation.Logger.Error(w.Header())

	_, err = os.Stat(file)

	if err != nil {
		housekeeper.SharedInformation.Logger.Warningf("%s does not exist!", file)
		body, _ := ioutil.ReadFile(configuration.Webserver.WebPath + "/" + configuration.Webserver.Index)
		w.Write([]byte(body))
	} else {
		body, _ := ioutil.ReadFile(file)
		w.Write([]byte(body))
	}
}*/

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

func tryRead(conn *websocket.Conn) {
	clientResponse := s_websocketResponse{}

	for {
		_, p, err := conn.ReadMessage()

		if err != nil {
			housekeeper.SharedInformation.Logger.Error(err)
			break
		}

		json.Unmarshal(p, &clientResponse)

		housekeeper.SharedInformation.Logger.Info(clientResponse)

//		housekeeper.SharedInformation.Logger.Info(p)

		housekeeper.SharedInformation.MQTTClient.Publish(clientResponse.Topic, 0, true, clientResponse.Message)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		housekeeper.SharedInformation.Logger.Error(err)
		return
	}

	housekeeper.SharedInformation.Logger.Critical("blip")

	go tryRead(c)

	housekeeper.SharedInformation.MQTTClient.Subscribe("#", 0, func(client MQTT.Client, msg MQTT.Message) {
//		housekeeper.SharedInformation.Logger.Criticalf("webbie says * [%s] %s\n", msg.Topic(), string(msg.Payload()))

//			serverResponse.Topic = string(msg.Topic())
//		serverResponse.Message = string(msg.Payload())
		serverResponse := &s_websocketResponse{ Topic: string(msg.Topic()), Message: string(msg.Payload())}

		packedResponse, err := json.Marshal(serverResponse)

		if err != nil {
			housekeeper.SharedInformation.Logger.Error(err)
		}

		err = c.WriteMessage(websocket.TextMessage, packedResponse)

		if err != nil {
			housekeeper.SharedInformation.Logger.Critical("write:", err)
		}
	})
/*	if err != nil {
		housekeeper.SharedInformation.Logger.Critical("upgrade:", err)
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()

		if err != nil {
			housekeeper.SharedInformation.Logger.Critical("read:", err)
			break
		}

		housekeeper.SharedInformation.Logger.Criticalf("recv: %s", message)
		err = c.WriteMessage(mt, message)

		if err != nil {
			housekeeper.SharedInformation.Logger.Critical("write:", err)
			break
		}
	}*/
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
	fs := http.FileServer(http.Dir(configuration.Webserver.WebPath))

	mux.Handle("/", fs)
	mux.HandleFunc("/echo", echo)

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
