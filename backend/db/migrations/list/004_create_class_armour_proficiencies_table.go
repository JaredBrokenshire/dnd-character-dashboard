package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassArmourProficienciesTable struct{}

func (m *CreateClassArmourProficienciesTable) GetName() string {
	return "CreateClassArmourProficienciesTable"
}

func (m *CreateClassArmourProficienciesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("class_armour_proficiencies", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.Column("class_id").Type("INT UNSIGNED").NotNull()
	table.Column("armour").Type("ENUM('Light','Medium','Heavy','Shields')").NotNull()
	table.MustExec()
}

func (m *CreateClassArmourProficienciesTable) Down(con *sqlx.DB) {
	builder.DropTable("class_armour_proficiencies", con).MustExec()
}
