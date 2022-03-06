package models

import (
	"context"
	"fmt"
	"time"
)

func (m *DBModel) 	GetTasks(param TaskList) ([]*TaskList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	blank := TaskList{
		Task:   "",
		Status: "",
	}

	if param != blank {
		where = "WHERE "
		if param.Task != "" {
			where += fmt.Sprintf("task = '%s' ", param.Task)
			if param.Status != "" {
				where += fmt.Sprintf("AND status = '%s' ", param.Status)
			}
		} else if param.Status != "" {
			where += fmt.Sprintf("status = '%s' ", param.Status)
		}
	}

	stmt := fmt.Sprintf(`SELECT *
			 FROM todo_list
			 %s`, where)

	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*TaskList

	for rows.Next() {
		var task TaskList
		err := rows.Scan(
			&task.Task,
			&task.Status,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}
