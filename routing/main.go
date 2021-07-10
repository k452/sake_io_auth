package routing

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"sake_io_auth/model"
	"sake_io_auth/typeFile"
)

func Routing(e *echo.Echo) {
	e.POST("/login", func(c echo.Context) error {
		var body typeFile.LoginType
		if err := c.Bind(&body); err != nil {
			logrus.Error(err)
		}
		return c.JSON(200, model.CreateToken(&body))
	})
}
