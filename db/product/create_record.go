package product

import (
	"fmt"
	"time"

	"github.com/timileyin19/gofarm/db"
	"github.com/timileyin19/gofarm/models"
)

func CreateRecord(productID string, cornColor string, quantity int64, timestamp time.Time) error {
	db, err := db.ConnectToDb()

	if err != nil {
		fmt.Println("Error connecting to database")
		return err
	}
	
	// migrate the schema
	db.AutoMigrate(&models.Products{})

	db.Create(&models.Products{
		ProductID: productID,
		Color: cornColor,
		Quantity: quantity,
		Timestamp: timestamp,
	})

	return nil
}