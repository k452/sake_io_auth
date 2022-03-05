package routing

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	//"sake_io_auth/conf"
	"sake_io_auth/model"
	"sake_io_auth/typeFile"
	"time"
)

func Routing(e *echo.Echo) {
	e.POST("/login", func(c echo.Context) error {
		var body typeFile.LoginType
		body.Mail = c.FormValue("mail")
		body.Pass = c.FormValue("pass")
		tokens := model.Login(body)

		if tokens.IDToken != "" && tokens.RefreshToken != "" {
			cookie := new(http.Cookie)
			cookie.Name = "IDToken"
			cookie.Value = tokens.IDToken
			// cookie.Domain = conf.GetDomain()
			cookie.Secure = true
			cookie.HttpOnly = true
			cookie.Expires = time.Now().Add(time.Minute * 30)
			c.SetCookie(cookie)

			cookie2 := new(http.Cookie)
			cookie2.Name = "RefreshToken"
			cookie2.Value = tokens.RefreshToken
			// cookie2.Domain = conf.GetDomain()
			cookie2.Secure = true
			cookie2.HttpOnly = true
			cookie2.Expires = time.Now().AddDate(0, 0, 1)
			c.SetCookie(cookie2)

			return c.JSON(http.StatusOK, "Success Setting Cookie Token")
		} else {
			return c.JSON(http.StatusUnauthorized, "")
		}
	})
	e.GET("/all", func(c echo.Context) error {
		for _, cookie := range c.Cookies() {
			fmt.Println(cookie.Name)
			fmt.Println(cookie.Value)
		}
		// cookie, err := c.Cookie("IDToken")
		// if err != nil {
		// 	return c.JSON(http.StatusInternalServerError, "")
		// }
		// fmt.Println(cookie.Value)
		return c.JSON(http.StatusOK, model.AllToken())
	})
}
