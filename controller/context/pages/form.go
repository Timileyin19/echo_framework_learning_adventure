package pages

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	productDatabase "github.com/timileyin19/gofarm/db/product"
)

func FormContext(c echo.Context) error {
	productId := c.FormValue("product_id")	
	quantity := c.FormValue("quantity")
	cornColor := c.FormValue("corn_color")

	qtyToNum, _ := strconv.ParseInt(quantity, 10, 64)

	// current time
	currentTime := time.Now()

	productDatabase.CreateRecord(productId, cornColor, qtyToNum, currentTime)

	pathUrl := fmt.Sprintf("/%v", productId)

	return c.Redirect(http.StatusMovedPermanently, pathUrl)
}