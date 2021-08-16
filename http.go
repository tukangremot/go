package main

type HttpResponseBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildHttpResponseBody(success bool, message string, data interface{}) HttpResponseBody {
	return HttpResponseBody{
		Success: success,
		Message: message,
		Data:    data,
	}
}
