package housekeeper

import (
	"encoding/json"
	"net/http"
	//"net"
	"errors"
	//"io/ioutil"
//	"strings"
	//"os"
//	"crypto/tls"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
)

type s_websocketResponse struct {
	Topic string `json:"topic"`
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{}

/*func PongServer(w http.ResponseWriter, req *http.Request) {
//	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	_, _, err := net.SplitHostPort(req.RemoteAddr)

//	SharedInformation.Logger.Error(req.URL.Path)
//	SharedInformation.Logger.Error(configuration.Webserver.WebPath)

	if err != nil {
		SharedInformation.Logger.Error(err)
	}

//	if configuration.Webserver.LogRequests {
//		SharedInformation.Logger.Infof("%s:%s made an ip request", ip, port)
//	}

//	file := configuration.Webserver.WebPath + "/" + req.URL.Path[strings.LastIndex(req.URL.Path, "/") + 1:]
	file := configuration.Webserver.WebPath + req.URL.Path

	SharedInformation.Logger.Infof("> %s", file)

//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	SharedInformation.Logger.Error(w.Header())

	_, err = os.Stat(file)

	if err != nil {
		SharedInformation.Logger.Warningf("%s does not exist!", file)
		body, _ := ioutil.ReadFile(configuration.Webserver.WebPath + "/" + configuration.Webserver.Index)
		w.Write([]byte(body))
	} else {
		body, _ := ioutil.ReadFile(file)
		w.Write([]byte(body))
	}
}*/

func validateConfiguration() error {
	if SharedInformation.Configuration.Webserver.Protocol == "" {
		return errors.New("Webserver.Protocol missing from configuration")
	}

	if SharedInformation.Configuration.Webserver.Listen == "" {
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
			SharedInformation.Logger.Error(err)
			break
		}

		json.Unmarshal(p, &clientResponse)

		SharedInformation.Logger.Info(clientResponse)

//		SharedInformation.Logger.Info(p)

		SharedInformation.MQTTClient.Publish(clientResponse.Topic, 0, true, clientResponse.Message)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		SharedInformation.Logger.Error(err)
		return
	}

	SharedInformation.Logger.Critical("blip")

	go tryRead(c)

	SharedInformation.MQTTClient.Subscribe("#", 0, func(client MQTT.Client, msg MQTT.Message) {
//		SharedInformation.Logger.Criticalf("webbie says * [%s] %s\n", msg.Topic(), string(msg.Payload()))

//			serverResponse.Topic = string(msg.Topic())
//		serverResponse.Message = string(msg.Payload())
		serverResponse := &s_websocketResponse{ Topic: string(msg.Topic()), Message: string(msg.Payload())}

		packedResponse, err := json.Marshal(serverResponse)

		if err != nil {
			SharedInformation.Logger.Error(err)
		}

		err = c.WriteMessage(websocket.TextMessage, packedResponse)

		if err != nil {
			SharedInformation.Logger.Critical("write:", err)
		}
	})
/*	if err != nil {
		SharedInformation.Logger.Critical("upgrade:", err)
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()

		if err != nil {
			SharedInformation.Logger.Critical("read:", err)
			break
		}

		SharedInformation.Logger.Criticalf("recv: %s", message)
		err = c.WriteMessage(mt, message)

		if err != nil {
			SharedInformation.Logger.Critical("write:", err)
			break
		}
	}*/
}

func StartWebserver() error {
	err := validateConfiguration()

	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(SharedInformation.Configuration.Webserver.WebPath))

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
		Addr: SharedInformation.Configuration.Webserver.Listen,
		Handler: mux,
/*		TLSConfig: cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),*/
	}
	
	if err != nil {
		return err
	}

	SharedInformation.Logger.Info("Serving webserver!")
	
	err = srv.ListenAndServe()
//	err = srv.ListenAndServeTLS(housekeeper.Configuration.Webserver.Certificate, configuration.Webserver.CertificateKey)

	return err
}