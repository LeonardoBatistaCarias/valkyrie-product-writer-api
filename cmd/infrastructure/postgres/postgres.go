package postgres

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/constants"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"dbName"`
	SSLMode  bool   `yaml:"sslMode"`
	Password string `yaml:"password"`
}

func NewPgxConn(cfg *Config) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.Password,
	)

	poolCfg, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = constants.MAX_CONN
	poolCfg.HealthCheckPeriod = constants.HEALTH_CHECK_PERIOD
	poolCfg.MaxConnIdleTime = constants.MAX_CONN_IDLE_TIME
	poolCfg.MaxConnLifetime = constants.MAX_CONN_LIFETIME
	poolCfg.MinConns = constants.MIN_CONN
	poolCfg.LazyConnect = constants.LAZY_CONNECT

	connPool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgx.ConnectConfig")
	}

	return connPool, nil
}
