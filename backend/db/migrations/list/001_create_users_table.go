package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateUsersTable struct{}

func (m *CreateUsersTable) GetName() string {
	return "CreateUsersTable"
}

func (m *CreateUsersTable) Up(con *sqlx.DB) {
	table := builder.NewTable("users", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.String("username", 255).NotNull().Unique()
	table.String("password", 255).NotNull()
	table.WithTimestamps()
	table.Column("deleted_at").Type("DATETIME").Nullable()
	table.MustExec()
}

func (m *CreateUsersTable) Down(con *sqlx.DB) {
	builder.DropTable("users", con).MustExec()
}
