package mocks

import "github.com/stretchr/testify/mock"

type MockApp struct {
	mock.Mock
  name string
}

func (mockApp *MockApp) Name() string {
 // args := mockApp.Called()
  return mockApp.name
}

func NewMock(name string) *MockApp{
  m := &MockApp{
    name: name,
  }
  m.setupExpectations("Name",name)
  return m
}

func (mockApp *MockApp) setupExpectations(methodName string,returnArgs ...interface{}) *mock.Call{
    // setup expectations
   return mockApp.On(methodName).Return(returnArgs)
}