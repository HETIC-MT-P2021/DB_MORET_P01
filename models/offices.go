package models

// Office : Architecture in db
type Office struct {
	OfficeCode   string     `json:"officeCode"`
	City         string     `json:"city"`
	Phone        string     `json:"phone"`
	AddressLine1 string     `json:"addressLine1"`
	AddressLine2 NullString `json:"addressLine2"`
	State        NullString `json:"state"`
	Country      string     `json:"country"`
	PostalCode   string     `json:"postalCode"`
}

// OfficeWithEmployees : Office with employees infos
type OfficeWithEmployees struct {
	Office
	Employees []Employee `json:"employees"`
}

// GetOffices : Get infos offices
func GetOffices() ([]Office, error) {
	var offices []Office
	var office Office

	query := `
		SELECT
			offices.officeCode,
			offices.city,
			offices.phone,
			offices.addressLine1,
			offices.addressLine2,
			offices.state,
			offices.country,
			offices.postalCode
		FROM offices
	`

	customerResults, err := DB.Query(query)

	if err != nil {
		return offices, err
	}

	for customerResults.Next() {
		err := customerResults.Scan(
			&office.OfficeCode,
			&office.City,
			&office.Phone,
			&office.AddressLine1,
			&office.AddressLine2,
			&office.State,
			&office.Country,
			&office.PostalCode)

		if err != nil {
			return offices, err
		}

		offices = append(offices, office)
	}

	return offices, nil
}
