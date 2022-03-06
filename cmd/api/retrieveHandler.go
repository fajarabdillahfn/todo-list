package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fajarabdillahfn/todo-list/models"
	util "github.com/fajarabdillahfn/todo-list/pkg"
)

func (app *Application) getTasks(w http.ResponseWriter, r *http.Request) {
	var param models.TaskList
	var err error

	if r.Body != nil {
		err = json.NewDecoder(r.Body).Decode(&param)
	}

	if r.Body == nil || err != nil {
		param = models.TaskList{
			Task:   strings.ToLower(r.FormValue("task")),
			Status: strings.ToLower(r.FormValue("status")),
		}
	}

	// get task(s) from db
	tasks, err := app.models.DB.GetTasks(param)
	if err != nil {
		app.logger.Println(err)
		util.WriteResponse(w, status_sql_error, err.Error(), nil)
		return
	}

	// check if no tasks with current filter
	if tasks == nil {
		if !IsValidStatus(param.Status) {
			util.WriteResponse(w, status_invalid_parameter, "invalid status", nil)
			return
		}

		err = util.WriteResponse(w, status_no_data, "no result", tasks)
		if err != nil {
			app.logger.Println(err)
			util.WriteResponse(w, status_process_error, err.Error(), nil)
			return
		}
		return
	}

	err = util.WriteResponse(w, status_ok, "", tasks)
	if err != nil {
		app.logger.Println(err)
		util.WriteResponse(w, status_process_error, err.Error(), nil)
		return
	}
}
