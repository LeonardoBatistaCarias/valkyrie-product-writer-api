package server

import (
	"context"
	"log"

	"net"
	"strconv"

	kafkaClient "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka"
	"github.com/pkg/errors"

	"github.com/segmentio/kafka-go"
)

const (
	stackSize = 1 << 10 // 1 KB
)

func (s *server) connectKafkaBrokers(ctx context.Context) error {
	kafkaConn, err := kafkaClient.NewKafkaConn(ctx, s.cfg.Kafka)
	if err != nil {
		return errors.Wrap(err, "kafka.NewKafkaCon")
	}

	s.kafkaConn = kafkaConn

	brokers, err := s.kafkaConn.Brokers()
	if err != nil {
		return errors.Wrapf(err, "kafkaConn.Brokers")
	}

	log.Printf("kafka connected to brokers: %+v", brokers)

	return nil
}

func (s *server) initKafkaTopics(ctx context.Context) {
	controller, err := s.kafkaConn.Controller()
	if err != nil {
		errors.Wrapf(err, "kafkaConn.Controller")
		return
	}

	controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	log.Printf("kafka controller uri: %s", controllerURI)

	conn, err := kafka.DialContext(ctx, "tcp", controllerURI)
	if err != nil {
		errors.Wrapf(err, "initKafkaTopicDialContext")
		return
	}
	defer conn.Close() // nolint: errcheck

	log.Printf("established new kafka controller connection: %s", controllerURI)

	productCreateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.ProductCreate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.ProductCreate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.ProductCreate.ReplicationFactor,
	}

	productDeleteTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.ProductDelete.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.ProductDelete.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.ProductDelete.ReplicationFactor,
	}

	productUpdateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.ProductUpdate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.ProductUpdate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.ProductUpdate.ReplicationFactor,
	}

	if err := conn.CreateTopics(
		productCreateTopic,
		productDeleteTopic,
		productUpdateTopic,
	); err != nil {
		log.Printf("kafkaConn.CreateTopics", err)
		return
	}

	log.Printf("kafka topics created or already exists: %+v", []kafka.TopicConfig{productCreateTopic})
}

func (s *server) getConsumerGroupTopics() []string {
	return []string{
		s.cfg.KafkaTopics.ProductCreate.TopicName,
		s.cfg.KafkaTopics.ProductDelete.TopicName,
		s.cfg.KafkaTopics.ProductUpdate.TopicName,
	}
}
