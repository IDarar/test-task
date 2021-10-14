package transport

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IDarar/test-task/internal/config"
	"github.com/IDarar/test-task/internal/domain"
	p "github.com/IDarar/test-task/test_task"
	"google.golang.org/grpc"
)

type UserService interface {
	Create(name string) (*p.User, error)
	UsersList() (*p.UsersListResp, error)
	Delete(id int) error
}

type UserServer struct {
	Service UserService
}

func RunUsersServer(cfg config.GRPCConfig, services UserService) error {
	s := &UserServer{Service: services}

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		return fmt.Errorf("listen failed: %w", err)
	}

	grpcSrv := grpc.NewServer()

	p.RegisterUsersServer(grpcSrv, s)

	sigCh := make(chan os.Signal, 1)

	//get signal from os to gracefully shutdown server
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		<-sigCh
		grpcSrv.GracefulStop()
		wg.Done()
	}()

	fmt.Println("Starting server ...")
	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("failed serving: %w", err)
	}

	wg.Wait()

	return nil
}

func (s *UserServer) Create(ctx context.Context, req *p.CreateUserReq) (*p.User, error) {
	if req.Name == "" {
		return nil, domain.ErrEmptyName
	}

	return s.Service.Create(req.Name)
}

func (s *UserServer) UsersList(context.Context, *p.UsersListReq) (*p.UsersListResp, error) {
	return s.Service.UsersList()
}

func (s *UserServer) Delete(ctx context.Context, user *p.DeleteUserReq) (*p.DeleteUserResp, error) {
	if user.ID == 0 {
		return &p.DeleteUserResp{}, domain.ErrIdNotSet
	}

	return &p.DeleteUserResp{}, s.Service.Delete(int(user.ID))
}
