package app

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Server struct {
	*gin.Engine
}

func NewServer(session *DbSession) *Server {
	router := gin.New()
	api := router.Group("/")
	api.Use(gin.Recovery())
	api.Use(session.Database())
	{
		api.GET("/check", check)
		api.GET("/ping", pong)
		api.POST("/upload", uploadImageS3)
	}

	return &Server{router}
}

func check(ctx *gin.Context) {
	ctx.String(200, "OK")
}

func pong(ctx *gin.Context) {
	db := getDatabase(ctx)

	pings := db.C("pings")
	pong := Ping{}
	err := pings.Find(bson.M{}).One(&pong)
	if err != nil {
		ctx.String(400, "Error!")
		ctx.Abort()
		return
	}

	ctx.JSON(200, pong)
}

func uploadImageS3(ctx *gin.Context) {
	photo := Photo{}
	ctx.BindJSON(&photo)

	photo.uploadImageS3()
}

func getDatabase(ctx *gin.Context) *mgo.Database {
	return ctx.MustGet("db").(*mgo.Database)
}
