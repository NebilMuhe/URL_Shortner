package initiator

import (
	"log"
	"url_shortener/platform/logger"

	"go.uber.org/zap"
)



func InitLogger() logger.Logger{
	zapLogger,err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initiate zap logger %v",err)
	}

	return logger.InitLogger(zapLogger)
}