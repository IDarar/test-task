package cache

import (
	"testing"
	"time"

	"github.com/IDarar/test-task/internal/config"
	"github.com/IDarar/test-task/internal/domain"
	"github.com/IDarar/test-task/pkg/redisdb"
	"github.com/stretchr/testify/require"
)

//test there both methods
func TestCache(t *testing.T) {
	usersCache, err := getTestUserCache()
	require.NoError(t, err)

	uList := []domain.User{}

	for i := 0; i < 10; i++ {
		uList = append(uList, domain.User{ID: i, Name: "name", CreatedAt: time.Now()})
	}

	err = usersCache.SetList(uList)
	require.NoError(t, err)

	cachedList, err := usersCache.GetList()
	require.NoError(t, err)

	require.Equal(t, len(uList), len(cachedList))

	for i, v := range uList {
		require.Equal(t, v.ID, cachedList[i].ID)
		require.Equal(t, v.Name, cachedList[i].Name)
		require.Equal(t, v.CreatedAt.Hour(), cachedList[i].CreatedAt.Hour())
		require.Equal(t, v.CreatedAt.Minute(), cachedList[i].CreatedAt.Minute())
		require.Equal(t, v.CreatedAt.Second(), cachedList[i].CreatedAt.Second())
	}

	//test if cache was deleted
	time.Sleep(500 * time.Millisecond)

	_, err = usersCache.GetList()
	require.Error(t, err)
}

func getTestUserCache() (*UsersCache, error) {
	rdb, err := redisdb.NewRedisDB(config.Config{Redis: config.RedisConfig{Addr: "localhost:6379", Password: "secret", DB: 0}})
	if err != nil {
		return nil, err
	}

	return &UsersCache{
		rdb:          rdb,
		usersListTTL: 500 * time.Millisecond,
	}, nil
}
