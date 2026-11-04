package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterSubclassesTable struct{}

func (m *CreateCharacterSubclassesTable) GetName() string {
	return "CreateCharacterSubclassesTable"
}

func (m *CreateCharacterSubclassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_subclasses", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("int unsigned").NotNull()
	table.Column("subclass_id").Type("int unsigned").NotNull()

	table.ForeignKey("character_id").Reference("characters").On("id").OnDelete("CASCADE").OnUpdate("CASCADE")
	table.ForeignKey("subclass_id").Reference("subclasses").On("id").OnDelete("CASCADE").OnUpdate("CASCADE")

	table.MustExec()
}

func (m *CreateCharacterSubclassesTable) Down(con *sqlx.DB) {
	builder.DropTable("character_subclasses", con).MustExec()
}
