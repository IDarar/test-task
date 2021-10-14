package service

import (
	"errors"
	"testing"

	"github.com/IDarar/test-task/internal/domain"
	"github.com/IDarar/test-task/internal/service/mock_service"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
)

func TestUsersList(t *testing.T) {
	tests := []struct {
		name     string
		cacheErr error
	}{
		{
			name:     "list is cached",
			cacheErr: nil,
		},
		{
			name:     "cache is empty",
			cacheErr: errors.New("empty cache"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			defer c.Finish()

			cache := mock_service.NewMockUserCache(c)
			repo := mock_service.NewMockUsersRepo(c)

			cache.EXPECT().GetList().Return([]domain.User{}, tc.cacheErr)
			if tc.cacheErr != nil {
				repo.EXPECT().List().Return([]domain.User{}, nil)
			}

			s := UsersService{cache: cache, repo: repo}

			_, err := s.UsersList()
			require.NoError(t, err)
		})
	}
}
