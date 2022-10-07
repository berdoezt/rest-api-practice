package routes

import (
	"rest-api-practice/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(ctl controller.Controller) error {
	r := gin.Default()

	r.GET("/persons", ctl.GetPersons)
	r.GET("/persons/:id")
	r.POST("/person", ctl.CreatePerson)
	r.PUT("/person/:id")
	r.DELETE("/person/:id")

	return r.Run("localhost:8080")
}
