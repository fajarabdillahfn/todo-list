package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fajarabdillahfn/todo-list/models"
	util "github.com/fajarabdillahfn/todo-list/pkg"
)

func (app *Application) deleteTask(w http.ResponseWriter, r *http.Request) {
	delTask := models.TaskList{
		Task: strings.ToLower(r.FormValue("task")),
	}

	err := ValidateParam([]int{task}, delTask)
	if err != nil {
		app.logger.Println(err)
		util.WriteResponse(w, status_invalid_parameter, err.Error(), nil)
		return
	}

	check, err := app.IsValidTask(delTask.Task)
	if err != nil {
		app.logger.Println(err)
		util.WriteResponse(w, status_no_data, fmt.Sprintf("failed to validate task named '%s'", delTask.Task), []*models.TaskList{&delTask})
		return
	}
	if !check {
		app.logger.Println(err)
		util.WriteResponse(w, status_no_data, fmt.Sprintf("delete task failed, no task named '%s'", delTask.Task), nil)
		return
	}

	err = app.models.DB.DeleteTask(delTask.Task)
	if err != nil {
		app.logger.Println(err)
		util.WriteResponse(w, status_sql_error, err.Error(), nil)

		return
	}

	err = util.WriteResponse(w, status_ok, "delete success", []*models.TaskList{&delTask})
	if err != nil {
		app.logger.Println(err)
		util.WriteResponse(w, status_process_error, err.Error(), nil)
		return
	}
}
