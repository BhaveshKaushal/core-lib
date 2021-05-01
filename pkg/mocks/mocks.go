package mocks

import "github.com/stretchr/testify/mock"

type MockApp struct {
	mock.Mock
}

func (mockApp *MockApp) Name() string {
  args := mockApp.Mock.Called()
  return args.String(0)
}