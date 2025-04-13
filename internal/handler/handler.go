package handler

import (
	"real-time-ranking/internal/service"
)

type Server struct {
	s service.Service
}

func NewServer(s service.Service) *Server {
	return &Server{
		s: s,
	}
}
