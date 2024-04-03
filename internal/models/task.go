package models

import "database/sql"

type Task struct {
	Id     int
	Title  string
	Desc   sql.NullString
	Status uint8
}
