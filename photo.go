package main

type Photo struct {
	Name        string `json:"name"`
	DataBase64  string `json:"data"`
	ContentType string `json:"contentType"`
}
