package http

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

type HttpResponseBodyPagination struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

func BuildHttpResponseBodyPagination(success bool, message string, data interface{}, meta interface{}) HttpResponseBodyPagination {
	return HttpResponseBodyPagination{
		Success: success,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}
