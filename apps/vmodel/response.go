package vmodel

import "time"

type Meta struct {
	RequestId  string    `json:"request_id"`
	Timestamp  time.Time `json:"timestamp"`
	HttpCode   int       `json:"http_code"`
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
}

type ResponseModel struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
