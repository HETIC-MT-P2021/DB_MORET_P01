package models

// OrderDetails : Architecture in database
type OrderDetails struct {
	OrderNumber     int     `json:"orderNumber"`
	ProductCode     string  `json:"productCode"`
	QuantityOrdered int     `json:"quantityOrdered"`
	PriceEach       float32 `json:"priceEach"`
}

// TotalAndPriceProduct : Calculated with sql query
// NbProductOrdered : product * quantity (product and quantity are stored in db)
// TotalPrice : price of a product * quantity (price and quantity are stored in db)
type TotalAndPriceProduct struct {
	OrderID          int     `json:"orderNumber"`
	NbProductOrdered int     `json:"nbProduct"`
	TotalPrice       float32 `json:"totalPrice"`
}
