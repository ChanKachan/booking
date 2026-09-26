package http

type ErrorResponse struct {
	Error ApiError `json:"error"`
}

type ApiError struct {
	Message string            `json:"message"`
	Code    string            `json:"code"`
	Fields  map[string]string `json:"fields,omitempty"`
}

type Response struct {
	Code string `json:"code"`
	Data any    `json:"data"`
}
