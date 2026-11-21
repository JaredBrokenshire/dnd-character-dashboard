package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateSubclassesTable struct{}

func (m *CreateSubclassesTable) GetName() string {
	return "CreateSubclassesTable"
}

func (m *CreateSubclassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("subclasses", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.Column("class_id").Type("INT UNSIGNED").NotNull()
	table.String("name", 255).NotNull()
	table.Column("description").Type("TEXT").NotNull()
	table.MustExec()
}

func (m *CreateSubclassesTable) Down(con *sqlx.DB) {
	builder.DropTable("subclasses", con).MustExec()
}
