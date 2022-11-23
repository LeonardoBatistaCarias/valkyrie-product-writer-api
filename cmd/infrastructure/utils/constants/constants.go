package constants

import "time"

const (
	// Config Constants
	GRPC_PORT           = "GRPC_PORT"
	CONFIG_PATH         = "CONFIG_PATH"
	DEFAULT_CONFIG_TYPE = "yaml"
	HTTP_PORT           = "HTTP_PORT"
	POSTGRES_SQL_HOST   = "POSTGRES_HOST"
	POSTGRES_SQL_PORT   = "POSTGRES_PORT"

	// Kafka Config Constants
	KAFKA_BROKERS        = "KAFKA_BROKERS"
	WRITER_READ_TIMEOUT  = 10 * time.Second
	WRITER_WRITE_TIMEOUT = 10 * time.Second
	WRITER_REQUIRED_ACKS = -1
	WRITER_MAX_ATTEMPTS  = 3
)
