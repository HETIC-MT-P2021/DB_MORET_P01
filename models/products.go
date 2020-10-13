package models

// Product : Architecture in db
type Product struct {
	ProductName        string `json:"productName"`
	ProductLine        string `json:"productLine"`
	ProductVendor      string `json:"productVendor"`
	ProductDescription string `json:"productDescription"`
	QuantityInStock    int    `json:"quantityInStock"`
}

// GetProduct : Get product details
func GetProduct(productCode string) (Product, error) {
	var product Product

	query := `
		SELECT
			products.productName,
			products.productLine,
			products.productVendor,
			products.productDescription,
			products.quantityInStock
		FROM products
		WHERE products.productCode = ?
	`

	err := DB.QueryRow(query, productCode).Scan(
		&product.ProductName,
		&product.ProductLine,
		&product.ProductVendor,
		&product.ProductDescription,
		&product.QuantityInStock)

	if err != nil {
		return product, err
	}

	return product, nil
}
