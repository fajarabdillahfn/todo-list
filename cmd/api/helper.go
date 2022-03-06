package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/fajarabdillahfn/todo-list/models"
)

const (
	task = iota
	status
)

func IsValidStatus(status string) bool {
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

func ValidateParam(identifier []int, param models.TaskList) (err error) {
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

func (app *Application) IsValidTask(task string) (bool, error) {
	check, err := app.models.DB.GetTasks(models.TaskList{Task: task})
	if err != nil {
		return false, err
	}

	if check == nil {
		return false, nil
	}

	return true, nil
}
