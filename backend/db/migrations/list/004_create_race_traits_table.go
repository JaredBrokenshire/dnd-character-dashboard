package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRaceTraitsTable struct{}

func (m *CreateRaceTraitsTable) GetName() string {
	return "CreateRaceTraitsTable"
}

func (m *CreateRaceTraitsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("race_traits", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("race_id").Type("INT UNSIGNED").NotNull()
	table.Column("trait_id").Type("INT UNSIGNED").NotNull()

	table.ForeignKey("race_id").Reference("races").On("id").OnDelete("CASCADE").OnUpdate("CASCADE")
	table.ForeignKey("trait_id").Reference("traits").On("id").OnDelete("CASCADE").OnUpdate("CASCADE")

	table.MustExec()
}

func (m *CreateRaceTraitsTable) Down(con *sqlx.DB) {
	builder.DropTable("race_traits", con).MustExec()
}
