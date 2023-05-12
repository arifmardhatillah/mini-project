package main

import (
	"project_structure/config"
	"project_structure/middleware"
	"project_structure/route"

	"github.com/labstack/echo"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	middleware.Logmiddleware(e)

	route.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
