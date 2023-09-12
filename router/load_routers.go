package router

import (
	"github.com/labstack/echo/v4"
	"github.com/timileyin19/gofarm/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	http.IndexRouter(app)

	// form router 
	http.FormRouter(app)

	// details router
	http.DetailsRouter(app)
}

