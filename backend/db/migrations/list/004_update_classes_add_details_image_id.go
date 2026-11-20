package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassesAddDetailsImageId struct{}

func (m *UpdateClassesAddDetailsImageId) GetName() string {
	return "UpdateClassesAddDetailsImageId"
}

func (m *UpdateClassesAddDetailsImageId) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.Column("details_image_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *UpdateClassesAddDetailsImageId) Down(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.DropColumn("details_image_id")
	table.MustExec()
}
