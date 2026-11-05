package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateTraitsTable struct{}

func (m *CreateTraitsTable) GetName() string {
	return "CreateTraitsTable"
}

func (m *CreateTraitsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("traits", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.MustExec()
}

func (m *CreateTraitsTable) Down(con *sqlx.DB) {
	builder.DropTable("traits", con).MustExec()
}
