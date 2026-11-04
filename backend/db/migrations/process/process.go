package process

import (
	"dnd-character-dashboard/db/migrations/list"
	"dnd-character-dashboard/pkg/go-migrations"
	"dnd-character-dashboard/pkg/go-migrations/store"
)

func Run() {
	go_migrations.Run(List())
}

func List() []store.Migratable {
	return []store.Migratable{
		&list.CreateRolesTable{},
		&list.CreateUsersTable{},
		&list.CreateUserRolesTable{},
		&list.CreateRacesTable{},
		&list.CreateClassesTable{},
		&list.CreateSubclassesTable{},
		&list.CreateCampaignsTable{},
		&list.CreateCharactersTable{},
		&list.CreateCharacterClassesTable{},
		&list.CreateCharacterSubclassesTable{},
		&list.CreateFilesTable{},
		&list.UpdateCharactersAddProfilePictureFileId{},
		&list.UpdateClassesAddLogoFileId{},
		&list.UpdateCampaignsAddUserId{},
		&list.UpdateClassesAddCardBackgroundFileId{},
		&list.UpdateClassesAddDescriptions{},
	}
}
