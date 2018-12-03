package main

var s *server

func init() {
	s = newServer()
	s.serverInit()
}
func main() {
	s.serverRun()
}
