package mq

import (
	"encoding/json"
	"fmt"

	"github.com/IDarar/test-task/internal/domain"
	"github.com/Shopify/sarama"
)

type UsersKafka struct {
	producer        sarama.SyncProducer
	usersEventTopic string
}

func NewUsersKafka(p sarama.SyncProducer, usersEventTopic string) *UsersKafka {
	return &UsersKafka{
		producer:        p,
		usersEventTopic: usersEventTopic,
	}
}

func (k *UsersKafka) SendCreationEvent(user domain.User) {
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("err marshaling user object: ", err)
		return
	}

	_, _, err = k.producer.SendMessage(&sarama.ProducerMessage{Topic: k.usersEventTopic, Value: sarama.StringEncoder(b)})
	if err != nil {
		fmt.Println("err sending message to kafka: ", err)
	}
}
