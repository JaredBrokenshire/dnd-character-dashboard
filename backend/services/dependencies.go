package services

import (
	"dnd-character-dashboard/db/repositories"
	"dnd-character-dashboard/pkg/file_storage"
	"github.com/jinzhu/gorm"
	"os"
)

type Dependencies struct {
	db        *gorm.DB
	repos     *repositories.Repos
	jwtSecret string

	// File Storage
	fileStore file_storage.Store

	// Services
	authService AuthServiceInterface
	userService UserServiceInterface
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

func (d *Dependencies) GetFileStore() file_storage.Store {
	if d.fileStore == nil {
		d.fileStore = file_storage.NewLocalStorage("assets/files")
	}

	return d.fileStore
}

func (d *Dependencies) SetFileStore(fileStore file_storage.Store) {
	d.fileStore = fileStore
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

func (d *Dependencies) GetUserService() UserServiceInterface {
	if d.userService == nil {
		d.userService = NewUserService(d.repos.User)
	}

	return d.userService
}

func (d *Dependencies) SetUserService(userService UserServiceInterface) {
	d.userService = userService
}
