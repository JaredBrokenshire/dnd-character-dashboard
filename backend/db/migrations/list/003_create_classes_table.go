package list

import (
	"dnd/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassesTable struct{}

func (m *CreateClassesTable) GetName() string {
	return "CreateClassesTable"
}

func (m *CreateClassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("classes", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement().Primary()
	table.String("name", 255).NotNull()
	table.Column("logo_file_id").Type("INT UNSIGNED").Nullable()
	table.Column("card_background_file_id").Type("INT UNSIGNED").Nullable()
	table.String("tagline", 255).Nullable()
	table.Column("card_description").Type("TEXT").Nullable()
	table.String("primary_ability", 255).NotNull()
	table.Integer("hit_point_die_value").NotNull()
	table.Column("save").Type("ENUM('Strength','Dexterity','Constitution','Intelligence','Wisdom','Charisma')").NotNull()
	table.Column("secondary_save").Type("ENUM('Strength','Dexterity','Constitution','Intelligence','Wisdom','Charisma')").Nullable()
	table.MustExec()
}

func (m *CreateClassesTable) Down(con *sqlx.DB) {
	builder.DropTable("classes", con).MustExec()
}
