package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fajarabdillahfn/todo-list/models"
)

func (app *Application) updateTask(w http.ResponseWriter, r *http.Request) {
	updTask := models.TaskList{
		Task:   strings.ToLower(r.FormValue("task")),
		Status: strings.ToLower(r.FormValue("status")),
	}

	err := validateParam([]int{task, status}, updTask)
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_invalid_parameter, err.Error(), nil)
		return
	}

	if !isValidStatus(updTask.Status) {
		writeResponse(w, status_invalid_parameter, "invalid status", nil)
		return
	}

	check, err := app.isValidTask(updTask.Task)
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_no_data, fmt.Sprintf("failed to validate task named '%s'", updTask.Task), []*models.TaskList{&updTask})
		return
	}
	if !check {
		app.logger.Println(err)
		writeResponse(w, status_no_data, fmt.Sprintf("update task failed, no task named '%s'", updTask.Task), nil)
		return
	}

	err = app.models.DB.UpdateTask(updTask.Task, updTask.Status)
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_sql_error, err.Error(), nil)

		return
	}

	err = writeResponse(w, status_ok, "update success", []*models.TaskList{&updTask})
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_process_error, err.Error(), nil)
		return
	}
}
