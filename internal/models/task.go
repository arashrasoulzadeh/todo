package models

import "database/sql"

type Task struct {
	Id     int `json:"id,nullable"`
	Title  string
	Desc   sql.NullString
	Status uint8 // pending , progress, done
}
