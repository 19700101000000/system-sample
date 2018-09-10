package main

import (
	"github.com/19700101000000/system-sample/server/handler"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler.Root)
}
func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
