package configuration

/**
this go file will be load the file response_message.json into [] object
this aim for easy response handling messages either error nor success response
*/
import (
	"encoding/json"
	"log"
	"os"
	"path"
	"runtime"
)

var (
	MapMessage map[string]MessageModel
)

type MessageModel struct {
	InternalMessage string
	Httpcode        int
	StatusCode      int
	Message         string
}

/**
load message configuration and binding into [] MessageModel struct
*/
func LoadMessage() {
	var _, filename, _, _ = runtime.Caller(0)
	configurationFilePath := path.Join(path.Dir(filename), "./json/response_message.json")
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
	err = jsonParser.Decode(&MapMessage)
	if err != nil {
		log.Panic(err)
	}
}

//get message base on map with key and return message struct
func GetMessage(errMessage string) MessageModel {
	var message MessageModel
	message = MapMessage[errMessage]
	if (MessageModel{}) == message {
		log.Panic("message info not found on configuration")
	}
	return message
}
