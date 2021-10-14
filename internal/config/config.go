package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	GRPC       GRPCConfig       `json:"grpc,omitempty"`
	Postgres   PostgresConfig   `json:"postgres,omitempty"`
	Redis      RedisConfig      `json:"redis,omitempty"`
	Kafka      KafkaConfig      `json:"kafka,omitempty"`
	ClickHouse ClickHouseConfig `json:"click_house,omitempty"`
}

type GRPCConfig struct {
	Port string `json:"port,omitempty"`
}

type PostgresConfig struct {
	URL string `json:"url,omitempty"`
}

type RedisConfig struct {
	Addr         string `json:"addr,omitempty"`
	Password     string `json:"password,omitempty"`
	DB           int    `json:"db,omitempty"`
	UsersListTTL int    `json:"users_list_ttl,omitempty"` // in seconds
}

type KafkaConfig struct {
	Brokers                         string `json:"brokers,omitempty"`
	UsernameSASL                    string `json:"username_sasl,omitempty"`
	PaswordSASL                     string `json:"pasword_sasl,omitempty"`
	Ca                              string `json:"ca,omitempty"`
	UsersEventTopic                 string `json:"users_event_topic,omitempty"`
	UsersEventTopicNumPartition     int    `json:"users_event_topic_num_partition,omitempty"`
	UsersEventTopicReplicationFator int    `json:"users_event_topic_replication_fator,omitempty"`
}

type ClickHouseConfig struct {
}

func Init(cfgPath string) (Config, error) {
	b, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
