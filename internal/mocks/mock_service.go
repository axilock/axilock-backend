package mocks

import mock "github.com/stretchr/testify/mock"

type MockServiceStore struct {
	User UserServiceInterface
}

func NewMockServiceStore(t interface {
	mock.TestingT
	Cleanup(func())
},
) *MockServiceStore {
	return &MockServiceStore{
		User: *NewUserServiceInterface(t),
	}
}
