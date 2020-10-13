# DB Exercise

Exercise to build a GoLang API with some employees and offices datas

## Install and run project

```bash
go build
go run main.go
```

## API Routes

```
Get all customers
GET    /api/customers/
```

```
Get customer by its id
GET    /api/customers/{id}
```

```
Get orders of a customer
GET    /api/customers/{id}/orders 
```

```
Get all orders
GET    /api/orders/              
```

```
Get all employees
GET    /api/employees/           
```

```http
Get an employee details by its id
GET    /api/employees/{id}        
```

```http
Get offices and their employees
GET    /api/offices/             
```

```http
Get employees by their office id
GET    /api/offices/{id}
```

## License
[MIT](https://github.com/HETIC-MT-P2021/DB_MORET_P01/blob/master/LICENSE)
