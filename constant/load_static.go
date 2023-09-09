package constant

import "github.com/labstack/echo/v4"

func LoadStatic(app *echo.Echo) {
	// load all static files
	app.Static("static", "../repository/assets")

}