package models

import (
	"context"
	"time"
)

func (m *DBModel) ArchiveTask(task string) error {
	getDoneTasks, err := m.GetTasks(TaskList{Task: task, Status: "done"})
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO archived_task (task, achived_date)
			 VALUES ($1, $2)`

	for _, doneTask := range getDoneTasks {
		err := m.DeleteTask(doneTask.Task)
		if err != nil {
			return err
		}

		_, err = m.DB.ExecContext(ctx, stmt, doneTask.Task, time.Now().Format("02-01-2006"))
		if err != nil {
			return err
		}
	}

	return nil
}
