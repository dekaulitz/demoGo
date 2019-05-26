package configuration

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

//MySQL struct
type MySQL struct {
	Hostname     string
	Username     string
	Password     string
	MaxOpenConns string
	MaxIdleConns string
	Schema       string
}

var (
	engine                  = &xorm.Engine{}
	conf                    = GetConfiguration()
	ErrFailedToConnectToSQL = "Failed to connect to mysql %v\n"
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		conf.Database.Username, conf.Database.Password, conf.Database.Hostname, conf.Database.Schema)
	var err error
	fmt.Println(dsn)
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		err = fmt.Errorf(ErrFailedToConnectToSQL, err)
		panic(err.Error())
	}
	maxCon, err := strconv.Atoi(conf.Database.MaxOpenConns)
	if err != nil {
		panic(err)
	}
	maxIdleCon, err := strconv.Atoi(conf.Database.MaxIdleConns)
	engine.SetMaxOpenConns(maxCon)
	engine.SetMaxIdleConns(maxIdleCon)
	engine.SetConnMaxLifetime(-1)
}

func GetConnection() *xorm.Session {
	return engine.NewSession()
}

func GetEngine() *xorm.Engine {
	return engine
}
