// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-example-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-example-rpc/ent/teacher"
)

// TeacherUpdate is the builder for updating Teacher entities.
type TeacherUpdate struct {
	config
	hooks    []Hook
	mutation *TeacherMutation
}

// Where appends a list predicates to the TeacherUpdate builder.
func (tu *TeacherUpdate) Where(ps ...predicate.Teacher) *TeacherUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TeacherUpdate) SetUpdatedAt(t time.Time) *TeacherUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetName sets the "name" field.
func (tu *TeacherUpdate) SetName(s string) *TeacherUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetAge sets the "age" field.
func (tu *TeacherUpdate) SetAge(i int) *TeacherUpdate {
	tu.mutation.ResetAge()
	tu.mutation.SetAge(i)
	return tu
}

// AddAge adds i to the "age" field.
func (tu *TeacherUpdate) AddAge(i int) *TeacherUpdate {
	tu.mutation.AddAge(i)
	return tu
}

// SetAgeInt32 sets the "age_int32" field.
func (tu *TeacherUpdate) SetAgeInt32(i int32) *TeacherUpdate {
	tu.mutation.ResetAgeInt32()
	tu.mutation.SetAgeInt32(i)
	return tu
}

// AddAgeInt32 adds i to the "age_int32" field.
func (tu *TeacherUpdate) AddAgeInt32(i int32) *TeacherUpdate {
	tu.mutation.AddAgeInt32(i)
	return tu
}

// SetAgeInt64 sets the "age_int64" field.
func (tu *TeacherUpdate) SetAgeInt64(i int64) *TeacherUpdate {
	tu.mutation.ResetAgeInt64()
	tu.mutation.SetAgeInt64(i)
	return tu
}

// AddAgeInt64 adds i to the "age_int64" field.
func (tu *TeacherUpdate) AddAgeInt64(i int64) *TeacherUpdate {
	tu.mutation.AddAgeInt64(i)
	return tu
}

// SetAgeUint sets the "age_uint" field.
func (tu *TeacherUpdate) SetAgeUint(u uint) *TeacherUpdate {
	tu.mutation.ResetAgeUint()
	tu.mutation.SetAgeUint(u)
	return tu
}

// AddAgeUint adds u to the "age_uint" field.
func (tu *TeacherUpdate) AddAgeUint(u int) *TeacherUpdate {
	tu.mutation.AddAgeUint(u)
	return tu
}

// SetAgeUint32 sets the "age_uint32" field.
func (tu *TeacherUpdate) SetAgeUint32(u uint32) *TeacherUpdate {
	tu.mutation.ResetAgeUint32()
	tu.mutation.SetAgeUint32(u)
	return tu
}

// AddAgeUint32 adds u to the "age_uint32" field.
func (tu *TeacherUpdate) AddAgeUint32(u int32) *TeacherUpdate {
	tu.mutation.AddAgeUint32(u)
	return tu
}

// SetAgeUint64 sets the "age_uint64" field.
func (tu *TeacherUpdate) SetAgeUint64(u uint64) *TeacherUpdate {
	tu.mutation.ResetAgeUint64()
	tu.mutation.SetAgeUint64(u)
	return tu
}

// AddAgeUint64 adds u to the "age_uint64" field.
func (tu *TeacherUpdate) AddAgeUint64(u int64) *TeacherUpdate {
	tu.mutation.AddAgeUint64(u)
	return tu
}

// SetWeightFloat sets the "weight_float" field.
func (tu *TeacherUpdate) SetWeightFloat(f float64) *TeacherUpdate {
	tu.mutation.ResetWeightFloat()
	tu.mutation.SetWeightFloat(f)
	return tu
}

// AddWeightFloat adds f to the "weight_float" field.
func (tu *TeacherUpdate) AddWeightFloat(f float64) *TeacherUpdate {
	tu.mutation.AddWeightFloat(f)
	return tu
}

// SetWeightFloat32 sets the "weight_float32" field.
func (tu *TeacherUpdate) SetWeightFloat32(f float32) *TeacherUpdate {
	tu.mutation.ResetWeightFloat32()
	tu.mutation.SetWeightFloat32(f)
	return tu
}

// AddWeightFloat32 adds f to the "weight_float32" field.
func (tu *TeacherUpdate) AddWeightFloat32(f float32) *TeacherUpdate {
	tu.mutation.AddWeightFloat32(f)
	return tu
}

// SetClassID sets the "class_id" field.
func (tu *TeacherUpdate) SetClassID(u uuid.UUID) *TeacherUpdate {
	tu.mutation.SetClassID(u)
	return tu
}

// SetEnrollAt sets the "enroll_at" field.
func (tu *TeacherUpdate) SetEnrollAt(t time.Time) *TeacherUpdate {
	tu.mutation.SetEnrollAt(t)
	return tu
}

// SetStatusBool sets the "status_bool" field.
func (tu *TeacherUpdate) SetStatusBool(b bool) *TeacherUpdate {
	tu.mutation.SetStatusBool(b)
	return tu
}

// Mutation returns the TeacherMutation object of the builder.
func (tu *TeacherUpdate) Mutation() *TeacherMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeacherUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeacherUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeacherUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeacherUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TeacherUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := teacher.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

func (tu *TeacherUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(teacher.Table, teacher.Columns, sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(teacher.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(teacher.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Age(); ok {
		_spec.SetField(teacher.FieldAge, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedAge(); ok {
		_spec.AddField(teacher.FieldAge, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AgeInt32(); ok {
		_spec.SetField(teacher.FieldAgeInt32, field.TypeInt32, value)
	}
	if value, ok := tu.mutation.AddedAgeInt32(); ok {
		_spec.AddField(teacher.FieldAgeInt32, field.TypeInt32, value)
	}
	if value, ok := tu.mutation.AgeInt64(); ok {
		_spec.SetField(teacher.FieldAgeInt64, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedAgeInt64(); ok {
		_spec.AddField(teacher.FieldAgeInt64, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AgeUint(); ok {
		_spec.SetField(teacher.FieldAgeUint, field.TypeUint, value)
	}
	if value, ok := tu.mutation.AddedAgeUint(); ok {
		_spec.AddField(teacher.FieldAgeUint, field.TypeUint, value)
	}
	if value, ok := tu.mutation.AgeUint32(); ok {
		_spec.SetField(teacher.FieldAgeUint32, field.TypeUint32, value)
	}
	if value, ok := tu.mutation.AddedAgeUint32(); ok {
		_spec.AddField(teacher.FieldAgeUint32, field.TypeUint32, value)
	}
	if value, ok := tu.mutation.AgeUint64(); ok {
		_spec.SetField(teacher.FieldAgeUint64, field.TypeUint64, value)
	}
	if value, ok := tu.mutation.AddedAgeUint64(); ok {
		_spec.AddField(teacher.FieldAgeUint64, field.TypeUint64, value)
	}
	if value, ok := tu.mutation.WeightFloat(); ok {
		_spec.SetField(teacher.FieldWeightFloat, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedWeightFloat(); ok {
		_spec.AddField(teacher.FieldWeightFloat, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.WeightFloat32(); ok {
		_spec.SetField(teacher.FieldWeightFloat32, field.TypeFloat32, value)
	}
	if value, ok := tu.mutation.AddedWeightFloat32(); ok {
		_spec.AddField(teacher.FieldWeightFloat32, field.TypeFloat32, value)
	}
	if value, ok := tu.mutation.ClassID(); ok {
		_spec.SetField(teacher.FieldClassID, field.TypeUUID, value)
	}
	if value, ok := tu.mutation.EnrollAt(); ok {
		_spec.SetField(teacher.FieldEnrollAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.StatusBool(); ok {
		_spec.SetField(teacher.FieldStatusBool, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teacher.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TeacherUpdateOne is the builder for updating a single Teacher entity.
type TeacherUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeacherMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TeacherUpdateOne) SetUpdatedAt(t time.Time) *TeacherUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TeacherUpdateOne) SetName(s string) *TeacherUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetAge sets the "age" field.
func (tuo *TeacherUpdateOne) SetAge(i int) *TeacherUpdateOne {
	tuo.mutation.ResetAge()
	tuo.mutation.SetAge(i)
	return tuo
}

// AddAge adds i to the "age" field.
func (tuo *TeacherUpdateOne) AddAge(i int) *TeacherUpdateOne {
	tuo.mutation.AddAge(i)
	return tuo
}

// SetAgeInt32 sets the "age_int32" field.
func (tuo *TeacherUpdateOne) SetAgeInt32(i int32) *TeacherUpdateOne {
	tuo.mutation.ResetAgeInt32()
	tuo.mutation.SetAgeInt32(i)
	return tuo
}

// AddAgeInt32 adds i to the "age_int32" field.
func (tuo *TeacherUpdateOne) AddAgeInt32(i int32) *TeacherUpdateOne {
	tuo.mutation.AddAgeInt32(i)
	return tuo
}

// SetAgeInt64 sets the "age_int64" field.
func (tuo *TeacherUpdateOne) SetAgeInt64(i int64) *TeacherUpdateOne {
	tuo.mutation.ResetAgeInt64()
	tuo.mutation.SetAgeInt64(i)
	return tuo
}

// AddAgeInt64 adds i to the "age_int64" field.
func (tuo *TeacherUpdateOne) AddAgeInt64(i int64) *TeacherUpdateOne {
	tuo.mutation.AddAgeInt64(i)
	return tuo
}

// SetAgeUint sets the "age_uint" field.
func (tuo *TeacherUpdateOne) SetAgeUint(u uint) *TeacherUpdateOne {
	tuo.mutation.ResetAgeUint()
	tuo.mutation.SetAgeUint(u)
	return tuo
}

// AddAgeUint adds u to the "age_uint" field.
func (tuo *TeacherUpdateOne) AddAgeUint(u int) *TeacherUpdateOne {
	tuo.mutation.AddAgeUint(u)
	return tuo
}

// SetAgeUint32 sets the "age_uint32" field.
func (tuo *TeacherUpdateOne) SetAgeUint32(u uint32) *TeacherUpdateOne {
	tuo.mutation.ResetAgeUint32()
	tuo.mutation.SetAgeUint32(u)
	return tuo
}

// AddAgeUint32 adds u to the "age_uint32" field.
func (tuo *TeacherUpdateOne) AddAgeUint32(u int32) *TeacherUpdateOne {
	tuo.mutation.AddAgeUint32(u)
	return tuo
}

// SetAgeUint64 sets the "age_uint64" field.
func (tuo *TeacherUpdateOne) SetAgeUint64(u uint64) *TeacherUpdateOne {
	tuo.mutation.ResetAgeUint64()
	tuo.mutation.SetAgeUint64(u)
	return tuo
}

// AddAgeUint64 adds u to the "age_uint64" field.
func (tuo *TeacherUpdateOne) AddAgeUint64(u int64) *TeacherUpdateOne {
	tuo.mutation.AddAgeUint64(u)
	return tuo
}

// SetWeightFloat sets the "weight_float" field.
func (tuo *TeacherUpdateOne) SetWeightFloat(f float64) *TeacherUpdateOne {
	tuo.mutation.ResetWeightFloat()
	tuo.mutation.SetWeightFloat(f)
	return tuo
}

// AddWeightFloat adds f to the "weight_float" field.
func (tuo *TeacherUpdateOne) AddWeightFloat(f float64) *TeacherUpdateOne {
	tuo.mutation.AddWeightFloat(f)
	return tuo
}

// SetWeightFloat32 sets the "weight_float32" field.
func (tuo *TeacherUpdateOne) SetWeightFloat32(f float32) *TeacherUpdateOne {
	tuo.mutation.ResetWeightFloat32()
	tuo.mutation.SetWeightFloat32(f)
	return tuo
}

// AddWeightFloat32 adds f to the "weight_float32" field.
func (tuo *TeacherUpdateOne) AddWeightFloat32(f float32) *TeacherUpdateOne {
	tuo.mutation.AddWeightFloat32(f)
	return tuo
}

// SetClassID sets the "class_id" field.
func (tuo *TeacherUpdateOne) SetClassID(u uuid.UUID) *TeacherUpdateOne {
	tuo.mutation.SetClassID(u)
	return tuo
}

// SetEnrollAt sets the "enroll_at" field.
func (tuo *TeacherUpdateOne) SetEnrollAt(t time.Time) *TeacherUpdateOne {
	tuo.mutation.SetEnrollAt(t)
	return tuo
}

// SetStatusBool sets the "status_bool" field.
func (tuo *TeacherUpdateOne) SetStatusBool(b bool) *TeacherUpdateOne {
	tuo.mutation.SetStatusBool(b)
	return tuo
}

// Mutation returns the TeacherMutation object of the builder.
func (tuo *TeacherUpdateOne) Mutation() *TeacherMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TeacherUpdate builder.
func (tuo *TeacherUpdateOne) Where(ps ...predicate.Teacher) *TeacherUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeacherUpdateOne) Select(field string, fields ...string) *TeacherUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Teacher entity.
func (tuo *TeacherUpdateOne) Save(ctx context.Context) (*Teacher, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeacherUpdateOne) SaveX(ctx context.Context) *Teacher {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeacherUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeacherUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TeacherUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := teacher.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

func (tuo *TeacherUpdateOne) sqlSave(ctx context.Context) (_node *Teacher, err error) {
	_spec := sqlgraph.NewUpdateSpec(teacher.Table, teacher.Columns, sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Teacher.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teacher.FieldID)
		for _, f := range fields {
			if !teacher.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != teacher.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(teacher.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(teacher.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Age(); ok {
		_spec.SetField(teacher.FieldAge, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedAge(); ok {
		_spec.AddField(teacher.FieldAge, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AgeInt32(); ok {
		_spec.SetField(teacher.FieldAgeInt32, field.TypeInt32, value)
	}
	if value, ok := tuo.mutation.AddedAgeInt32(); ok {
		_spec.AddField(teacher.FieldAgeInt32, field.TypeInt32, value)
	}
	if value, ok := tuo.mutation.AgeInt64(); ok {
		_spec.SetField(teacher.FieldAgeInt64, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedAgeInt64(); ok {
		_spec.AddField(teacher.FieldAgeInt64, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AgeUint(); ok {
		_spec.SetField(teacher.FieldAgeUint, field.TypeUint, value)
	}
	if value, ok := tuo.mutation.AddedAgeUint(); ok {
		_spec.AddField(teacher.FieldAgeUint, field.TypeUint, value)
	}
	if value, ok := tuo.mutation.AgeUint32(); ok {
		_spec.SetField(teacher.FieldAgeUint32, field.TypeUint32, value)
	}
	if value, ok := tuo.mutation.AddedAgeUint32(); ok {
		_spec.AddField(teacher.FieldAgeUint32, field.TypeUint32, value)
	}
	if value, ok := tuo.mutation.AgeUint64(); ok {
		_spec.SetField(teacher.FieldAgeUint64, field.TypeUint64, value)
	}
	if value, ok := tuo.mutation.AddedAgeUint64(); ok {
		_spec.AddField(teacher.FieldAgeUint64, field.TypeUint64, value)
	}
	if value, ok := tuo.mutation.WeightFloat(); ok {
		_spec.SetField(teacher.FieldWeightFloat, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedWeightFloat(); ok {
		_spec.AddField(teacher.FieldWeightFloat, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.WeightFloat32(); ok {
		_spec.SetField(teacher.FieldWeightFloat32, field.TypeFloat32, value)
	}
	if value, ok := tuo.mutation.AddedWeightFloat32(); ok {
		_spec.AddField(teacher.FieldWeightFloat32, field.TypeFloat32, value)
	}
	if value, ok := tuo.mutation.ClassID(); ok {
		_spec.SetField(teacher.FieldClassID, field.TypeUUID, value)
	}
	if value, ok := tuo.mutation.EnrollAt(); ok {
		_spec.SetField(teacher.FieldEnrollAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.StatusBool(); ok {
		_spec.SetField(teacher.FieldStatusBool, field.TypeBool, value)
	}
	_node = &Teacher{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teacher.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
