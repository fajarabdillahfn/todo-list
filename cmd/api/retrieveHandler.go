package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fajarabdillahfn/todo-list/models"
)

func (app *Application) getTasks(w http.ResponseWriter, r *http.Request) {
	var param models.TaskList
	var err error

	err = app.archiveDoneTask("")
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_process_error, "error while achieving done task(s)", nil)
		return
	}

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
		writeResponse(w, status_sql_error, err.Error(), nil)
		return
	}

	// check if no tasks with current filter
	if tasks == nil {
		if !isValidStatus(param.Status) {
			writeResponse(w, status_invalid_parameter, "invalid status", nil)
			return
		}

		err = writeResponse(w, status_no_data, "no result", tasks)
		if err != nil {
			app.logger.Println(err)
			writeResponse(w, status_process_error, err.Error(), nil)
			return
		}
		return
	}

	err = writeResponse(w, status_ok, "", tasks)
	if err != nil {
		app.logger.Println(err)
		writeResponse(w, status_process_error, err.Error(), nil)
		return
	}
}
