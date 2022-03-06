package models

import (
	"context"
	"time"
)

func (m *DBModel) DeleteTask(task string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `DELETE FROM todo_list
			 WHERE task = $1`

	_, err := m.DB.ExecContext(ctx, stmt, task)
	if err != nil {
		return err
	}

	return nil
}
