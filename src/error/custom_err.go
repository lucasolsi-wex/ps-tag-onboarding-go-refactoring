package error

import "net/http"

type CustomErr struct {
	Message string   `json:"message,omitempty"`
	Err     string   `json:"err,omitempty"`
	Code    int      `json:"code,omitempty"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *CustomErr) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *CustomErr {
	return &CustomErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *CustomErr {
	return &CustomErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewUserNotFoundError(message string) *CustomErr {
	return &CustomErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}
