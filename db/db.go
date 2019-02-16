package db

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

// Db is an interface that any database concrete type must implement in order to access database later.
// It has several important methods mostly for SQL database.
type Db interface {
	InsertData(data interface{})
	ExecRaw(query string, args ...interface{}) *gorm.DB
	InsertRaw(query string, args ...interface{})
	Migrate(models ...interface{})
	CheckExists(field string, data interface{}, table string) bool
	Update(model interface{})
	Delete(field string, data interface{}, table string)
	Close()
	GetSession() *mgo.Session
}