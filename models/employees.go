package models

// Employee : Architecture in db
type Employee struct {
	EmployeeNumber int    `json:"employeeNumber"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	Extension      string `json:"extension"`
	Email          string `json:"email"`
	JobTitle       string `json:"jobTitle"`
}

// EmployeeWithOffice : Employee with office associated
type EmployeeWithOffice struct {
	Employee
	Office Office `json:"office"`
}

// GetEmployeesWithOffice : Get all amployees with office associated
func GetEmployeesWithOffice() ([]EmployeeWithOffice, error) {
	var employeesWithOffice []EmployeeWithOffice
	var employeeWithOffice EmployeeWithOffice

	query := `
		SELECT
			offices.officeCode,
			offices.city,
			offices.phone,
			offices.addressLine1,
			offices.addressLine2,
			offices.state,
			offices.country,
			offices.postalCode,
			employees.employeeNumber,
			employees.lastName,
			employees.firstName,
			employees.extension,
			employees.email,
			employees.jobTitle
		FROM offices
		JOIN employees ON offices.officeCode = employees.officeCode
		WHERE employees.employeeNumber
	`
	employeesResult, err := DB.Query(query)

	if err != nil {
		return employeesWithOffice, err
	}

	for employeesResult.Next() {
		err := employeesResult.Scan(
			&employeeWithOffice.Office.OfficeCode,
			&employeeWithOffice.Office.City,
			&employeeWithOffice.Office.Phone,
			&employeeWithOffice.Office.AddressLine1,
			&employeeWithOffice.Office.AddressLine2,
			&employeeWithOffice.Office.State,
			&employeeWithOffice.Office.Country,
			&employeeWithOffice.Office.PostalCode,
			&employeeWithOffice.EmployeeNumber,
			&employeeWithOffice.LastName,
			&employeeWithOffice.FirstName,
			&employeeWithOffice.Extension,
			&employeeWithOffice.Email,
			&employeeWithOffice.JobTitle)

		if err != nil {
			return employeesWithOffice, err
		}

		employeesWithOffice = append(employeesWithOffice, employeeWithOffice)
	}

	return employeesWithOffice, nil
}

// GetEmployeesOffice : Get all employess which are associated to an office
func GetEmployeesOffice(officeCode string) ([]Employee, error) {
	var employee Employee
	var employees []Employee

	query := `
		SELECT
			employees.employeeNumber,
			employees.lastName,
			employees.firstName,
			employees.extension,
			employees.email,
			employees.jobTitle
		FROM employees
		JOIN offices ON employees.officeCode = offices.officeCode
		WHERE offices.officeCode = ?
	`

	employeesResult, err := DB.Query(query, officeCode)

	if err != nil {
		return employees, err
	}

	for employeesResult.Next() {
		err := employeesResult.Scan(
			&employee.EmployeeNumber,
			&employee.LastName,
			&employee.FirstName,
			&employee.Extension,
			&employee.Email,
			&employee.JobTitle)

		if err != nil {
			return employees, err
		}

		employees = append(employees, employee)
	}

	return employees, nil

}
