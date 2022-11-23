package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"dbName"`
	SSLMode  bool   `yaml:"sslMode"`
	Password string `yaml:"password"`
}

const (
	maxConn           = 50
	healthCheckPeriod = 1 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConn           = 10
	lazyConnect       = false
)

// NewPgxConn pool
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

	poolCfg.MaxConns = maxConn
	poolCfg.HealthCheckPeriod = healthCheckPeriod
	poolCfg.MaxConnIdleTime = maxConnIdleTime
	poolCfg.MaxConnLifetime = maxConnLifetime
	poolCfg.MinConns = minConn
	poolCfg.LazyConnect = lazyConnect

	connPool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("pgx.ConnectConfig", err)
	}

	return connPool, nil
}
