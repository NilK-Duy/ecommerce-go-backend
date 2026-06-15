package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age: %d", "John", 30) // Like fmt.Printf(format, a)

	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "John"), zap.Int("age", 30))

	// 2.
	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")

	// // Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// // Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// 3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("info log", zap.Int("line", 1))
	logger.Error("error log", zap.Int("line", 2))
}

// format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1781511085.2000034 -> 2026-06-15T15:11:25.199+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encodeConfig.TimeKey = "Time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.log.go:24"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stdout)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
