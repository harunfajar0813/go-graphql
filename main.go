package main

import (
	"graphi/graphql"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"graphi/datastore"
	"graphi/handler"
)

func main() {
	db, err := datastore.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.Welcome())

	hUser, err := graphql.NewUserHandler(db)
	if err != nil {
		log.Fatal(err)
	}
	e.POST("/graphql", echo.WrapHandler(hUser))

	if err := e.Start(":3000"); err != nil {
		log.Fatalln(err)
	}
}
