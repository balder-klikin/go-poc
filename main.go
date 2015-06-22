package main

import "github.com/balder-klikin/go-poc/app"

func main() {
	DbSession := app.NewDbSession("go-poc")
	server := app.NewServer(DbSession)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
