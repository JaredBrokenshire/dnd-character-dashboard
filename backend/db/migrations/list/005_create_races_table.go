package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRacesTable struct{}

func (m *CreateRacesTable) GetName() string {
	return "CreateRacesTable"
}

func (m *CreateRacesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("races", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.String("name", 255).NotNull()
	table.Column("card_description").Type("TEXT").NotNull()
	table.Column("card_background_file_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *CreateRacesTable) Down(con *sqlx.DB) {
	builder.DropTable("races", con).MustExec()
}
