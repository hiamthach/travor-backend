package model

type ErrorResponse struct {
	Error interface{} `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
