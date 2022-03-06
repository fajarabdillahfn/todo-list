package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/tasks", app.getTasks)
	router.HandlerFunc(http.MethodPost, "/v1/add_task", app.insertTask)
	router.HandlerFunc(http.MethodDelete, "/v1/del_task", app.deleteTask)

	return router
}
