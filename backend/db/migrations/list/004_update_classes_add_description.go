package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassesAddDescription struct{}

func (m *UpdateClassesAddDescription) GetName() string {
	return "UpdateClassesAddDescription"
}

func (m *UpdateClassesAddDescription) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.Column("description").Type("LONGTEXT").NotNull()
	table.MustExec()
}

func (m *UpdateClassesAddDescription) Down(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.DropColumn("description")
	table.MustExec()
}
