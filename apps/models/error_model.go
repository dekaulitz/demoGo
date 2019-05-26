package models

type ErrorModel struct {
	Error   error
	Message string
}

func ThrowError() ErrorModel {
	var err ErrorModel
	return err
}
