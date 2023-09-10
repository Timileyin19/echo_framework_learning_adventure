package router

import (
	"github.com/labstack/echo/v4"
	"github.com/timileyin19/gofarm/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	http.IndexRouter(app)
}

