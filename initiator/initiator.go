package initiator

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	persistencedb "url_shortener/internal/constant/model/persistenceDB"
	"url_shortener/internal/handler/middleware"

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
	pool := InitDB(context.Background(), Options{
		Url:             viper.GetString("db.url"),
		MaxConnIdleTime: viper.GetDuration("db.max_idle_time"),
	}, log)
	log.Info(context.Background(), "database initialized")

	log.Info(context.Background(), "initializing migration")
	m := InitMigration(viper.GetString("migrations.file"), viper.GetString("db.url"), log)
	log.Info(context.Background(), "migration initialized")

	log.Info(context.Background(), "initializing up migration")
	UpMigration(m, log)
	log.Info(context.Background(), "Up migration initialized")

	log.Info(context.Background(), "intializing persistence layer")
	persistence := InitPersistence(persistencedb.New(pool, log), log)
	log.Info(context.Background(), "persistence layer initialized")

	log.Info(context.Background(), "initializing module layer")
	module := InitModule(persistence, log)
	log.Info(context.Background(), "module layer initialized")

	log.Info(context.Background(), "initializing handler layer")
	handler := InitHandler(module, log)
	log.Info(context.Background(), "handler layer initialized")

	log.Info(context.Background(), "initializing routes")
	gin_server := gin.Default()
	v1 := gin_server.Group("/v1")
	gin_server.Use(middleware.ErrorHandler())
	InitRoute(v1, handler)
	log.Info(context.Background(), "route initialized")

	log.Info(context.Background(), "initializing server")
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
