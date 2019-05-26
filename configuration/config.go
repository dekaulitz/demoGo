package configuration

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
)

type Configuration struct {
	Database MySQL
	Host     Host
}

var (
	//ConfigurationData contains all configuration data
	ConfigurationData = &Configuration{}
)

func init() {
	goEnvironment, _ := os.LookupEnv("GODEMOENV")
	if goEnvironment == "" {
		goEnvironment = "test"
		os.Setenv("CONF_PATH", "$GOPATH/src/demoGo")

	}
	log.Printf("load configuration from .env.%s", goEnvironment)
	err := godotenv.Load(os.ExpandEnv(path.Join(os.Getenv("CONF_PATH"), ".env."+goEnvironment)))
	if err != nil {
		log.Panic("ErrorMapper loading .env file", err)
	}

	ConfigurationData.Database = MySQL{os.Getenv("HOSTNAME"), os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"), os.Getenv("MAX_OPEN_CONNECTIONS"),
		os.Getenv("MAX_IDLE_CONNECTIONS"), os.Getenv("SCHEMA")}
	ConfigurationData.Host = Host{os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")}

	//log.SetOutput(&lumberjack.Logger{
	//	Filename:   "logs/application.log",
	//	MaxSize:    500, // megabytes
	//	MaxBackups: 3,
	//	MaxAge:     28,   // days
	//	Compress:   true, // disabled by default
	//})
	//govalidator.SetFieldsRequiredByDefault(true)
	LoadError()
}

//GetConfiguration Get the new configuration
func GetConfiguration() *Configuration {
	return ConfigurationData
}
