package config

import (
	"fmt"
	"time"
)

var Database DatabaseConfig

type DatabaseConfig struct {
	Name               string
	Host               string
	User               string
	Password           string
	Port               int
	MaxPoolSize        int
	ConnectionURL      string
	ReadTimeoutSecond  time.Duration
	WriteTimeoutSecond time.Duration
}

func initDatabaseConfig() {
	Database = DatabaseConfig{
		Name:               mustGetString("DB_NAME"),
		Host:               mustGetString("DB_HOST"),
		User:               mustGetString("DB_USER"),
		Password:           mustGetString("DB_PASSWORD"),
		Port:               mustGetInt("DB_PORT"),
		MaxPoolSize:        mustGetInt("DB_MAX_POOL_SIZE"),
		ReadTimeoutSecond:  mustGetDurationS("DB_READ_TIMEOUT_SECOND"),
		WriteTimeoutSecond: mustGetDurationS("DB_WRITE_TIMEOUT_SECOND"),
	}
	Database.ConnectionURL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		Database.User,
		Database.Password,
		Database.Host,
		Database.Port,
		Database.Name)
}
