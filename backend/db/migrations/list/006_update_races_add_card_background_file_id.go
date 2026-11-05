package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateRacesAddCardFields struct{}

func (m *UpdateRacesAddCardFields) GetName() string {
	return "UpdateRacesAddCardFields"
}

func (m *UpdateRacesAddCardFields) Up(con *sqlx.DB) {
	table := builder.ChangeTable("races", con)
	table.Column("card_background_file_id").Type("int unsigned").Nullable()
	table.Column("card_description").Type("text").Nullable()
	table.MustExec()
}

func (m *UpdateRacesAddCardFields) Down(con *sqlx.DB) {
	table := builder.ChangeTable("races", con)
	table.DropColumn("card_background_file_id")
	table.MustExec()
}
