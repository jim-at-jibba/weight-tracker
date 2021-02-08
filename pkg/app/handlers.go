package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"jamesbest.tech/weight-tracker/pkg/api"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "weight tracker API running",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		// body, _ := ioutil.ReadAll(c.Request.Body)
		// println(string(body))
		var newUser api.NewUserRequest

		err := c.ShouldBindJSON(&newUser)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		log.Printf("Hello %+v", newUser)
		err = s.userService.New(newUser)

		if err != nil {

			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new user created",
		}

		c.JSON(http.StatusOK, response)
	}
}
