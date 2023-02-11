package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugaredLogger *zap.SugaredLogger

func Init() {
	encoder := getEncoder()
	writerSyncer := getWriterSyncer()
	core := zapcore.NewCore(encoder, writerSyncer, zap.DebugLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugaredLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSyncer() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}

func Info(message string) {
	sugaredLogger.Info(message)
}
func Debug(message string) {
	sugaredLogger.Debug(message)
}
func Error(message string) {
	sugaredLogger.Errorf(message)
}
func Panic(message string) {
	sugaredLogger.Panic(message)
}
func Sync() {
	sugaredLogger.Sync()
}
