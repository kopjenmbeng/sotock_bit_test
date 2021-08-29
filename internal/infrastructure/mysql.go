package infrastructure

import (
	_ "fmt"
	"strings"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	log "github.com/sirupsen/logrus"
)

type IMySql interface {
	Write() *sqlx.DB
	Read() *sqlx.DB
}

type MySql struct {
	log   *log.Logger
	read  string
	write string
}

func NewMySql(read string, write string, log *log.Logger) IMySql {
	return &MySql{read: read, write: write, log: log}
}

func (pg *MySql) Write() *sqlx.DB {
	db := sqlx.MustOpen("mysql", pg.write)
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	// ensure the database are reachable
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	pg.log.Println("write db connected successfully !")
	return db
}

func (pg *MySql) Read() *sqlx.DB {
	db := sqlx.MustOpen("mysql", pg.read)
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	// ensure the database are reachable
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	pg.log.Println("read db connected successfully !")
	return db
}
