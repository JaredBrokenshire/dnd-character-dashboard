package mocks

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/services"
)

type MockAuthService struct {
	RegisterCalls []services.RegisterRequestDTO
	RegisterError error

	GenerateAccessTokenCalls []*models.User
	GenerateAccessTokenError error

	GenerateRefreshTokenCalls []*models.User
	GenerateRefreshTokenError error
}

func NewMockAuthService() *MockAuthService {
	return &MockAuthService{
		RegisterCalls: []services.RegisterRequestDTO{},
		RegisterError: nil,

		GenerateAccessTokenCalls: []*models.User{},
		GenerateAccessTokenError: nil,

		GenerateRefreshTokenCalls: []*models.User{},
		GenerateRefreshTokenError: nil,
	}
}

func (m *MockAuthService) Register(dto services.RegisterRequestDTO) (*services.AuthResponse, string, error) {
	m.RegisterCalls = append(m.RegisterCalls, dto)
	return nil, "", m.RegisterError
}

func (m *MockAuthService) GenerateAccessToken(user *models.User) (string, error) {
	m.GenerateAccessTokenCalls = append(m.GenerateAccessTokenCalls, user)
	return "", m.GenerateAccessTokenError
}

func (m *MockAuthService) GenerateRefreshToken(user *models.User) (string, error) {
	m.GenerateRefreshTokenCalls = append(m.GenerateRefreshTokenCalls, user)
	return "", m.GenerateRefreshTokenError
}

func (m *MockAuthService) Reset() {
	m.RegisterCalls = []services.RegisterRequestDTO{}
	m.RegisterError = nil

	m.GenerateAccessTokenCalls = []*models.User{}
	m.GenerateAccessTokenError = nil

	m.GenerateRefreshTokenCalls = []*models.User{}
	m.GenerateRefreshTokenError = nil
}
