package model

type GenericRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Meta    interface{} `json:"meta"`
}
