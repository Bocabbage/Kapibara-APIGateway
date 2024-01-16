package main

import (
	"kapibara-apigateway/internal/config" // config init
	_ "kapibara-apigateway/internal/logger"
	. "kapibara-apigateway/internal/router"
)

func main() {
	serverConfig := config.GlobalConfig.ServerConf
	serverPoint := serverConfig.ServerAddr + ":" + serverConfig.ServerPort
	if serverConfig.IsTLS {
		ServerEngine.RunTLS(serverPoint, serverConfig.CertFile, serverConfig.KeyFile)
	} else {
		ServerEngine.Run(serverPoint)
	}
}
