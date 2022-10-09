package main

import (
	"log"
	handler "myapp/internal/handlers"
	repository "myapp/internal/repository"

	"github.com/labstack/echo"
)

func main() {
	if err := repository.OpenTable(); err != nil {
		log.Fatal(err)
	}
	router := echo.New()
	router.GET("/person", handler.GetAll)
	router.GET("/person/:id", handler.GetById)
	router.POST("/person", handler.Add)
	router.DELETE("/person/:id", handler.Delete)
	router.PUT("/person/:id", handler.Update)
	router.Logger.Fatal(router.Start(":8080"))
}
