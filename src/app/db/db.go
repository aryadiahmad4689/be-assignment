package db

import (
	"be-assesment/src/app/db/utils"
	"context"
	"errors"
)

type Config struct {
	Driver   string
	Source   string
	Host     string
	Password string
	Db       string
}

func Init(ctx context.Context, config Config) (interface{}, error) {
	switch config.Driver {
	case "postgres":
		return utils.InitPostgres(config.Driver, config.Source)
	default:
		return nil, errors.New("database driver not found")
	}
}
