package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRacesTable struct{}

func (m *CreateRacesTable) GetName() string {
	return "CreateRacesTable"
}

func (m *CreateRacesTable) Up(conn *sqlx.DB) {
	table := builder.NewTable("races", conn)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.MustExec()
}

func (m *CreateRacesTable) Down(conn *sqlx.DB) {
	builder.DropTable("races", conn).MustExec()
}
