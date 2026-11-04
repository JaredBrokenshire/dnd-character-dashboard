package mocks

import (
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type MockCharacterService struct {
	GetCharactersCalls []struct {
		UserID interface{}
		scopes []func(db *gorm.DB) *gorm.DB
	}

	CountCalls []int

	GetCharacterByIDCalls []struct {
		UserID interface{}
		ID     interface{}
	}
	GetCharacterByIdError error
}

func NewMockCharacterService() *MockCharacterService {
	return &MockCharacterService{
		GetCharactersCalls: []struct {
			UserID interface{}
			scopes []func(db *gorm.DB) *gorm.DB
		}{},
		CountCalls: []int{},

		GetCharacterByIDCalls: []struct {
			UserID interface{}
			ID     interface{}
		}{},
		GetCharacterByIdError: nil,
	}
}

func (m *MockCharacterService) GetCharacters(_ echo.Context, userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Character, int, int) {
	m.GetCharactersCalls = append(m.GetCharactersCalls, struct {
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

func (m *MockCharacterService) GetCharacterByID(userId, id interface{}) (*models.Character, error) {
	m.GetCharacterByIDCalls = append(
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

	return nil, m.GetCharacterByIdError
}

func (m *MockCharacterService) Reset() {
	m.GetCharactersCalls = []struct {
		UserID interface{}
		scopes []func(db *gorm.DB) *gorm.DB
	}{}
	m.CountCalls = []int{}

	m.GetCharacterByIDCalls = []struct {
		UserID interface{}
		ID     interface{}
	}{}
	m.GetCharacterByIdError = nil
}
