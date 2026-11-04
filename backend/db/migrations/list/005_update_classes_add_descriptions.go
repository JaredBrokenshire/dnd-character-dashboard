package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassesAddDescriptions struct{}

func (m *UpdateClassesAddDescriptions) GetName() string {
	return "UpdateClassesAddDescriptions"
}

func (m *UpdateClassesAddDescriptions) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.String("tagline", 255).Nullable()
	table.Column("card_description").Type("text").Nullable()
	table.String("primary_ability", 255).NotNull()
	table.Integer("hit_point_die_value").NotNull()
	table.Column("save").Type("ENUM('Strength','Constitution','Dexterity','Intelligence','Wisdom','Charisma')").NotNull()
	table.Column("secondary_save").Type("ENUM('Strength','Constitution','Dexterity','Intelligence','Wisdom','Charisma')").NotNull()
	table.MustExec()
}

func (m *UpdateClassesAddDescriptions) Down(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.DropColumn("tagline")
	table.DropColumn("card_description")
	table.DropColumn("primary_ability")
	table.DropColumn("hit_point_die_value")
	table.DropColumn("save")
	table.DropColumn("secondary_save")
	table.MustExec()
}
