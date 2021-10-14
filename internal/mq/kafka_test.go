package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/IDarar/test-task/internal/config"
	"github.com/IDarar/test-task/internal/domain"
	"github.com/IDarar/test-task/pkg/kafka"
	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/require"
)

var (
	testUsersCreationEventTopic = "hoq0cmws-user_creation_test"

	// almost same config, because I have only one kafka cluster
	testKafkaCfg = config.KafkaConfig{Brokers: "glider-01.srvs.cloudkafka.com:9094,glider-02.srvs.cloudkafka.com:9094,glider-03.srvs.cloudkafka.com:9094",
		UsernameSASL:                    "hoq0cmws",
		PaswordSASL:                     "4zNgnuNsWbb3MkbjC3zvNxG1rUTG4YHb",
		Ca:                              "-----BEGIN CERTIFICATE-----\nMIIF3jCCA8agAwIBAgIQAf1tMPyjylGoG7xkDjUDLTANBgkqhkiG9w0BAQwFADCB\niDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCk5ldyBKZXJzZXkxFDASBgNVBAcTC0pl\ncnNleSBDaXR5MR4wHAYDVQQKExVUaGUgVVNFUlRSVVNUIE5ldHdvcmsxLjAsBgNV\nBAMTJVVTRVJUcnVzdCBSU0EgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMTAw\nMjAxMDAwMDAwWhcNMzgwMTE4MjM1OTU5WjCBiDELMAkGA1UEBhMCVVMxEzARBgNV\nBAgTCk5ldyBKZXJzZXkxFDASBgNVBAcTC0plcnNleSBDaXR5MR4wHAYDVQQKExVU\naGUgVVNFUlRSVVNUIE5ldHdvcmsxLjAsBgNVBAMTJVVTRVJUcnVzdCBSU0EgQ2Vy\ndGlmaWNhdGlvbiBBdXRob3JpdHkwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK\nAoICAQCAEmUXNg7D2wiz0KxXDXbtzSfTTK1Qg2HiqiBNCS1kCdzOiZ/MPans9s/B\n3PHTsdZ7NygRK0faOca8Ohm0X6a9fZ2jY0K2dvKpOyuR+OJv0OwWIJAJPuLodMkY\ntJHUYmTbf6MG8YgYapAiPLz+E/CHFHv25B+O1ORRxhFnRghRy4YUVD+8M/5+bJz/\nFp0YvVGONaanZshyZ9shZrHUm3gDwFA66Mzw3LyeTP6vBZY1H1dat//O+T23LLb2\nVN3I5xI6Ta5MirdcmrS3ID3KfyI0rn47aGYBROcBTkZTmzNg95S+UzeQc0PzMsNT\n79uq/nROacdrjGCT3sTHDN/hMq7MkztReJVni+49Vv4M0GkPGw/zJSZrM233bkf6\nc0Plfg6lZrEpfDKEY1WJxA3Bk1QwGROs0303p+tdOmw1XNtB1xLaqUkL39iAigmT\nYo61Zs8liM2EuLE/pDkP2QKe6xJMlXzzawWpXhaDzLhn4ugTncxbgtNMs+1b/97l\nc6wjOy0AvzVVdAlJ2ElYGn+SNuZRkg7zJn0cTRe8yexDJtC/QV9AqURE9JnnV4ee\nUB9XVKg+/XRjL7FQZQnmWEIuQxpMtPAlR1n6BB6T1CZGSlCBst6+eLf8ZxXhyVeE\nHg9j1uliutZfVS7qXMYoCAQlObgOK6nyTJccBz8NUvXt7y+CDwIDAQABo0IwQDAd\nBgNVHQ4EFgQUU3m/WqorSs9UgOHYm8Cd8rIDZsswDgYDVR0PAQH/BAQDAgEGMA8G\nA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQEMBQADggIBAFzUfA3P9wF9QZllDHPF\nUp/L+M+ZBn8b2kMVn54CVVeWFPFSPCeHlCjtHzoBN6J2/FNQwISbxmtOuowhT6KO\nVWKR82kV2LyI48SqC/3vqOlLVSoGIG1VeCkZ7l8wXEskEVX/JJpuXior7gtNn3/3\nATiUFJVDBwn7YKnuHKsSjKCaXqeYalltiz8I+8jRRa8YFWSQEg9zKC7F4iRO/Fjs\n8PRF/iKz6y+O0tlFYQXBl2+odnKPi4w2r78NBc5xjeambx9spnFixdjQg3IM8WcR\niQycE0xyNN+81XHfqnHd4blsjDwSXWXavVcStkNr/+XeTWYRUc+ZruwXtuhxkYze\nSf7dNXGiFSeUHM9h4ya7b6NnJSFd5t0dCy5oGzuCr+yDZ4XUmFF0sbmZgIn/f3gZ\nXHlKYC6SQK5MNyosycdiyA5d9zZbyuAlJQG03RoHnHcAP9Dc1ew91Pq7P8yF1m9/\nqS3fuQL39ZeatTXaw2ewh0qpKJ4jjv9cJ2vhsE/zB+4ALtRZh8tSQZXq9EfX7mRB\nVXyNWQKV3WKdwrnuWih0hKWbt5DHDAff9Yk2dDLWKMGwsAvgnEzDHNb842m1R0aB\nL6KCq9NjRHDEjf8tM7qtj3u1cIiuPhnPQCjY/MiQu12ZIvVS5ljFH4gxQ+6IHdfG\njjxDah2nGN59PRbxYvnKkKj9\n-----END CERTIFICATE-----",
		UsersEventTopic:                 testUsersCreationEventTopic,
		UsersEventTopicNumPartition:     1,
		UsersEventTopicReplicationFator: 1,
	}
)

func TestSendCreationEvent(t *testing.T) {
	userKafka, err := getTestUsersKafka()
	require.NoError(t, err)

	defer userKafka.producer.Close()

	//empty topic before tests
	err = purgeTestTopic()
	require.NoError(t, err)

	user := domain.User{ID: 1, Name: fmt.Sprint(time.Now().Second()), CreatedAt: time.Now()}

	//send message
	userKafka.SendCreationEvent(user)

	//check if it is on topic
	c, err := getTestConsumer()
	require.NoError(t, err)

	defer c.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	select {
	case err := <-c.Errors():
		t.Fatal(err)
	case msg := <-c.Messages():
		consumedUser := domain.User{}

		err = json.Unmarshal(msg.Value, &consumedUser)
		require.NoError(t, err)

		//time is not correct on decoding, so compare field by field
		require.Equal(t, user.ID, consumedUser.ID)
		require.Equal(t, user.Name, consumedUser.Name)

		require.Equal(t, user.CreatedAt.Hour(), consumedUser.CreatedAt.Hour())
		require.Equal(t, user.CreatedAt.Minute(), consumedUser.CreatedAt.Minute())
		require.Equal(t, user.CreatedAt.Second(), consumedUser.CreatedAt.Second())

	case <-ctx.Done():
		t.Fatal("timeout reading from kafka")
	}
}

func getTestUsersKafka() (*UsersKafka, error) {
	conf, err := kafka.GetKafkaConfig(testKafkaCfg)
	if err != nil {
		return nil, err
	}

	userKafka := &UsersKafka{}
	userKafka.usersEventTopic = testUsersCreationEventTopic

	p, err := kafka.NewKafkaProducer(testKafkaCfg, conf)
	if err != nil {
		return nil, err
	}

	userKafka.producer = p

	return userKafka, nil
}

func getTestConsumer() (sarama.PartitionConsumer, error) {
	conf, err := kafka.GetKafkaConfig(testKafkaCfg)
	if err != nil {
		return nil, err
	}

	brokers := strings.Split(testKafkaCfg.Brokers, ",")

	master, err := sarama.NewConsumer(brokers, conf)
	if err != nil {
		return nil, err
	}

	consumer, err := master.ConsumePartition(testKafkaCfg.UsersEventTopic, 0, sarama.OffsetOldest)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func purgeTestTopic() error {
	conf, err := kafka.GetKafkaConfig(testKafkaCfg)
	if err != nil {
		return err
	}

	ad, err := kafka.GetKafkaAdmin(testKafkaCfg, conf)
	if err != nil {
		return err
	}

	topicsOffsets := make(map[int32]int64)

	topicsOffsets[0] = -1

	return ad.DeleteRecords(testUsersCreationEventTopic, topicsOffsets)
}
