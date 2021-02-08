package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	// group all routers under /v1/api
	v1 := router.Group("/v1/api")
	{
		v1.GET("/status", s.ApiStatus())
		// prefix the user Routes
		user := v1.Group("/user")
		{
			user.POST("", s.CreateUser())
		}
	}

	return router
}
