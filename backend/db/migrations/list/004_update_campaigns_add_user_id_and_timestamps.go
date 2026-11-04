package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCampaignsAddUserId struct{}

func (m *UpdateCampaignsAddUserId) GetName() string {
	return "UpdateCampaignsAddUserId"
}

func (m *UpdateCampaignsAddUserId) Up(con *sqlx.DB) {
	table := builder.ChangeTable("campaigns", con)
	table.Column("user_id").Type("int unsigned").NotNull()
	table.MustExec()
}

func (m *UpdateCampaignsAddUserId) Down(con *sqlx.DB) {
	table := builder.ChangeTable("campaigns", con)
	table.DropColumn("user_id")
	table.DropColumn("created_at")
	table.DropColumn("updated_at")
	table.MustExec()
}
