// Code generated by ent, DO NOT EDIT.

package car

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the car type in the database.
	Label = "car"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldBrand holds the string denoting the brand field in the database.
	FieldBrand = "brand"
	// FieldModelYear holds the string denoting the model_year field in the database.
	FieldModelYear = "model_year"
	// Table holds the table name of the car in the database.
	Table = "cars"
)

// Columns holds all SQL columns for car fields.
var Columns = []string{
	FieldID,
	FieldNickname,
	FieldBrand,
	FieldModelYear,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cars"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_cars",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Car queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByBrand orders the results by the brand field.
func ByBrand(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBrand, opts...).ToFunc()
}

// ByModelYear orders the results by the model_year field.
func ByModelYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelYear, opts...).ToFunc()
}