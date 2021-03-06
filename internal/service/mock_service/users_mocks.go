// Code generated by MockGen. DO NOT EDIT.
// Source: users.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	domain "github.com/IDarar/test-task/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockUsersRepo is a mock of UsersRepo interface.
type MockUsersRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepoMockRecorder
}

// MockUsersRepoMockRecorder is the mock recorder for MockUsersRepo.
type MockUsersRepoMockRecorder struct {
	mock *MockUsersRepo
}

// NewMockUsersRepo creates a new mock instance.
func NewMockUsersRepo(ctrl *gomock.Controller) *MockUsersRepo {
	mock := &MockUsersRepo{ctrl: ctrl}
	mock.recorder = &MockUsersRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepo) EXPECT() *MockUsersRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsersRepo) Create(name string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUsersRepoMockRecorder) Create(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsersRepo)(nil).Create), name)
}

// Delete mocks base method.
func (m *MockUsersRepo) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUsersRepoMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsersRepo)(nil).Delete), id)
}

// List mocks base method.
func (m *MockUsersRepo) List() ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUsersRepoMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUsersRepo)(nil).List))
}

// MockUserCache is a mock of UserCache interface.
type MockUserCache struct {
	ctrl     *gomock.Controller
	recorder *MockUserCacheMockRecorder
}

// MockUserCacheMockRecorder is the mock recorder for MockUserCache.
type MockUserCacheMockRecorder struct {
	mock *MockUserCache
}

// NewMockUserCache creates a new mock instance.
func NewMockUserCache(ctrl *gomock.Controller) *MockUserCache {
	mock := &MockUserCache{ctrl: ctrl}
	mock.recorder = &MockUserCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCache) EXPECT() *MockUserCacheMockRecorder {
	return m.recorder
}

// GetList mocks base method.
func (m *MockUserCache) GetList() ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList")
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockUserCacheMockRecorder) GetList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockUserCache)(nil).GetList))
}

// SetList mocks base method.
func (m *MockUserCache) SetList(arg0 []domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetList", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetList indicates an expected call of SetList.
func (mr *MockUserCacheMockRecorder) SetList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetList", reflect.TypeOf((*MockUserCache)(nil).SetList), arg0)
}

// MockUserKafka is a mock of UserKafka interface.
type MockUserKafka struct {
	ctrl     *gomock.Controller
	recorder *MockUserKafkaMockRecorder
}

// MockUserKafkaMockRecorder is the mock recorder for MockUserKafka.
type MockUserKafkaMockRecorder struct {
	mock *MockUserKafka
}

// NewMockUserKafka creates a new mock instance.
func NewMockUserKafka(ctrl *gomock.Controller) *MockUserKafka {
	mock := &MockUserKafka{ctrl: ctrl}
	mock.recorder = &MockUserKafkaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserKafka) EXPECT() *MockUserKafkaMockRecorder {
	return m.recorder
}

// SendCreationEvent mocks base method.
func (m *MockUserKafka) SendCreationEvent(user domain.User) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendCreationEvent", user)
}

// SendCreationEvent indicates an expected call of SendCreationEvent.
func (mr *MockUserKafkaMockRecorder) SendCreationEvent(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCreationEvent", reflect.TypeOf((*MockUserKafka)(nil).SendCreationEvent), user)
}
