package main

import "github.com/balder-klikin/go-poc/gopoc"

func main() {
	server := gopoc.NewServer()

	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
