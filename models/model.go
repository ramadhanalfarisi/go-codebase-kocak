package models

import "database/sql"

type Model struct {
	DB        *sql.DB
	Model     interface{}
	Parameter *string
}
