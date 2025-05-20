package initiator

import (
	"context"
	"fmt"
	"strings"
	"url_shortener/platform/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

func InitMigration(path string, url string, logger logger.Logger) *migrate.Migrate {
	m, err := migrate.New(
		fmt.Sprintf("file://%s", path),
		fmt.Sprintf("cockroachdb://%s", strings.Split(url, "://")[1]))
	if err != nil {
		logger.Fatal(context.Background(), "failed to migrate schema", zap.Error(err))
	}

	return m

}

func UpMigration(m *migrate.Migrate, logger logger.Logger) {
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Fatal(context.Background(), "failed to up migration", zap.Error(err))
	}
}
