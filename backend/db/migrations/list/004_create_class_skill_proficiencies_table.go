package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassSkillProficienciesTable struct{}

func (m *CreateClassSkillProficienciesTable) GetName() string {
	return "CreateClassSkillProficienciesTable"
}

func (m *CreateClassSkillProficienciesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("class_skill_proficiencies", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.Column("class_id").Type("INT UNSIGNED").NotNull()
	table.Column("skill").Type("ENUM('Acrobatics','Animal Handling','Arcana','Athletics','Deception','History','Insight','Intimidation','Investigation','Medicine','Nature','Perception','Performance','Persuasion','Religion','Sleight of Hand','Stealth','Survival')").NotNull()
	table.MustExec()
}

func (m *CreateClassSkillProficienciesTable) Down(con *sqlx.DB) {}
