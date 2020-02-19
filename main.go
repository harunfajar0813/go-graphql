package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"graphi/autoload"
	"graphi/datastore"
	"graphi/graphql"
)

func init() {
	autoload.Load()
}

func main() {
	db, err := datastore.NewMyqlDB()
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	allHandlers, err := graphql.NewUserHandler(db)
	if err != nil {
		log.Fatal(err)
	}
	e.POST("/graphql", echo.WrapHandler(allHandlers))

	if err := e.Start(":3000"); err != nil {
		log.Fatalln(err)
	}
}
