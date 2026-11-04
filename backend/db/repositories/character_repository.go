package repositories

import (
	"dnd-character-dashboard/db"
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type CharacterRepository struct {
	*Repository
}

func NewCharacterRepository(db *gorm.DB) *CharacterRepository {
	return &CharacterRepository{
		Repository: &Repository{
			Db: db,
		},
	}
}

func (r *CharacterRepository) GetCharacters(c echo.Context, userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Character, int, int) {
	var characters []*models.Character
	page, pageSize, paginateFunc := db.Paginate(c)

	r.Db.
		Preload("ProfilePicture").
		Preload("Race").
		Preload("Classes").
		Preload("Classes.Logo").
		Preload("Subclasses").
		Preload("Campaign").
		Scopes(paginateFunc).
		Scopes(scopes...).
		Where("user_id = ?", userID).
		Order("name ASC").
		Find(&characters)

	return characters, page, pageSize
}

func (r *CharacterRepository) Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int {
	var count int64

	r.Db.
		Model(&models.Character{}).
		Scopes(scopes...).
		Where("user_id = ?", userID).
		Count(&count)

	return int(count)
}

func (r *CharacterRepository) GetByID(userID, id interface{}) *models.Character {
	var character models.Character

	r.Db.
		Preload("ProfilePicture").
		Preload("Race").
		Preload("Classes").
		Preload("Classes.Logo").
		Preload("Subclasses").
		Preload("Campaign").
		Where("user_id = ?", userID).
		Where("id = ?", id).
		First(&character)

	return &character
}
