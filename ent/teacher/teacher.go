// Code generated by ent, DO NOT EDIT.

package teacher

import (
	"time"

	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
)

const (
	// Label holds the string label denoting the teacher type in the database.
	Label = "teacher"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldAgeInt32 holds the string denoting the age_int32 field in the database.
	FieldAgeInt32 = "age_int32"
	// FieldAgeInt64 holds the string denoting the age_int64 field in the database.
	FieldAgeInt64 = "age_int64"
	// FieldAgeUint holds the string denoting the age_uint field in the database.
	FieldAgeUint = "age_uint"
	// FieldAgeUint32 holds the string denoting the age_uint32 field in the database.
	FieldAgeUint32 = "age_uint32"
	// FieldAgeUint64 holds the string denoting the age_uint64 field in the database.
	FieldAgeUint64 = "age_uint64"
	// FieldWeightFloat holds the string denoting the weight_float field in the database.
	FieldWeightFloat = "weight_float"
	// FieldWeightFloat32 holds the string denoting the weight_float32 field in the database.
	FieldWeightFloat32 = "weight_float32"
	// FieldClassID holds the string denoting the class_id field in the database.
	FieldClassID = "class_id"
	// FieldEnrollAt holds the string denoting the enroll_at field in the database.
	FieldEnrollAt = "enroll_at"
	// FieldStatusBool holds the string denoting the status_bool field in the database.
	FieldStatusBool = "status_bool"
	// Table holds the table name of the teacher in the database.
	Table = "teachers"
)

// Columns holds all SQL columns for teacher fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldAge,
	FieldAgeInt32,
	FieldAgeInt64,
	FieldAgeUint,
	FieldAgeUint32,
	FieldAgeUint64,
	FieldWeightFloat,
	FieldWeightFloat32,
	FieldClassID,
	FieldEnrollAt,
	FieldStatusBool,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Order defines the ordering method for the Teacher queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByAgeInt32 orders the results by the age_int32 field.
func ByAgeInt32(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldAgeInt32, opts...).ToFunc()
}

// ByAgeInt64 orders the results by the age_int64 field.
func ByAgeInt64(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldAgeInt64, opts...).ToFunc()
}

// ByAgeUint orders the results by the age_uint field.
func ByAgeUint(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldAgeUint, opts...).ToFunc()
}

// ByAgeUint32 orders the results by the age_uint32 field.
func ByAgeUint32(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldAgeUint32, opts...).ToFunc()
}

// ByAgeUint64 orders the results by the age_uint64 field.
func ByAgeUint64(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldAgeUint64, opts...).ToFunc()
}

// ByWeightFloat orders the results by the weight_float field.
func ByWeightFloat(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldWeightFloat, opts...).ToFunc()
}

// ByWeightFloat32 orders the results by the weight_float32 field.
func ByWeightFloat32(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldWeightFloat32, opts...).ToFunc()
}

// ByClassID orders the results by the class_id field.
func ByClassID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldClassID, opts...).ToFunc()
}

// ByEnrollAt orders the results by the enroll_at field.
func ByEnrollAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldEnrollAt, opts...).ToFunc()
}

// ByStatusBool orders the results by the status_bool field.
func ByStatusBool(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldStatusBool, opts...).ToFunc()
}
