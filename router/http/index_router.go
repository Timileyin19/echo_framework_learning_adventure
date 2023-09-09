package http

import (
	"github.com/labstack/echo/v4"
	context "github.com/timileyin19/go_farm/controller/context/pages"
)



func IndexRouter(app *echo.Echo) {
	app.GET("/", context.IndexContext)
}