package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fajarabdillahfn/todo-list/models"
	util "github.com/fajarabdillahfn/todo-list/pkg"
)

func (app *Application) insertTask(w http.ResponseWriter, r *http.Request) {
	newTask := models.TaskList{
		Task:   strings.ToLower(r.FormValue("task")),
		Status: strings.ToLower(r.FormValue("status")),
	}

	err := ValidateParam([]int{task}, newTask)
	if err != nil {
		app.logger.Println(err)
		util.WriteResponse(w, status_invalid_parameter, err.Error(), nil)
		return
	}

	if newTask.Status == "" {
		newTask.Status = "backlog"
	}

	err = app.models.DB.InsertTask(newTask)
	if err != nil {
		app.logger.Println(err)
		if strings.Contains(err.Error(), "duplicate") {
			util.WriteResponse(w, status_data_exist, fmt.Sprintf("task '%s' is already exist and not done yet", newTask.Task), nil)
		} else {
			util.WriteResponse(w, status_sql_error, err.Error(), nil)
		}
		return
	}

	err = util.WriteResponse(w, status_ok, "insert success", []*models.TaskList{&newTask})
	if err != nil {
		app.logger.Println(err)
		util.WriteResponse(w, status_process_error, err.Error(), nil)
		return
	}
}
