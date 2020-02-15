package server

// setControllers ...
func (s *Server) setControllers()  {

	s.router.HandleFunc("/", s.Controllers.Index.HandleIndex()).Methods("GET")

}


