package main

import (
	"fmt"
	"kapibara-apigateway/internal/config" // config init
	"kapibara-apigateway/internal/logger"
	_ "kapibara-apigateway/internal/logger"
	. "kapibara-apigateway/internal/router"
)

func main() {
	serverConfig := config.GlobalConfig.ServerConf

	// configuration load test
	msg := config.GlobalConfig.HealthConf.TestConfig
	logger.Info(fmt.Sprintf("test config: %s\n", msg))

	serverPoint := serverConfig.ServerAddr + ":" + serverConfig.ServerPort
	if serverConfig.IsTLS {
		logger.Info("Start to serve in HTTPS-mode.")
		ServerEngine.RunTLS(serverPoint, serverConfig.CertFile, serverConfig.KeyFile)
	} else {
		logger.Info("Start to serve in HTTP-mode.")
		ServerEngine.Run(serverPoint)
	}
}
