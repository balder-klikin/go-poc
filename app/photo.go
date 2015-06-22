package app

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/nfnt/resize"
)

type Photo struct {
	Name        string `json:"name"`
	DataBase64  string `json:"data"`
	ContentType string `json:"contentType"`
}

func (p *Photo) uploadImageS3() {
	data, err := base64.StdEncoding.DecodeString(p.DataBase64)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}

	imageReader := bytes.NewReader(data)
	image, format, err := image.Decode(imageReader)
	fmt.Println(format)
	if err != nil {
		fmt.Println("image.Decode error:", err)
		panic(err)
	}
	small := resize.Thumbnail(50, 50, image, resize.Bicubic)
	medium := resize.Thumbnail(100, 100, image, resize.Bicubic)

	smallBytes := new(bytes.Buffer)
	jpeg.Encode(smallBytes, small, nil)

	mediumBytes := new(bytes.Buffer)
	jpeg.Encode(mediumBytes, medium, nil)

	auth, _ := aws.SharedAuth() // reads auth from ~/.aws/credentials
	region := aws.EUWest
	s3client := s3.New(auth, region)
	bucket := s3client.Bucket("media-poc-baldercm")

	var uploadErr error
	uploadErr = bucket.Put(p.Name, data, p.ContentType, "public-read", s3.Options{})
	if uploadErr != nil {
		panic(uploadErr)
	}
	uploadErr = bucket.Put(p.Name+"_small", smallBytes.Bytes(), p.ContentType, "public-read", s3.Options{})
	if uploadErr != nil {
		panic(uploadErr)
	}
	uploadErr = bucket.Put(p.Name+"_medium", mediumBytes.Bytes(), p.ContentType, "public-read", s3.Options{})
	if uploadErr != nil {
		panic(uploadErr)
	}
}
