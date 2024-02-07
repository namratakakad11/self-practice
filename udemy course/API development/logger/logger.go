package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	// zap.NewDevelopment()

	config := zap.NewProductionConfig()

	encoderconfig := zap.NewProductionEncoderConfig()
	encoderconfig.TimeKey = "timestamp"
	encoderconfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderconfig.StacktraceKey = ""
	config.EncoderConfig = encoderconfig

	log, err = config.Build(zap.AddCallerSkip(1))

	// log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
func Debug(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
func Error(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
