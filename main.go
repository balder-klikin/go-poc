package main

import "github.com/balder-klikin/go-poc/gopoc"

func main() {
	mgoSession := gopoc.NewMgoSession("go-poc")
	server := gopoc.NewServer(mgoSession)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
