package main

func main() {
	DbSession := NewDbSession("go-poc")
	server := NewServer(DbSession)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
