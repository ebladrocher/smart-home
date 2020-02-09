package server

func (s *Server) setControllers()  {

	s.router.HandleFunc("/", s.Controllers.Index.HandleIndex())

}


