package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharactersTable struct{}

func (m *CreateCharactersTable) GetName() string {
	return "CreateCharactersTable"
}

func (m *CreateCharactersTable) Up(con *sqlx.DB) {
	table := builder.NewTable("characters", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("user_id").Type("int unsigned").NotNull()
	table.Column("campaign_id").Type("int unsigned").Nullable()
	table.String("name", 255).Nullable()
	table.Column("race_id").Type("int unsigned").Nullable()
	table.Integer("level").Default("1")
	table.MustExec()
}

func (m *CreateCharactersTable) Down(con *sqlx.DB) {
	builder.DropTable("characters", con).MustExec()
}
