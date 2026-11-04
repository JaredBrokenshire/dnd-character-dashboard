package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassesAddLogoFileId struct{}

func (m *UpdateClassesAddLogoFileId) GetName() string {
	return "UpdateClassesAddLogoFileId"
}

func (m *UpdateClassesAddLogoFileId) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.Column("logo_file_id").Type("int unsigned").Nullable()
	table.MustExec()
}

func (m *UpdateClassesAddLogoFileId) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("logo_file_id")
	table.MustExec()
}
