package list

import (
	"dnd/pkg/go-migrations/builder"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CreateClassWeaponProficienciesTable struct{}

func (m *CreateClassWeaponProficienciesTable) GetName() string {
	return "CreateClassWeaponProficienciesTable"
}

func (m *CreateClassWeaponProficienciesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("class_weapon_proficiencies", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.Column("class_id").Type("INT UNSIGNED").NotNull()
	table.Column("weapon_type").Type("ENUM('Simple','Martial')").NotNull()
	table.String("`condition`", 255).Nullable()
	fmt.Println(">>> ", table.GetSQL())
	table.MustExec()
}

func (m *CreateClassWeaponProficienciesTable) Down(con *sqlx.DB) {
	builder.DropTable("class_weapon_proficiencies", con).MustExec()
}
