package models

import "database/sql"

// NullString : Extend sql.NullString which allow to get
// {
//   String: "",
//   Valid: false
// }
// That process avoid error in program
// Ref : https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267
type NullString struct {
	sql.NullString
}
