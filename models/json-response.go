package models


type JSONResponse struct {
	ErrorCode int			`json:"error_code,omitempty"`
	Message string			`json:"message,omitempty"`
	Data    interface{}     `json:"data,omitempty"`
}