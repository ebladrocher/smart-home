package server

func (s *Server) setControllers() {

	r := s.engine

	basePath := r.Group("/api")

	v1 := basePath.Group("/v1")
	v1.GET("/", s.Controllers.Index.Index)
}
