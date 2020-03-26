package app

import (
	"github.com/imjuanleonard/palu-covid/config"
	"github.com/imjuanleonard/palu-covid/pkg/db"
	"github.com/imjuanleonard/palu-covid/pkg/logger"
	"time"
)

func Init() error {
	logger.SetupLogger(config.Log)
	return initDB()
}

func initDB() error {
	if err := db.Init(&db.Config{
		Driver:          "postgres",
		URL:             config.Database.ConnectionURL,
		MaxIdleConns:    config.Database.MaxPoolSize,
		MaxOpenConns:    config.Database.MaxPoolSize,
		ConnMaxLifeTime: 30 * time.Minute,
	}); err != nil {
		return err
	}

	return nil
}

func Close() {
	if err := db.Close(); err != nil {
		logger.Error(err)
	}
}
