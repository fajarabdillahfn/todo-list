package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/fajarabdillahfn/todo-list/models"
)

type header struct {
	ResponseCode int    `json:"response_code"`
	Message      string `json:"message"`
}

type Response struct {
	Header header            `json:"header"`
	Data   []*models.TaskList `json:"data"`
}

func WriteResponse(w http.ResponseWriter, code int, message string, data []*models.TaskList) error {
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
