package services

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/db/repositories"
	"dnd-character-dashboard/pkg/middleware"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type AuthServiceInterface interface {
	Register(dto RegisterRequestDTO) (*AuthResponse, string, error)
	Login(dto LoginRequestDTO) (*AuthResponse, string, error)
	Refresh(refreshToken string) (*AuthResponse, error)
	GenerateAccessToken(user *models.User) (string, error)
	GenerateRefreshToken(user *models.User) (string, error)
}

type AuthResponse struct {
	AccessToken string
	User        *models.User
}

type RegisterRequestDTO struct {
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
}

type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Common sentinel errors.
var (
	ErrEmailTaken       = errors.New("email already registered")
	ErrUsernameTaken    = errors.New("username already taken")
	ErrInvalidCreds     = errors.New("invalid username or password")
	ErrUserNotFound     = errors.New("user not found")
	ErrPasswordNotMatch = errors.New("passwords do not match")
)

// AuthService handles user registration, login and token operations.
type AuthService struct {
	db        *gorm.DB
	userRepo  *repositories.UserRepository
	roleRepo  *repositories.RoleRepository
	jwtSecret []byte
}

// NewAuthService creates a new AuthService.
func NewAuthService(db *gorm.DB, userRepo *repositories.UserRepository, roleRepo *repositories.RoleRepository, jwtSecret string) *AuthService {
	return &AuthService{
		db:        db,
		userRepo:  userRepo,
		roleRepo:  roleRepo,
		jwtSecret: []byte(jwtSecret),
	}
}

// Register creates a new user account with the default "user" role.
func (s *AuthService) Register(dto RegisterRequestDTO) (*AuthResponse, string, error) {
	if dto.Password != dto.ConfirmPassword {
		return nil, "", ErrPasswordNotMatch
	}

	existingUser := s.userRepo.GetByUsername(dto.Username)
	if existingUser.ID != 0 {
		return nil, "", ErrUsernameTaken
	}

	existingEmail := s.userRepo.GetByEmail(dto.Email)
	if existingEmail.ID != 0 {
		return nil, "", ErrEmailTaken
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	role := s.roleRepo.GetByName(models.RoleUser)

	user := &models.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: string(passwordHash),
		Roles:    []models.Role{*role},
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, "", err
	}

	accessToken, err := s.GenerateAccessToken(user)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := s.GenerateRefreshToken(user)
	if err != nil {
		return nil, "", err
	}

	return &AuthResponse{
		AccessToken: accessToken,
		User:        user,
	}, refreshToken, nil

}

// Login authenticates an existing user by username and returns a signed JWT.
func (s *AuthService) Login(dto LoginRequestDTO) (*AuthResponse, string, error) {
	user := s.userRepo.GetByUsername(dto.Username)
	if user.ID == 0 {
		return nil, "", ErrInvalidCreds
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return nil, "", ErrInvalidCreds
	}

	accessToken, err := s.GenerateAccessToken(user)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := s.GenerateRefreshToken(user)
	if err != nil {
		return nil, "", err
	}

	return &AuthResponse{
		AccessToken: accessToken,
		User:        user,
	}, refreshToken, nil
}

// Refresh refreshes an existing token
func (s *AuthService) Refresh(refreshToken string) (*AuthResponse, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})
	if err != nil {
		return nil, ErrInvalidCreds
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidCreds
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return nil, ErrInvalidCreds
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, ErrInvalidCreds
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		return nil, ErrInvalidCreds
	}

	user := s.userRepo.GetById(userId)
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}

	accessToken, err := s.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken: accessToken,
		User:        user,
	}, nil
}

func (s *AuthService) GenerateAccessToken(user *models.User) (string, error) {
	claims := middleware.JWTClaims{
		UserID: user.ID,

		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "localdnd-api",
			Subject:  strconv.FormatUint(uint64(user.ID), 10),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(15 * time.Minute),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(s.jwtSecret)
}

func (s *AuthService) GenerateRefreshToken(user *models.User) (string, error) {

	claims := jwt.MapClaims{
		"sub":  strconv.FormatUint(uint64(user.ID), 10),
		"type": "refresh",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iss":  "localdnd-api",
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(s.jwtSecret)
}
