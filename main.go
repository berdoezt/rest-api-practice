package main

import (
	"fmt"
	"rest-api-practice/controller"
	"rest-api-practice/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Start()
	if err != nil {
		fmt.Println("error start database", err)
		return
	}

	ctl := controller.New(db)

	r := gin.Default()

	r.GET("/persons", ctl.GetPersons)
	r.GET("/persons/:id")
	r.POST("/person", ctl.CreatePerson)
	r.PUT("/person/:id")
	r.DELETE("/person/:id")

	r.Run("localhost:8080")
}
