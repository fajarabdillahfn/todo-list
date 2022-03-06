package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fajarabdillahfn/todo-list/models"
)

func (app *Application) deleteTask(w http.ResponseWriter, r *http.Request) {
	delTask := models.TaskList{
		Task: strings.ToLower(r.FormValue("task")),
	}

	err := validateParam([]int{task}, delTask)
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_invalid_parameter, err.Error(), nil)
		return
	}

	check, err := app.isValidTask(delTask.Task)
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_no_data, fmt.Sprintf("failed to validate task named '%s'", delTask.Task), []*models.TaskList{&delTask})
		return
	}
	if !check {
		app.logger.Println(err)
		writeResponse(w, status_no_data, fmt.Sprintf("delete task failed, no task named '%s'", delTask.Task), nil)
		return
	}

	err = app.models.DB.DeleteTask(delTask.Task)
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_sql_error, err.Error(), nil)

		return
	}

	err = writeResponse(w, status_ok, "delete success", []*models.TaskList{&delTask})
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_process_error, err.Error(), nil)
		return
	}
}
