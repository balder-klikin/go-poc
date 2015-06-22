package app

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"

	"github.com/gin-gonic/gin"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/nfnt/resize"
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
	auth, _ := aws.SharedAuth() // reads auth from ~/.aws/credentials
	region := aws.EUWest
	s3client := s3.New(auth, region)
	bucket := s3client.Bucket("media-poc-baldercm")

	photo := Photo{}
	ctx.BindJSON(&photo)

	data, err := base64.StdEncoding.DecodeString(photo.DataBase64)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}

	imageReader := bytes.NewReader(data)
	image, err := jpeg.Decode(imageReader)
	if err != nil {
		fmt.Println("jpeg error:", err)
		panic(err)
	}
	small := resize.Thumbnail(50, 50, image, resize.Bicubic)
	medium := resize.Thumbnail(100, 100, image, resize.Bicubic)

	smallBytes := new(bytes.Buffer)
	jpeg.Encode(smallBytes, small, nil)

	mediumBytes := new(bytes.Buffer)
	jpeg.Encode(mediumBytes, medium, nil)

	bucket.Put(photo.Name, data, photo.ContentType, "public-read", s3.Options{})
	bucket.Put(photo.Name+"_small", smallBytes.Bytes(), photo.ContentType, "public-read", s3.Options{})
	bucket.Put(photo.Name+"_medium", mediumBytes.Bytes(), photo.ContentType, "public-read", s3.Options{})
}
