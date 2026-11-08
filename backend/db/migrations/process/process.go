package process

import (
	"dnd-character-dashboard/db/migrations/list"
	go_migrations "dnd-character-dashboard/pkg/go-migrations"
	"dnd-character-dashboard/pkg/go-migrations/store"
)

func Run() {
	go_migrations.Run(List())
}

func List() []store.Migratable {
	return []store.Migratable{
		&list.CreateUsersTable{},
		&list.CreateRolesTable{},
		&list.CreateUserRolesTable{},
	}
}
