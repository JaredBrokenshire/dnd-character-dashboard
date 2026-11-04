package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateSubclassesTable struct{}

func (m *CreateSubclassesTable) GetName() string {
	return "CreateSubclassesTable"
}

func (m *CreateSubclassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("subclasses", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("class_id").Type("int unsigned").NotNull()
	table.String("name", 255).NotNull()
	table.MustExec()
}

func (m *CreateSubclassesTable) Down(con *sqlx.DB) {
	builder.DropTable("subclasses", con).MustExec()
}
