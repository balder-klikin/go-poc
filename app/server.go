package app

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type Server struct {
	*gin.Engine
}

func NewServer(session *DbSession) *Server {
	api := gin.Default()
	api.Use(SetDatabase(session))

	api.GET("/check", check)
	api.GET("/ping", pong)
	api.POST("/upload", uploadImageS3)

	return &Server{api}
}

func check(ctx *gin.Context) {
	ctx.String(200, "OK")
}

func pong(ctx *gin.Context) {
	db := GetDatabase(ctx)

	pings := db.C("pings")
	pong := Ping{}
	err := pings.Find(bson.M{}).One(&pong)
	if err != nil {
		ctx.String(400, "Error!")
		return
	}

	ctx.JSON(200, pong)
}

func uploadImageS3(ctx *gin.Context) {
	photo := Photo{}
	ctx.BindJSON(&photo)

	photo.uploadImageS3()
}
