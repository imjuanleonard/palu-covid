package v1

import (
	"encoding/json"
	"github.com/imjuanleonard/palu-covid/pkg/logger"
	"net/http"
)

const (
	CodeServerError    = "covid::server_error"
	CodeInvalidRequest = "covid::invalid_request"
)

type Error struct {
	Code         string `json:"code"`
	Message      string `json:"message"`
	MessageTitle string `json:"message_title"`
}

type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Errors  []Error     `json:"errors"`
}

func NewSuccessResponse(data interface{}) *Response {
	r := &Response{
		Data:    data,
		Success: true,
		Errors:  []Error{},
	}
	return r
}

func NewErrorResponse(code, message, title string) *Response {
	r := &Response{
		Data:    struct{}{},
		Success: false,
		Errors:  []Error{{code, message, title}},
	}
	return r
}

func (r *Response) Write(w http.ResponseWriter, status int) {
	body, err := json.Marshal(r)
	if err != nil {
		logger.Errorf("could not marshall response : %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write(body); err != nil {
		logger.Errorf("could not write to response: %v", err)
	}
}
