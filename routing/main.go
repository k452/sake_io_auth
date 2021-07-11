package routing

import (
	"sake_io_auth/model"
	"sake_io_auth/typeFile"

	"github.com/labstack/echo"
)

func Routing(e *echo.Echo) {
	e.POST("/login", func(c echo.Context) error {
		var body typeFile.LoginType
		body.Mail = c.FormValue("mail")
		body.Pass = c.FormValue("pass")
		return c.JSON(200, model.CreateToken(body))
	})
}
