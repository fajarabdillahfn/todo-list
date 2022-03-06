package models

import (
	"context"
	"time"
)

func (m *DBModel) InsertTask(param TaskList) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO todo_list (task, status)
			 VALUES ($1, $2)`

	_, err := m.DB.ExecContext(ctx, stmt, param.Task, param.Status)
	if err != nil {
		return err
	}

	return nil
}
