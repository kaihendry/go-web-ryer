package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleGreeting("Howdy %s"))
}
