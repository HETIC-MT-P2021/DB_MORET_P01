package models


import (
	_ "database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Customer struct {
	CustomerNumber int `json:"customerNumber"`
	CustomerName string `json:"customerName"`
}

func GetCustomer(customerId int) Customer {
	var customer Customer

	ConnectDB()
	rows, err := DB.Query("SELECT c.customerName," +
		"c.contactFirstName," +
		"c.contactLastName," +
		"c.addressLine1," +
		"c.addressLine2," +
		"c.city," +
		"c.postalCode," +
		"c.country," +
		"c.phone" +
		"FROM customers AS c" +
		"WHERE c.customerNumber = 119;")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&customer.CustomerNumber, &customer.CustomerName)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(customer)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return customer
}
