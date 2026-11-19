package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRolesTable struct{}

func (m *CreateRolesTable) GetName() string {
	return "CreateRolesTable"
}

func (m *CreateRolesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("roles", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.String("name", 255).NotNull().Unique()
	table.MustExec()
}

func (m *CreateRolesTable) Down(con *sqlx.DB) {
	builder.DropTable("roles", con).MustExec()
}
