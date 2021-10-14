package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/IDarar/test-task/internal/domain"
	"github.com/go-redis/redis/v8"
)

type UsersCache struct {
	rdb          *redis.Client
	usersListTTL time.Duration
}

func NewUsersCache(usersListTTL int, rdb *redis.Client) *UsersCache {
	return &UsersCache{
		rdb:          rdb,
		usersListTTL: time.Second * time.Duration(usersListTTL),
	}
}

func (r *UsersCache) GetList() ([]domain.User, error) {
	bs, err := r.rdb.Get(context.Background(), "list").Bytes()
	if err != nil {
		return []domain.User{}, fmt.Errorf("failed get list: %w", err)
	}

	userRes := []domain.User{}

	bytesRes := bytes.NewReader(bs)

	err = gob.NewDecoder(bytesRes).Decode(&userRes)
	if err != nil {
		return userRes, fmt.Errorf("failed to decode users list: %w", err)
	}

	return userRes, nil
}

func (r *UsersCache) SetList(uList []domain.User) error {
	var b bytes.Buffer

	err := gob.NewEncoder(&b).Encode(uList)
	if err != nil {
		return fmt.Errorf("failed to encode users list: %w", err)
	}

	err = r.rdb.Set(context.Background(), "list", b.Bytes(), r.usersListTTL).Err()
	if err != nil {
		return fmt.Errorf("failed set cache: %w", err)
	}

	return nil
}
