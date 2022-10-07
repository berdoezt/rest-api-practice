package main

import (
	"fmt"
	"rest-api-practice/controller"
	"rest-api-practice/database"
	"rest-api-practice/routes"
)

func main() {
	db, err := database.Start()
	if err != nil {
		fmt.Println("error start database", err)
		return
	}

	ctl := controller.New(db)

	err = routes.StartServer(ctl)
	if err != nil {
		fmt.Println("error start server", err)
		return
	}
}
