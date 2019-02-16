package db

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

// Mongo is the concrete type that holds the Mongo DB session for later usage
type Mongo struct {
	db *mgo.Session
}

var session *mgo.Session

// createMgoSession is a singleton function that creates Mongo object from configuration data
func createMgoSession() error {
	addr := viper.GetString("database.host") + ":" + viper.GetString("database.port")

	var err error
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{addr},
		Timeout:  5 * time.Second,
		Database: viper.GetString("database.name"),
		Username: viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
	}
	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		return err
	}
	log.Println("mongo: connected")
	return nil
	//session.SetMode(mgo.Monotonic, true)
	//session.SetPoolLimit(10000)
}

// NewMongo is a function adheres to the singleton pattern to make sure that only one instance of the database
// is used in order to minimize resource wastage and data duplication
func NewMongo() (*Mongo, error) {
	// checks if mongo session is already initiated or not
	log.Println(chalk.Blue, "mongo: connecting...")
	if session != nil {
		log.Println(chalk.Green, "mongo: connected")
		return &Mongo{session}, nil
	} else {
		if err := createMgoSession(); err != nil {
			return nil, err
		}
		return &Mongo{session}, nil
	}

}

// GetSession clones  the mongo session and returns it
func (m *Mongo) GetSession() *mgo.Session {
	return m.db.Clone()
}

// Dummy implementation for Db interface
func (m *Mongo) InsertData(data interface{}) {

}

// Dummy implementation for Db interface
func (m *Mongo) ExecRaw(query string, args ...interface{}) *gorm.DB {
	return nil
}

// Dummy implementation for Db interface
func (m *Mongo) InsertRaw(query string, args ...interface{}) {

}

// Dummy implementation for Db interface
func (m *Mongo) Migrate(models ...interface{}) {

}

// Dummy implementation for Db interface
func (m *Mongo) Update(model interface{}) {

}

// Dummy implementation for Db interface
func (m *Mongo) CheckExists(field string, data interface{}, table string) bool {

	return false
}

// Dummy implementation for Db interface
func (m *Mongo) Delete(field string, data interface{}, table string) {

}

// Dummy implementation for Db interface
func (m *Mongo) Close() {

}
