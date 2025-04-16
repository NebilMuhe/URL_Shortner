package initiator

import (
	"context"
	"url_shortener/platform/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig(ctx context.Context,path, name, file_type string, log logger.Logger) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(file_type)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(ctx, "failed to read config", zap.Error(err))
	}

	viper.WatchConfig()
}
