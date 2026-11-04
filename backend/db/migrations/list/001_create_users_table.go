package list

import "github.com/jmoiron/sqlx"

type CreateUsersTable struct{}

func (m *CreateUsersTable) GetName() string {
	return "CreateUsersTable"
}

func (m *CreateUsersTable) Up(con *sqlx.DB) {
	con.MustExec(`CREATE TABLE IF NOT EXISTS users (
		id         INT UNSIGNED AUTO_INCREMENT,
		username   VARCHAR(50)  NOT NULL,
		email      VARCHAR(255) NOT NULL,
		password   VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		CONSTRAINT users_pk PRIMARY KEY (id),
		CONSTRAINT users_username_unique UNIQUE (username),
		CONSTRAINT users_email_unique    UNIQUE (email)
	)`)
}

func (m *CreateUsersTable) Down(con *sqlx.DB) {
	con.MustExec(`DROP TABLE IF EXISTS users`)
}
