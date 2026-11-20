package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassesAddSkillProficiencyMaxOptions struct{}

func (m *UpdateClassesAddSkillProficiencyMaxOptions) GetName() string {
	return "UpdateClassesAddSkillProficiencyMaxOptions"
}

func (m *UpdateClassesAddSkillProficiencyMaxOptions) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.Column("skill_proficiency_max_options").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *UpdateClassesAddSkillProficiencyMaxOptions) Down(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.DropColumn("skill_proficiency_max_options")
	table.MustExec()
}
