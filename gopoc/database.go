package gopoc

import (
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo"
	"log"
)


type MgoSession struct {
	*mgo.Session
	db string
}


func NewMgoSession(name string) *MgoSession {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	createIndexes(session.DB(name))
	return &MgoSession{session, name}
}

func createIndexes(db *mgo.Database) {
	index := mgo.Index{
		Key:      []string{"value"},
		Unique:   false,
		DropDups: true,
	}
	indexErr := db.C("pings").EnsureIndex(index)
	if indexErr != nil {
		panic(indexErr)
	}
}

func SetDatabase(mgoSession *MgoSession) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := mgoSession.Copy()
		c.Set("db", s.DB(mgoSession.db))
		defer s.Close()

		c.Next()
	}
}

func GetDatabase(c *gin.Context) *mgo.Database {
	_db, exists := c.Get("db")
	if !exists {
		log.Fatal("Database !exists")
	}

	db, ok := _db.(*mgo.Database)
	if !ok {
		log.Fatal("Database !ok")
	}

	return db
}
