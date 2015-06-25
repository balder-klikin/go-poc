package app

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type DbSession struct {
	*mgo.Session
	db string
}

func NewDbSession(name string) *DbSession {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	createIndexes(session.DB(name))
	return &DbSession{session, name}
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

func (DbSession *DbSession) Database() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s := DbSession.Copy()
		ctx.Set("db", s.DB(DbSession.db))
		defer s.Close()

		ctx.Next()
	}
}
