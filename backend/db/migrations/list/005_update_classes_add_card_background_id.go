package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassesAddCardBackgroundFileId struct{}

func (m *UpdateClassesAddCardBackgroundFileId) GetName() string {
	return "UpdateClassesAddCardBackgroundFileId"
}

func (m *UpdateClassesAddCardBackgroundFileId) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.Column("card_background_file_id").Type("int unsigned").Nullable()
	table.MustExec()
}

func (m *UpdateClassesAddCardBackgroundFileId) Down(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.DropColumn("card_background_file_id")
	table.MustExec()
}
