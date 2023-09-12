package http

import (
	"github.com/labstack/echo/v4"
	"github.com/timileyin19/gofarm/controller/context/pages"
)


func DetailsRouter(app *echo.Echo) {
	app.GET("/:productId", pages.DetailsContext)
}