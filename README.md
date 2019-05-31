## Technology Stack

Following is technology stack that is used in this microservice

| Name | Description |
|------|-------------|
| [Golang](https://golang.org/) | Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. |
| [Glide](https://glide.sh/) | Package Management for Go. |
| [Mysql Connection](http://xorm.io/) | Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle |
| [Mysql Driver](https://github.com/go-sql-driver/mysql) | Mysql driver for database connection|


### Backend Setup
First, clone this repo on golang src directory (`$GOPATH/src`) by using this command:

```
cd $GOPATH/src && git clone https://github.com/dekaulitz/demoGo
```


and then, **edit** `.env` file to be the expected configuration to `.env.test`. 
if you using the different path you can edit the configuration path on file config.go

use Glide to install this repo dependencies:

```bash
glide install
```
and then build the project

```bash
go build
```

Finally, we can start the server by running:
```bash
go run main.go
```
