package main

import (
	"flag"
	"fmt"
	"kapibara-apigateway/internal/config" // config init
	"kapibara-apigateway/internal/logger"
	"kapibara-apigateway/internal/router"
)

// @BasePath /

func main() {
	serverConfig := config.GlobalConfig.ServerConf

	var useTLS bool
	var devMode bool
	flag.BoolVar(&useTLS, "tls", false, "enable TLS.")
	flag.BoolVar(&devMode, "dev", false, "enable dev-mode.")
	flag.Parse()

	// configuration load test
	msg := config.GlobalConfig.HealthConf.TestConfig
	logger.Info(fmt.Sprintf("test config: %s\n", msg))

	// if devMode {

	// }

	serverPoint := serverConfig.ServerAddr + ":" + serverConfig.ServerPort
	logger.Debug(fmt.Sprintf("serverPoint: %s\n", serverPoint))
	if useTLS {
		logger.Info("Start to serve in HTTPS-mode.")
		router.ServerEngine.RunTLS(serverPoint, serverConfig.CertFile, serverConfig.KeyFile)
	} else {
		logger.Info("Start to serve in HTTP-mode.")
		router.ServerEngine.Run(serverPoint)
	}
}
