package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"

	kafkaClient "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/postgres"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/constants"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Writer microservice microservice config path")
}

type Config struct {
	ServiceName string              `mapstructure:"serviceName"`
	KafkaTopics KafkaTopics         `mapstructure:"kafkaTopics"`
	GRPC        GRPC                `mapstructure:"grpc"`
	Postgresql  *postgres.Config    `mapstructure:"postgres"`
	Kafka       *kafkaClient.Config `mapstructure:"kafka"`
}

type GRPC struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

type KafkaTopics struct {
	ProductCreate  kafkaClient.TopicConfig `mapstructure:"productCreate"`
	ProductCreated kafkaClient.TopicConfig `mapstructure:"productCreated"`
	ProductUpdate  kafkaClient.TopicConfig `mapstructure:"productUpdate"`
	ProductUpdated kafkaClient.TopicConfig `mapstructure:"productUpdated"`
	ProductDelete  kafkaClient.TopicConfig `mapstructure:"productDelete"`
	ProductDeleted kafkaClient.TopicConfig `mapstructure:"productDeleted"`
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

	grpcPort := os.Getenv(constants.GRPC_PORT)
	if grpcPort != "" {
		cfg.GRPC.Port = grpcPort
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
