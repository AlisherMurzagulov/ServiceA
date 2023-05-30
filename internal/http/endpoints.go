package http

func (s *server) initEndpoints() {
	s.router.POST("/payment", s.handlers.TransferHandler)
	s.router.GET("/health")
}
