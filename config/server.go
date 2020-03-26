package config

import (
	"fmt"
	"time"
)

var Server ServerConfig

type ServerConfig struct {
	Port               int
	ReadTimeoutSecond  time.Duration
	WriteTimeoutSecond time.Duration
	IdleTimeoutSecond  time.Duration
	Address            string
}

func initServerConfig() {
	Server = ServerConfig{
		Port:               mustGetInt("SERVER_PORT"),
		ReadTimeoutSecond:  mustGetDurationS("SERVER_READ_TIMEOUT_SECOND"),
		WriteTimeoutSecond: mustGetDurationS("SERVER_WRITE_TIMEOUT_SECOND"),
		IdleTimeoutSecond:  mustGetDurationS("SERVER_IDLE_TIMEOUT_SECOND"),
	}
	Server.Address = fmt.Sprintf("0.0.0.0:%d", Server.Port)
}
