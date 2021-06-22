package httputil

import (
	"strings"
)

// JSONResponse
type JSONResponse struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	Error     string      `json:"error"`
	Validator interface{} `json:"validate,omitempty"`
}

// JSONResponse
func NewOkJSON(data interface{}, message ...string) *JSONResponse {
	msg := "ok"
	if len(message) > 0 {
		msg = strings.Join(message, " ")
	}
	return &JSONResponse{Code: 0, Data: data, Message: msg}
}

// JSONResponse
func NewFailJSON(bizCode int, message string) *JSONResponse {
	return &JSONResponse{Code: bizCode, Message: message}
}

// JSONResponse
func NewValidatorFailJSON(bizCode int, validate interface{}, message ...string) *JSONResponse {
	msg := "参数校验不通过"
	if len(message) > 0 {
		msg = strings.Join(message, " ")
	}
	return &JSONResponse{Code: bizCode, Message: msg, Validator: validate}
}

// JSONResponse
func NewErrorJSON(bizCode int, error string, message ...string) *JSONResponse {
	return &JSONResponse{Code: bizCode, Error: error, Message: strings.Join(message, " ")}
}
