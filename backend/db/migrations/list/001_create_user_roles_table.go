package list

import "github.com/jmoiron/sqlx"

type CreateUserRolesTable struct{}

func (m *CreateUserRolesTable) GetName() string {
	return "CreateUserRolesTable"
}

func (m *CreateUserRolesTable) Up(con *sqlx.DB) {
	con.MustExec(`CREATE TABLE IF NOT EXISTS user_roles (
		user_id INT UNSIGNED NOT NULL,
		role_id INT UNSIGNED NOT NULL,
		CONSTRAINT user_roles_pk          PRIMARY KEY (user_id, role_id),
		CONSTRAINT fk_user_roles_user     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		CONSTRAINT fk_user_roles_role     FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
	)`)
}

func (m *CreateUserRolesTable) Down(con *sqlx.DB) {
	con.MustExec(`DROP TABLE IF EXISTS user_roles`)
}
