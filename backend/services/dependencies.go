package services

import (
	"dnd-character-dashboard/db/repositories"
	"github.com/jinzhu/gorm"
	"os"
)

type Dependencies struct {
	db        *gorm.DB
	repos     *repositories.Repos
	jwtSecret string

	// Services
	authService AuthServiceInterface
}

func NewDependencies(db *gorm.DB, repos *repositories.Repos) *Dependencies {
	return &Dependencies{
		db:        db,
		repos:     repos,
		jwtSecret: os.Getenv("JWT_SECRET"),
	}
}

func (d *Dependencies) GetDB() *gorm.DB {
	return d.db
}

func (d *Dependencies) GetAuthService() AuthServiceInterface {
	if d.authService == nil {
		d.authService = NewAuthService(d.db, d.repos.User, d.repos.Role, d.jwtSecret)
	}

	return d.authService
}

func (d *Dependencies) SetAuthService(authService AuthServiceInterface) {
	d.authService = authService
}
