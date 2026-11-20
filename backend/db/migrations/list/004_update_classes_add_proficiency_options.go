package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassesAddProficiencyOptions struct{}

func (m *UpdateClassesAddProficiencyOptions) GetName() string {
	return "UpdateClassesAddProficiencyOptions"
}

func (m *UpdateClassesAddProficiencyOptions) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.Column("skill_proficiency_max_options").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *UpdateClassesAddProficiencyOptions) Down(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.DropColumn("skill_proficiency_max_options")
	table.MustExec()
}
