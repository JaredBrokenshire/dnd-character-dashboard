package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRaceVariantsTable struct{}

func (m *CreateRaceVariantsTable) GetName() string {
	return "CreateRaceVariantsTable"
}

func (m *CreateRaceVariantsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("race_variants", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("race_id").Type("INT UNSIGNED").NotNull()
	table.String("name", 255).NotNull()
	table.Column("description").Type("TEXT").NotNull()
	table.MustExec()
}

func (m *CreateRaceVariantsTable) Down(con *sqlx.DB) {
	builder.DropTable("race_variants", con).MustExec()
}
