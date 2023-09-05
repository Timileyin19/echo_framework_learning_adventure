package router

import (
	"github.com/labstack/echo/v4"
	"github.com/timileyin19/echo_framework/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	http.IndexRouter(app)
}

