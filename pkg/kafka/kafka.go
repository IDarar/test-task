package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"strings"

	"github.com/IDarar/test-task/pkg/scram"

	"github.com/IDarar/test-task/internal/config"
	"github.com/Shopify/sarama"
)

func GetKafkaConfig(cfg config.KafkaConfig) (*sarama.Config, error) {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Version = sarama.V2_0_0_0
	conf.ClientID = "sasl_scram_client"
	conf.Net.SASL.Enable = true
	conf.Net.SASL.Handshake = true
	conf.Net.SASL.User = cfg.UsernameSASL
	conf.Net.SASL.Password = cfg.PaswordSASL
	conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &scram.XDGSCRAMClient{HashGeneratorFcn: scram.SHA512} }
	conf.Net.SASL.Mechanism = sarama.SASLMechanism(sarama.SASLTypeSCRAMSHA512)

	certs := x509.NewCertPool()

	certs.AppendCertsFromPEM([]byte(cfg.Ca))

	conf.Net.TLS.Enable = true
	conf.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            certs,
	}

	return conf, nil
}

func InitTopics(cfg config.KafkaConfig, kafkaConf *sarama.Config) error {
	brokers := cfg.Brokers
	splitBrokers := strings.Split(brokers, ",")

	ad, err := sarama.NewClusterAdmin(splitBrokers, kafkaConf)
	if err != nil {
		return fmt.Errorf("failed get admin: %w", err)
	}

	err = ad.CreateTopic(cfg.UsersEventTopic, &sarama.TopicDetail{NumPartitions: int32(cfg.UsersEventTopicNumPartition), ReplicationFactor: int16(cfg.UsersEventTopicReplicationFator)}, false)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}

	return nil
}

func NewKafkaProducer(cfg config.KafkaConfig, kafkaConf *sarama.Config) (sarama.SyncProducer, error) {
	brokers := cfg.Brokers
	splitBrokers := strings.Split(brokers, ",")

	syncProducer, err := sarama.NewSyncProducer(splitBrokers, kafkaConf)
	if err != nil {
		return nil, fmt.Errorf("failed create producer: %w", err)
	}

	return syncProducer, nil
}

func GetKafkaAdmin(cfg config.KafkaConfig, kafkaConf *sarama.Config) (sarama.ClusterAdmin, error) {
	brokers := cfg.Brokers
	splitBrokers := strings.Split(brokers, ",")

	ad, err := sarama.NewClusterAdmin(splitBrokers, kafkaConf)
	if err != nil {
		return nil, fmt.Errorf("failed get admin: %w", err)
	}

	return ad, nil
}
