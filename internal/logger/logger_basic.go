// logger definition
package logger

import (
	"kapibara-apigateway/internal/config"
	"os"

	"github.com/sirupsen/logrus"
)

// global log fileds
var serviceName logrus.Fields

func newLogger() *standardLogger {
	logLevel := config.GlobalConfig.LogConf.Level
	serviceName = logrus.Fields{
		"service": config.GlobalConfig.LogConf.ServiceName,
	}

	var baseLogger = logrus.New()
	var standardLogger = &standardLogger{baseLogger}

	// Close call-train tracing
	standardLogger.SetReportCaller(false)

	// Set loglevel
	switch {
	case logLevel == "DEBUG":
		standardLogger.SetLevel(logrus.DebugLevel)
	case logLevel == "WARNING":
		standardLogger.SetLevel(logrus.WarnLevel)
	case logLevel == "ERROR":
		standardLogger.SetLevel(logrus.ErrorLevel)
	case logLevel == "FATAL":
		standardLogger.SetLevel(logrus.FatalLevel)
	case logLevel == "INFO":
	default:
		standardLogger.SetLevel(logrus.InfoLevel)
	}

	// Set logformat
	standardLogger.SetFormatter(&logrus.JSONFormatter{})

	// Set output
	standardLogger.SetOutput(os.Stdout)

	return standardLogger
}

// --------------------

// logger type definition
type standardLogger struct {
	*logrus.Logger
}

func (log *standardLogger) info(v ...interface{}) {
	log.WithFields(serviceName).Info(v...)
}

func (log *standardLogger) warn(v ...interface{}) {
	log.WithFields(serviceName).Warn(v...)
}

func (log *standardLogger) error(v ...interface{}) {
	log.WithFields(serviceName).Error(v...)
}

func (log *standardLogger) fatal(v ...interface{}) {
	log.WithFields(serviceName).Fatal(v...)
}

func (log *standardLogger) debug(v ...interface{}) {
	log.WithFields(serviceName).Debug(v...)
}
