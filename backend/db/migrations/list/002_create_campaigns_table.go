package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCampaignsTable struct{}

func (m *CreateCampaignsTable) GetName() string {
	return "CreateCampaignsTable"
}

func (m *CreateCampaignsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("campaigns", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateCampaignsTable) Down(con *sqlx.DB) {
	builder.DropTable("campaigns", con).MustExec()
}
