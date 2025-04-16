package initiator

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Initiator() {
	log := InitLogger()
	log.Info(context.Background(), "logger initialized")

	log.Info(context.Background(), "initalizing config")
	InitConfig(context.Background(), "config", "config", "yaml", log)
	log.Info(context.Background(), "config initialized")

	log.Info(context.Background(), "initializing database")
	InitDB(context.Background(), Options{
		Url:             viper.GetString("db.url"),
		MaxConnIdleTime: viper.GetDuration("db.max_idle_time"),
	}, log)
	log.Info(context.Background(), "database initialized")

	log.Info(context.Background(), "intializing persistence layer")
	InitPersistence()
	log.Info(context.Background(), "persistence layer initialized")

	log.Info(context.Background(), "initializing module layer")
	InitModule()
	log.Info(context.Background(), "module layer initialized")

	log.Info(context.Background(), "initializing handler layer")
	InitHandler()
	log.Info(context.Background(), "handler layer initialized")

	log.Info(context.Background(), "initializing routes")
	InitRoute()
	log.Info(context.Background(), "route initialized")

	log.Info(context.Background(), "initializing server")
	gin_server := gin.Default()

	server := &http.Server{
		Addr:              viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		ReadTimeout:       viper.GetDuration("server.read_timeout"),
		ReadHeaderTimeout: viper.GetDuration("server.read_header_timeout"),
		WriteTimeout:      viper.GetDuration("server.write_timeout"),
		IdleTimeout:       viper.GetDuration("server.idle_timeout"),
		MaxHeaderBytes:    int(viper.GetSizeInBytes("server.max_header_bytes")),
		Handler:           gin_server,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	go func() {
		log.Info(context.Background(), "server listening",
			zap.String("host", viper.GetString("server.host")),
			zap.Int("port", viper.GetInt("server.port")),
		)
		log.Info(context.Background(), fmt.Sprintf("server stopped %v", server.ListenAndServe()))
	}()

	sig := <-quit
	log.Info(context.Background(), fmt.Sprintf("shutting down server %v", sig))
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("server.timeout"))
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(context.Background(), "error while shutting down server", zap.Error(err))
	}

	log.Info(context.Background(), "server shutdown successfully")
}
