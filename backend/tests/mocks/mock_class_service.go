package mocks

import (
	"dnd/db"
	"dnd/db/models"
)

type ClassServiceListCalls struct {
	Page     int
	PageSize int
	Scopes   db.Scopes
}

type MockClassService struct {
	ListCalls []ClassServiceListCalls

	CountCalls []db.Scopes

	GetCalls []interface{}
	GetError error
}

func NewMockClassService() *MockClassService {
	m := &MockClassService{}
	m.Reset()
	return m
}

func (m *MockClassService) List(page, pageSize int, scopes db.Scopes) ([]*models.Class, int, int) {
	m.ListCalls = append(m.ListCalls, ClassServiceListCalls{page, pageSize, scopes})
	return nil, 0, 0
}

func (m *MockClassService) Count(scopes db.Scopes) int {
	m.CountCalls = append(m.CountCalls, scopes)
	return 0
}

func (m *MockClassService) Get(id interface{}) (*models.Class, error) {
	m.GetCalls = append(m.GetCalls, id)
	return nil, m.GetError
}

func (m *MockClassService) Reset() {
	m.ListCalls = []ClassServiceListCalls{}

	m.CountCalls = []db.Scopes{}

	m.GetCalls = []interface{}{}
	m.GetError = nil
}
