package models

type ErrorResponse struct {
	Message   string `json:"message"`
	Throwable string `json:"throwable"`
}

func BuildError(message, throwable string) *ErrorResponse {
	return &ErrorResponse{
		Message:   message,
		Throwable: throwable,
	}
}
