package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassSaveProficienciesTable struct{}

func (m *CreateClassSaveProficienciesTable) GetName() string {
	return "CreateClassSaveProficienciesTable"
}

func (m *CreateClassSaveProficienciesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("class_save_proficiencies", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.Column("class_id").Type("INT UNSIGNED").NotNull()
	table.Column("ability").Type("ENUM('Strength','Dexterity','Constitution','Intelligence','Wisdom','Charisma')").NotNull()
	table.MustExec()
}

func (m *CreateClassSaveProficienciesTable) Down(con *sqlx.DB) {
	builder.DropTable("class_save_proficiencies", con).MustExec()
}
