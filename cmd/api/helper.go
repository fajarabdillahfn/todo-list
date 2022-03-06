package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/fajarabdillahfn/todo-list/models"
)

func isValidStatus(status string) bool {
	switch strings.ToLower(status) {
	case
		"",
		"done",
		"in progress",
		"todo",
		"backlog":
		return true
	}
	return false
}

const (
	task = iota
	status
)

func validateParam(identifier []int, param models.TaskList) (err error) {
	message := "Param '%s' is Required"

	var buffer bytes.Buffer

	for _, v := range identifier {
		switch v {
		case task:
			if param.Task == "" {
				buffer.WriteString(fmt.Sprintf(message, "task"))
			}
		case status:
			if param.Task == "" {
				buffer.WriteString(fmt.Sprintf(message, "status"))
			}
		}
	}

	errMsg := buffer.String()

	if errMsg != "" {
		err = errors.New(errMsg)
	}

	return
}

func (app *Application) isValidTask(task string) (bool, error) {
	check, err := app.models.DB.GetTasks(models.TaskList{Task: task})
	if err != nil {
		return false, err
	}

	if check == nil {
		return false, nil
	}

	return true, nil
}

type header struct {
	ResponseCode int    `json:"response_code"`
	Message      string `json:"message"`
}

type Response struct {
	Header header             `json:"header"`
	Data   []*models.TaskList `json:"data"`
}

func writeResponse(w http.ResponseWriter, code int, message string, data []*models.TaskList) error {
	responseData := data

	if data == nil {
		responseData = []*models.TaskList{
			{
				Task:   "",
				Status: "",
			},
		}
	}

	response := Response{
		Header: header{
			ResponseCode: code,
			Message:      message,
		},
		Data: responseData,
	}

	js, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

	return nil
}

func (app *Application) archiveDoneTask(task string) error {
	err := app.models.DB.ArchiveTask(task)
	if err != nil {
		return err
	}
	return nil
}
