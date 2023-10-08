// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-example-rpc/ent/teacher"
)

// Teacher is the model entity for the Teacher schema.
type Teacher struct {
	config `json:"-"`
	// ID of the ent.
	// UUID
	ID uuid.UUID `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// AgeInt32 holds the value of the "age_int32" field.
	AgeInt32 int32 `json:"age_int32,omitempty"`
	// AgeInt64 holds the value of the "age_int64" field.
	AgeInt64 int64 `json:"age_int64,omitempty"`
	// AgeUint holds the value of the "age_uint" field.
	AgeUint uint `json:"age_uint,omitempty"`
	// AgeUint32 holds the value of the "age_uint32" field.
	AgeUint32 uint32 `json:"age_uint32,omitempty"`
	// AgeUint64 holds the value of the "age_uint64" field.
	AgeUint64 uint64 `json:"age_uint64,omitempty"`
	// WeightFloat holds the value of the "weight_float" field.
	WeightFloat float64 `json:"weight_float,omitempty"`
	// WeightFloat32 holds the value of the "weight_float32" field.
	WeightFloat32 float32 `json:"weight_float32,omitempty"`
	// ClassID holds the value of the "class_id" field.
	ClassID uuid.UUID `json:"class_id,omitempty"`
	// EnrollAt holds the value of the "enroll_at" field.
	EnrollAt time.Time `json:"enroll_at,omitempty"`
	// StatusBool holds the value of the "status_bool" field.
	StatusBool   bool `json:"status_bool,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Teacher) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case teacher.FieldStatusBool:
			values[i] = new(sql.NullBool)
		case teacher.FieldWeightFloat, teacher.FieldWeightFloat32:
			values[i] = new(sql.NullFloat64)
		case teacher.FieldAge, teacher.FieldAgeInt32, teacher.FieldAgeInt64, teacher.FieldAgeUint, teacher.FieldAgeUint32, teacher.FieldAgeUint64:
			values[i] = new(sql.NullInt64)
		case teacher.FieldName:
			values[i] = new(sql.NullString)
		case teacher.FieldCreatedAt, teacher.FieldUpdatedAt, teacher.FieldEnrollAt:
			values[i] = new(sql.NullTime)
		case teacher.FieldID, teacher.FieldClassID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Teacher fields.
func (t *Teacher) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case teacher.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case teacher.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case teacher.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case teacher.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case teacher.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				t.Age = int(value.Int64)
			}
		case teacher.FieldAgeInt32:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age_int32", values[i])
			} else if value.Valid {
				t.AgeInt32 = int32(value.Int64)
			}
		case teacher.FieldAgeInt64:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age_int64", values[i])
			} else if value.Valid {
				t.AgeInt64 = value.Int64
			}
		case teacher.FieldAgeUint:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age_uint", values[i])
			} else if value.Valid {
				t.AgeUint = uint(value.Int64)
			}
		case teacher.FieldAgeUint32:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age_uint32", values[i])
			} else if value.Valid {
				t.AgeUint32 = uint32(value.Int64)
			}
		case teacher.FieldAgeUint64:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age_uint64", values[i])
			} else if value.Valid {
				t.AgeUint64 = uint64(value.Int64)
			}
		case teacher.FieldWeightFloat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field weight_float", values[i])
			} else if value.Valid {
				t.WeightFloat = value.Float64
			}
		case teacher.FieldWeightFloat32:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field weight_float32", values[i])
			} else if value.Valid {
				t.WeightFloat32 = float32(value.Float64)
			}
		case teacher.FieldClassID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field class_id", values[i])
			} else if value != nil {
				t.ClassID = *value
			}
		case teacher.FieldEnrollAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field enroll_at", values[i])
			} else if value.Valid {
				t.EnrollAt = value.Time
			}
		case teacher.FieldStatusBool:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status_bool", values[i])
			} else if value.Valid {
				t.StatusBool = value.Bool
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Teacher.
// This includes values selected through modifiers, order, etc.
func (t *Teacher) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Teacher.
// Note that you need to call Teacher.Unwrap() before calling this method if this Teacher
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Teacher) Update() *TeacherUpdateOne {
	return NewTeacherClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Teacher entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Teacher) Unwrap() *Teacher {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Teacher is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Teacher) String() string {
	var builder strings.Builder
	builder.WriteString("Teacher(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", t.Age))
	builder.WriteString(", ")
	builder.WriteString("age_int32=")
	builder.WriteString(fmt.Sprintf("%v", t.AgeInt32))
	builder.WriteString(", ")
	builder.WriteString("age_int64=")
	builder.WriteString(fmt.Sprintf("%v", t.AgeInt64))
	builder.WriteString(", ")
	builder.WriteString("age_uint=")
	builder.WriteString(fmt.Sprintf("%v", t.AgeUint))
	builder.WriteString(", ")
	builder.WriteString("age_uint32=")
	builder.WriteString(fmt.Sprintf("%v", t.AgeUint32))
	builder.WriteString(", ")
	builder.WriteString("age_uint64=")
	builder.WriteString(fmt.Sprintf("%v", t.AgeUint64))
	builder.WriteString(", ")
	builder.WriteString("weight_float=")
	builder.WriteString(fmt.Sprintf("%v", t.WeightFloat))
	builder.WriteString(", ")
	builder.WriteString("weight_float32=")
	builder.WriteString(fmt.Sprintf("%v", t.WeightFloat32))
	builder.WriteString(", ")
	builder.WriteString("class_id=")
	builder.WriteString(fmt.Sprintf("%v", t.ClassID))
	builder.WriteString(", ")
	builder.WriteString("enroll_at=")
	builder.WriteString(t.EnrollAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status_bool=")
	builder.WriteString(fmt.Sprintf("%v", t.StatusBool))
	builder.WriteByte(')')
	return builder.String()
}

// Teachers is a parsable slice of Teacher.
type Teachers []*Teacher
