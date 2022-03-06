package models

import "database/sql"

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{
			DB: db,
		},
	}
}

type DBModel struct {
	DB *sql.DB
}

type TaskList struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}