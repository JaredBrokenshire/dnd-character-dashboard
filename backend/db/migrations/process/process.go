package process

import (
	"dnd/db/migrations/list"
	go_migrations "dnd/pkg/go-migrations"
	"dnd/pkg/go-migrations/store"
)

func Run() {
	go_migrations.Run(List())
}

func List() []store.Migratable {
	return []store.Migratable{
		&list.CreateRolesTable{},
		&list.CreateUsersTable{},
		&list.CreateUserRolesTable{},
	}
}
