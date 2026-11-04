package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterClassesTable struct{}

func (m *CreateCharacterClassesTable) GetName() string {
	return "CreateCharacterClassesTable"
}

func (m *CreateCharacterClassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_classes", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("int unsigned").NotNull()
	table.Column("class_id").Type("int unsigned").NotNull()

	table.ForeignKey("character_id").Reference("characters").On("id").OnDelete("CASCADE").OnUpdate("CASCADE")
	table.ForeignKey("class_id").Reference("classes").On("id").OnDelete("CASCADE").OnUpdate("CASCADE")

	table.MustExec()
}

func (m *CreateCharacterClassesTable) Down(con *sqlx.DB) {
	builder.DropTable("character_classes", con).MustExec()
}
