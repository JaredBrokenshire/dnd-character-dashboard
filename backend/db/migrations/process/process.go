package process

import (
	go_migrations "dnd/pkg/go-migrations"
	"dnd/pkg/go-migrations/store"
)

func Run() {
	go_migrations.Run(List())
}

func List() []store.Migratable {
	return []store.Migratable{}
}
