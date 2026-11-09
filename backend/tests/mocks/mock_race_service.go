package mocks

import (
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
)

type MockRaceService struct {
	ListCalls int

	CountCalls int

	GetCalls []interface{}
	GetError error
}

func NewMockRaceService() *MockRaceService {
	return &MockRaceService{
		ListCalls: 0,

		CountCalls: 0,

		GetCalls: []interface{}{},
		GetError: nil,
	}
}

func (m *MockRaceService) List(_, _ int, _ []func(db *gorm.DB) *gorm.DB) ([]*models.Race, int, int) {
	m.ListCalls++
	return nil, 0, 0
}

func (m *MockRaceService) Count(_ []func(db *gorm.DB) *gorm.DB) int {
	m.CountCalls++
	return 0
}

func (m *MockRaceService) Get(id interface{}) (*models.Race, error) {
	m.GetCalls = append(m.GetCalls, id)
	return nil, m.GetError
}

func (m *MockRaceService) Reset() {
	m.ListCalls = 0
	m.CountCalls = 0

	m.GetCalls = []interface{}{}
	m.GetError = nil
}
