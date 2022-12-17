package config

import (
	"flag"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/logger"
	"os"

	"github.com/pkg/errors"

	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/postgres"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/constants"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Writer microservice microservice config path")
}

type Config struct {
	ServiceName string           `mapstructure:"serviceName"`
	Logger      *logger.Config   `mapstructure:"logger"`
	KafkaTopics KafkaTopics      `mapstructure:"kafkaTopics"`
	GRPC        GRPC             `mapstructure:"grpc"`
	Postgresql  *postgres.Config `mapstructure:"postgres"`
	Kafka       *kafka.Config    `mapstructure:"kafka"`
}

type GRPC struct {
	ReaderServicePort string `mapstructure:"readerServicePort"`
}

type KafkaTopics struct {
	ProductCreate     kafka.TopicConfig `mapstructure:"productCreate"`
	ProductDelete     kafka.TopicConfig `mapstructure:"productDelete"`
	ProductDeactivate kafka.TopicConfig `mapstructure:"productDeactivate"`
	ProductUpdate     kafka.TopicConfig `mapstructure:"productUpdate"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.CONFIG_PATH)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, fmt.Errorf("os.Getwd", err)
			}
			configPath = fmt.Sprintf("%s/%s", getwd, constants.BASE_CONFIG_PATH)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.DEFAULT_CONFIG_TYPE)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrapf(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	postgresHost := os.Getenv(constants.POSTGRES_SQL_HOST)
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}
	postgresPort := os.Getenv(constants.POSTGRES_SQL_PORT)
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}

	kafkaBrokers := os.Getenv(constants.KAFKA_BROKERS)
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}

	return cfg, nil
}
