package models

type APIError struct {
	OK        bool        `json:"ok"`
	ErrorCode string      `json:"code"`
	ErrorInfo interface{} `json:"info"`
}

type APIResponse struct {
	OK   bool        `json:"ok"`
	Data interface{} `json:"data"`
}
