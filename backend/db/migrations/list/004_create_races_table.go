package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRacesTable struct{}

func (m *CreateRacesTable) GetName() string {
	return "CreateRacesTable"
}

func (m *CreateRacesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("races", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.Column("card_description").Type("TEXT").NotNull()
	table.Column("description").Type("LONGTEXT").NotNull()
	table.String("creature_type", 255).NotNull()
	table.String("size", 255).NotNull()
	table.Integer("base_speed").NotNull()
	table.Column("card_background_file_id").Type("INT UNSIGNED").NotNull()
	table.Column("details_image_file_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *CreateRacesTable) Down(con *sqlx.DB) {
	builder.DropTable("races", con).MustExec()
}
