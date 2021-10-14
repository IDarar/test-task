package app

import (
	"fmt"

	"github.com/IDarar/test-task/internal/cache"
	"github.com/IDarar/test-task/internal/config"
	"github.com/IDarar/test-task/internal/mq"
	"github.com/IDarar/test-task/internal/repository"
	"github.com/IDarar/test-task/internal/service"
	"github.com/IDarar/test-task/internal/transport"
	"github.com/IDarar/test-task/pkg/kafka"
	"github.com/IDarar/test-task/pkg/postgresdb"
	"github.com/IDarar/test-task/pkg/redisdb"
)

func Run() error {
	cfg, err := config.Init("config.json")
	if err != nil {
		return fmt.Errorf("failed init config: %w", err)
	}

	pgDb, err := postgresdb.NewPostgresDB(cfg)
	if err != nil {
		return fmt.Errorf("failed init postgres: %w", err)
	}

	defer pgDb.Close()

	repo := repository.NewUsersRepo(pgDb)

	kafkaConf, err := kafka.GetKafkaConfig(cfg.Kafka)
	if err != nil {
		return fmt.Errorf("failed get kafka config: %w", err)
	}

	kafkaProducer, err := kafka.NewKafkaProducer(cfg.Kafka, kafkaConf)
	if err != nil {
		return fmt.Errorf("failed get kafka producer: %w", err)
	}

	kafkaUsers := mq.NewUsersKafka(kafkaProducer, cfg.Kafka.UsersEventTopic)

	rdb, err := redisdb.NewRedisDB(cfg)
	if err != nil {
		return fmt.Errorf("failed get redis client: %w", err)
	}

	cache := cache.NewUsersCache(cfg.Redis.UsersListTTL, rdb)

	services := service.NewUsersService(repo, cache, kafkaUsers)

	err = transport.RunUsersServer(cfg.GRPC, services)
	if err != nil {
		return fmt.Errorf("error during running server: %w", err)
	}

	return nil
}
