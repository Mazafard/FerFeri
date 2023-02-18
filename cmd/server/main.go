package main

import (
	"github.com/goraz/onion"
	"github.com/mazafard/ferferi/pkg/logger"
	"github.com/mazafard/ferferi/pkg/signaler"
	"github.com/mazafard/ferferi/pkg/turn"
	"github.com/mazafard/ferferi/pkg/websocket"
	"os"
)

func main() {

	l1, err := onion.NewFileLayer("/tmp/shared.json", nil)
	if err != nil {
		logger.Errorf("Fail to read file: %v", err)
		os.Exit(1)
	}
	l2 := onion.NewEnvLayerPrefix("_", "APP")

	o := onion.New(l1, l2)

	publicIP := o.GetStringDefault("turn.public_ip", "empty")
	stunPort := o.GetIntDefault("turn.port", 0)
	if err != nil {
		stunPort = 3478
	}
	realm := o.GetStringDefault("turn.realm", "empty")

	turnConfig := turn.DefaultConfig()
	turnConfig.PublicIP = publicIP
	turnConfig.Port = stunPort
	turnConfig.Realm = realm
	turnServer := turn.NewTurnServer(turnConfig)

	newSignaler := signaler.NewSignaler(turnServer)
	wsServer := websocket.NewWebSocketServer(newSignaler.HandleNewWebSocket, newSignaler.HandleTurnServerCredentials)

	sslKey := o.GetStringDefault("general.key", "empty")
	sslCert := o.GetStringDefault("general.cert", "empty")

	bindAddress := o.GetStringDefault("general.bind", "empty")

	port := o.GetIntDefault("general.port", 0)

	htmlRoot := o.GetStringDefault("general.html_root", "empty")

	config := websocket.DefaultConfig()
	config.Host = bindAddress
	config.Port = port
	config.CertFile = sslCert
	config.KeyFile = sslKey
	config.HTMLRoot = htmlRoot

	wsServer.Bind(config)
}
