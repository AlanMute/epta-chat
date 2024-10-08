package resp

type ErrorResponse struct {
	Error string `json:"error"`
}

func Error(message string) ErrorResponse {
	return ErrorResponse{message}
}
