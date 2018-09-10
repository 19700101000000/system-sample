package main

import (
	"./handler"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler.Root)
}
func main() {
	http.ListenAndServe(":8080", nil)
}
