package configuration

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"runtime"
)

const ERROR_DATABASE_ERROR = "DATABASE_ERROR"

var (
	ErrorMsg map[string]ErrorModel
)

type ErrorModel struct {
	InternalMessage string
	Httpcode        int
	StatusCode      int
	Message         string
}

func LoadError() {
	var _, filename, _, _ = runtime.Caller(0)
	configurationFilePath := path.Join(path.Dir(filename), "./json/error.message.json")
	absPath := configurationFilePath
	errFile, err := os.Open(absPath)
	defer func() {
		err = errFile.Close()
		if err != nil {
			log.Panic(err)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	jsonParser := json.NewDecoder(errFile)
	err = jsonParser.Decode(&ErrorMsg)
	if err != nil {
		log.Panic(err)
	}
}
func GetError(errMessage string) ErrorModel {
	var error ErrorModel
	error = ErrorMsg[errMessage]
	if (ErrorModel{}) == error {
		log.Panic("error info not found on configuration")
	}
	return error
}
