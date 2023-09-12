package product

import (
	"fmt"

	"github.com/timileyin19/gofarm/db"
	"github.com/timileyin19/gofarm/models"
)

func GetProductDetails(productID string) (map[string]interface{}, error) {
	db, err := db.ConnectToDb()

	if err != nil {
		fmt.Println("Failed to connect to database")
		return map[string]interface{}{}, err
	}

	var product models.Products

	// query the dabase for the product
	result := db.Where("product_id = ?", productID).First(&product)

	if result.Error == nil {
		return map[string]interface{}{
			"product_id":   product.ProductID,
			"color":        product.Color,
			"quantity":     product.Quantity,
			"timestamp":    product.Timestamp,
		}, nil
	}

	return map[string]interface{}{}, err
}