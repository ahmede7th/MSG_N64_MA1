package main

import (
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateStuff(stuff *Stuff) error {

	rets := m.Called(stuff)
	return rets.Error(0)
}

func (m *MockStore) GetStuffs() ([]*Stuff, error) {
	rets := m.Called()

	return rets.Get(0).([]*Stuff), rets.Error(1)
}

func InitMockStore() *MockStore {

	s := new(MockStore)
	store = s
	return s
}
