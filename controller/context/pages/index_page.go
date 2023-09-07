package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func IndexContext(c echo.Context) error {
	// return c.String(http.StatusOK, "Changed the name of the package")
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
}


