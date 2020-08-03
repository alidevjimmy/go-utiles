package errors

import "net/http"

// RestErr is the helper functino for return better and easeier error
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequestErr message for bad Requests
func NewBadRequestErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad_Request",
	}
}
