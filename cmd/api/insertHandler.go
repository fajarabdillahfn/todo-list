package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fajarabdillahfn/todo-list/models"
)

func (app *Application) insertTask(w http.ResponseWriter, r *http.Request) {
	newTask := models.TaskList{
		Task:   strings.ToLower(r.FormValue("task")),
		Status: strings.ToLower(r.FormValue("status")),
	}

	err := validateParam([]int{task}, newTask)
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_invalid_parameter, err.Error(), nil)
		return
	}

	if newTask.Status == "" {
		newTask.Status = "backlog"
	}

	err = app.models.DB.InsertTask(newTask)
	if err != nil {
		app.logger.Println(err)
		if strings.Contains(err.Error(), "duplicate") {
			writeResponse(w, status_data_exist, fmt.Sprintf("task '%s' is already exist and not done yet", newTask.Task), nil)
		} else {
			writeResponse(w, status_sql_error, err.Error(), nil)
		}
		return
	}

	err = writeResponse(w, status_ok, "insert success", []*models.TaskList{&newTask})
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_process_error, err.Error(), nil)
		return
	}
}
