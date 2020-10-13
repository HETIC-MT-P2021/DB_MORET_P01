package models

// Order : Architecture in database
type Order struct {
	OrderNumber  int    `json:"orderNumber"`
	OrderDate    string `json:"orderDate"`
	RequiredDate string `json:"requiredDate"`
	ShippedDate  string `json:"shippedDate"`
	Status       string `json:"status"`
	Comments     string `json:"comments"`
}

// GetProductsCode : Get orderID and products code associated
func GetProductsCode(orderID int) (int, []string, error) {
	var order Order
	var orderDetails OrderDetails
	var orderProducts []string

	query := `
		SELECT
			orderdetails.orderNumber,
			orderdetails.productCode
		FROM orderdetails
		JOIN orders ON orderdetails.orderNumber = orders.orderNumber
		WHERE orders.orderNumber = ?
	`

	orderResults, err := DB.Query(query, orderID)

	if err != nil {
		return orderID, orderProducts, err
	}

	for orderResults.Next() {
		err := orderResults.Scan(
			&order.OrderNumber,
			&orderDetails.ProductCode)

		if err != nil {
			return orderID, orderProducts, err
		}

		orderProducts = append(orderProducts, orderDetails.ProductCode)
	}

	return orderID, orderProducts, nil
}

// GetTotalAndPriceProduct : Get total and price of all products ordered
func GetTotalAndPriceProduct(orderID int) (TotalAndPriceProduct, error) {
	var totalAndPriceProduct TotalAndPriceProduct

	query := `
		SELECT
		    orderdetails.orderNumber,
			SUM(orderdetails.quantityOrdered) AS nbProductOrdered,
			SUM(orderdetails.quantityOrdered * orderdetails.priceEach) AS totalPrice
		FROM orderdetails
		JOIN orders ON orderdetails.orderNumber = orders.orderNumber
		WHERE orders.orderNumber = ?
		GROUP BY orderdetails.orderNumber
	`

	err := DB.QueryRow(query, orderID).Scan(
		&totalAndPriceProduct.OrderID,
		&totalAndPriceProduct.NbProductOrdered,
		&totalAndPriceProduct.TotalPrice)

	if err != nil {
		return totalAndPriceProduct, err
	}

	return totalAndPriceProduct, nil
}
