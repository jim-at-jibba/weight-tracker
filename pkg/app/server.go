package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"jamesbest.tech/weight-tracker/pkg/api"
)

type Server struct {
	router      *gin.Engine
	userService api.UserService
}

func NewServer(router *gin.Engine, userService api.UserService) *Server {
	return &Server{
		router:      router,
		userService: userService,
	}
}

func (s *Server) Run() error {
	// run funcion that inits the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
	}

	return nil
}
