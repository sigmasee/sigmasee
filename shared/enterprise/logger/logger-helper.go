package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateProductionLogger() (*zap.Logger, *zap.SugaredLogger) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	return logger, logger.Sugar()
}

func CreateDevelopmentLogger() (*zap.Logger, *zap.SugaredLogger) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	return logger, logger.Sugar()
}
