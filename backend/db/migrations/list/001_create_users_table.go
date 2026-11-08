package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateUsersTable struct{}

func (m *CreateUsersTable) GetName() string {
	return "CreateUsersTable"
}

func (m *CreateUsersTable) Up(con *sqlx.DB) {
	table := builder.NewTable("users", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("username", 255).NotNull().Unique()
	table.String("email", 255).NotNull().Unique()
	table.String("password", 255).NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateUsersTable) Down(con *sqlx.DB) {
	builder.DropTable("users", con).MustExec()
}
