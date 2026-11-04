package list

import "github.com/jmoiron/sqlx"

type CreateRolesTable struct{}

func (m *CreateRolesTable) GetName() string {
	return "CreateRolesTable"
}

func (m *CreateRolesTable) Up(con *sqlx.DB) {
	con.MustExec(`CREATE TABLE IF NOT EXISTS roles (
		id   INT UNSIGNED AUTO_INCREMENT,
		name VARCHAR(50) NOT NULL,
		CONSTRAINT roles_pk PRIMARY KEY (id),
		CONSTRAINT roles_name_unique UNIQUE (name)
	)`)
}

func (m *CreateRolesTable) Down(con *sqlx.DB) {
	con.MustExec(`DROP TABLE IF EXISTS roles`)
}
