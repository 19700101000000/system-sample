package main

import (
	"./handler"
	"log"
	"net/http"
)

type server struct {
}

func (s *server) init() {
	http.HandleFunc("/", handler.Root)
}
func (s *server) run() {
	//http.ListenAndServe(":3939", nil)
	if err := http.ListenAndServeTLS(":3939", "ssl/server.crt", "ssl/server.key", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	s := server{}
	s.init()
	s.run()
}
