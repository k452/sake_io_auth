package main

import (
	"sake_io_auth/routing"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routing.Routing(e)
	e.Logger.Fatal(e.Start(":8090"))
}
