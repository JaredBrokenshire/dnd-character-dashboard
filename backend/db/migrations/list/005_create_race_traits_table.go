package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRaceTraitsTable struct{}

func (m *CreateRaceTraitsTable) GetName() string {
	return "CreateRaceTraitsTable"
}

func (m *CreateRaceTraitsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("race_traits", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.Column("race_id").Type("INT UNSIGNED").NotNull()
	table.Column("trait_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *CreateRaceTraitsTable) Down(con *sqlx.DB) {
	builder.DropTable("race_traits", con).MustExec()
}
