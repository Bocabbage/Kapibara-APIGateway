// global log apis
package logger

// global logger
var logger *standardLogger

func init() {
	logger = newLogger()
}
