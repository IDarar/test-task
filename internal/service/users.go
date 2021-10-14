package service

import (
	"github.com/IDarar/test-task/internal/domain"
	p "github.com/IDarar/test-task/test_task"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:generate mockgen -source=users.go -destination=mock_service/users_mocks.go

type UsersRepo interface {
	Create(name string) (domain.User, error)
	Delete(id int) error
	List() ([]domain.User, error)
}

type UserCache interface {
	SetList([]domain.User) error
	GetList() ([]domain.User, error)
}

type UserKafka interface {
	SendCreationEvent(user domain.User)
}

type UsersService struct {
	repo  UsersRepo
	cache UserCache
	kafka UserKafka
}

func NewUsersService(repo UsersRepo, cache UserCache, kafka UserKafka) *UsersService {
	return &UsersService{
		repo:  repo,
		cache: cache,
		kafka: kafka,
	}
}

func (s *UsersService) Create(name string) (*p.User, error) {
	user, err := s.repo.Create(name)
	if err != nil {
		return &p.User{}, err
	}

	go s.kafka.SendCreationEvent(user)

	return &p.User{Name: user.Name, ID: int64(user.ID), CreatedAt: timestamppb.New(user.CreatedAt)}, nil
}

func (s *UsersService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *UsersService) UsersList() (*p.UsersListResp, error) {
	//get from cache
	users, err := s.cache.GetList()

	//if no err, return resluts from cache
	if err == nil {
		resp := &p.UsersListResp{}
		for _, user := range users {
			resp.Users = append(resp.Users, &p.User{ID: int64(user.ID), Name: user.Name, CreatedAt: timestamppb.New(user.CreatedAt)})
		}

		return resp, nil
	}

	//if was err, get from db
	users, err = s.repo.List()
	if err != nil {
		return &p.UsersListResp{}, err
	}

	//set cache
	go s.cache.SetList(users)

	resp := &p.UsersListResp{}
	for _, user := range users {
		resp.Users = append(resp.Users, &p.User{ID: int64(user.ID), Name: user.Name, CreatedAt: timestamppb.New(user.CreatedAt)})
	}

	return resp, nil
}
