package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateFilesTable struct{}

func (m *CreateFilesTable) GetName() string {
	return "CreateFilesTable"
}

func (m *CreateFilesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("files", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.String("name", 255).NotNull()
	table.Column("location").Type("TEXT").NotNull()
	table.WithTimestamps()
	table.Column("deleted_at").Type("DATETIME").Nullable()
	table.MustExec()
}

func (m *CreateFilesTable) Down(con *sqlx.DB) {
	builder.DropTable("files", con).MustExec()
}
