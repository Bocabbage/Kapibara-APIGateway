// global log apis
package logger

// global logger
var logger *standardLogger

func init() {
	logger = newLogger()
}

func Error(v ...interface{}) {
	logger.error(v...)
}

func Warn(v ...interface{}) {
	logger.warn(v...)
}

func Debug(v ...interface{}) {
	logger.debug(v...)
}

func Info(v ...interface{}) {
	logger.info(v...)
}

func Fatal(v ...interface{}) {
	logger.fatal(v...)
}
