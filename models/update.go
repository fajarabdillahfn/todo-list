package models

import (
	"context"
	"time"
)

func (m *DBModel) UpdateTask(task, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE todo_list
			 SET status = $1
			 WHERE task = $2`

	_, err := m.DB.ExecContext(ctx, stmt, status, task)
	if err != nil {
		return err
	}

	return nil
}
