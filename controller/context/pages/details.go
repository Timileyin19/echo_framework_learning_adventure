package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
	productDb "github.com/timileyin19/gofarm/db/product"
)

func DetailsContext(c echo.Context) error {
	productId := c.Param("productId")

	dbRecord, _ := productDb.GetProductDetails(productId)

	return c.Render(http.StatusOK, "detail.html", dbRecord)
}