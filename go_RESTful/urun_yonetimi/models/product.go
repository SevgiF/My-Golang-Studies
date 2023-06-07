package models

import "time"

// Represents the product for this app
// swagger:model
type Product struct {
	// the id for product
	// required: true
	// unique: true
	ID int `json:"id"`
	// the name for product name
	// required: true
	// min lenght: 5
	Name string `json:"name"`
	// the description for product description
	// required: false
	Description string `json:"description"`
	// the product creation time
	// example: 2022-01-01
	CreatedOn time.Time `json:"createdon"`
	// the product editing time
	// example: 2022-01-01
	ChangedOn time.Time `json:"changedon"`
}
