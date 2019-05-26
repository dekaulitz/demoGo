package vmodel

import "time"

type Meta struct {
	Timestamp  time.Time `json:"timestamp"`
	HttpCode   int       `json:"http_code"`
	StatusCode int       `json:"status"`
	Message    string    `json:"message"`
}

type ResponseModel struct {
	Meta Meta        `json:"meta"`
	Body interface{} `json:"body"`
}

func GetMeta(meta Meta) Meta {
	return meta
}
