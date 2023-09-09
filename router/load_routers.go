package router

import (
	"github.com/labstack/echo/v4"
	"github.com/timileyin19/go_farm/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	http.IndexRouter(app)
}

