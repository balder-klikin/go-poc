package main

import "github.com/balder-klikin/go-poc/gopoc"

func main() {
	DbSession := gopoc.NewDbSession("go-poc")
	server := gopoc.NewServer(DbSession)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
