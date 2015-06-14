package gopoc

import (
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo/bson"
)

func NewServer(mgoSession *MgoSession) *gin.Engine {
	api := gin.Default()
	api.Use(SetDatabase(mgoSession))

	api.GET("/ping", pong)

	return api
}

func pong(c *gin.Context) {
	db := GetDatabase(c)

	pings := db.C("pings")
	pong := Ping{}
	err := pings.Find(bson.M{}).One(&pong)
	if err != nil {
		c.String(400, "Error!")
		return
	}

	c.JSON(200, pong)
}
