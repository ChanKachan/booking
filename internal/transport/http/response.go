package http

type ErrorResponse struct {
	Error ApiError `json:"error"`
}

type ApiError struct {
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Fields  map[string]string `json:"fields,omitempty"`
}

type Response struct {
	Code int `json:"code"`
	Data any `json:"data"`
}
