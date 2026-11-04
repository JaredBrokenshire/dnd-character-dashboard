package list

import (
	"dnd-character-dashboard/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddProfilePictureFileId struct{}

func (m *UpdateCharactersAddProfilePictureFileId) GetName() string {
	return "UpdateCharactersAddProfilePictureFileId"
}

func (m *UpdateCharactersAddProfilePictureFileId) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("profile_picture_file_id").Type("int unsigned").Nullable()
	table.MustExec()
}

func (m *UpdateCharactersAddProfilePictureFileId) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("profile_picture_file_id")
	table.MustExec()
}
