package mocks

import (
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type MockCharacterService struct {
	ListCalls []struct {
		UserID interface{}
		scopes []func(db *gorm.DB) *gorm.DB
	}

	CountCalls []int

	GetCalls []struct {
		UserID interface{}
		ID     interface{}
	}
	GetError error
}

func NewMockCharacterService() *MockCharacterService {
	return &MockCharacterService{
		ListCalls: []struct {
			UserID interface{}
			scopes []func(db *gorm.DB) *gorm.DB
		}{},
		CountCalls: []int{},

		GetCalls: []struct {
			UserID interface{}
			ID     interface{}
		}{},
		GetError: nil,
	}
}

func (m *MockCharacterService) List(_ echo.Context, userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Character, int, int) {
	m.ListCalls = append(m.ListCalls, struct {
		UserID interface{}
		scopes []func(db *gorm.DB) *gorm.DB
	}{
		UserID: userID,
		scopes: scopes,
	})

	return nil, 0, 0
}

func (m *MockCharacterService) Count(_ interface{}, _ []func(db *gorm.DB) *gorm.DB) int {
	return 0
}

func (m *MockCharacterService) Get(userId, id interface{}) (*models.Character, error) {
	m.GetCalls = append(
		[]struct {
			UserID interface{}
			ID     interface{}
		}{
			{
				UserID: userId,
				ID:     id,
			},
		},
	)

	return nil, m.GetError
}

func (m *MockCharacterService) Reset() {
	m.ListCalls = []struct {
		UserID interface{}
		scopes []func(db *gorm.DB) *gorm.DB
	}{}
	m.CountCalls = []int{}

	m.GetCalls = []struct {
		UserID interface{}
		ID     interface{}
	}{}
	m.GetError = nil
}
