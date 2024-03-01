package entity

import "net/http"

type HTTPException struct {
	Message string
	Code    int
}

func (e *HTTPException) Error() string {
	return e.Message
}

func NewHTTPException(code int, message string) *HTTPException {
	if message == "" {
		message = getDefaultMessage(code)
	}

	return &HTTPException{
		Message: message,
		Code:    code,
	}
}

func getDefaultMessage(code int) string {
	switch code {
	case http.StatusBadRequest:
		return "Bad Request"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusUnprocessableEntity:
		return "Unprocessable Entity"
	case http.StatusConflict:
		return "Conflict"
	case http.StatusOK:
		return "OK"
	default:
		return "Unknown Error"
	}
}
