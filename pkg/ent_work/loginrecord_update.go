// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/loginrecord"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/predicate"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/user"
)

// LoginRecordUpdate is the builder for updating LoginRecord entities.
type LoginRecordUpdate struct {
	config
	hooks     []Hook
	mutation  *LoginRecordMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LoginRecordUpdate builder.
func (lru *LoginRecordUpdate) Where(ps ...predicate.LoginRecord) *LoginRecordUpdate {
	lru.mutation.Where(ps...)
	return lru
}

// SetCreatedBy sets the "created_by" field.
func (lru *LoginRecordUpdate) SetCreatedBy(i int64) *LoginRecordUpdate {
	lru.mutation.ResetCreatedBy()
	lru.mutation.SetCreatedBy(i)
	return lru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (lru *LoginRecordUpdate) SetNillableCreatedBy(i *int64) *LoginRecordUpdate {
	if i != nil {
		lru.SetCreatedBy(*i)
	}
	return lru
}

// AddCreatedBy adds i to the "created_by" field.
func (lru *LoginRecordUpdate) AddCreatedBy(i int64) *LoginRecordUpdate {
	lru.mutation.AddCreatedBy(i)
	return lru
}

// SetUpdatedBy sets the "updated_by" field.
func (lru *LoginRecordUpdate) SetUpdatedBy(i int64) *LoginRecordUpdate {
	lru.mutation.ResetUpdatedBy()
	lru.mutation.SetUpdatedBy(i)
	return lru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (lru *LoginRecordUpdate) SetNillableUpdatedBy(i *int64) *LoginRecordUpdate {
	if i != nil {
		lru.SetUpdatedBy(*i)
	}
	return lru
}

// AddUpdatedBy adds i to the "updated_by" field.
func (lru *LoginRecordUpdate) AddUpdatedBy(i int64) *LoginRecordUpdate {
	lru.mutation.AddUpdatedBy(i)
	return lru
}

// SetUpdatedAt sets the "updated_at" field.
func (lru *LoginRecordUpdate) SetUpdatedAt(t time.Time) *LoginRecordUpdate {
	lru.mutation.SetUpdatedAt(t)
	return lru
}

// SetDeletedAt sets the "deleted_at" field.
func (lru *LoginRecordUpdate) SetDeletedAt(t time.Time) *LoginRecordUpdate {
	lru.mutation.SetDeletedAt(t)
	return lru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lru *LoginRecordUpdate) SetNillableDeletedAt(t *time.Time) *LoginRecordUpdate {
	if t != nil {
		lru.SetDeletedAt(*t)
	}
	return lru
}

// SetUa sets the "ua" field.
func (lru *LoginRecordUpdate) SetUa(s string) *LoginRecordUpdate {
	lru.mutation.SetUa(s)
	return lru
}

// SetNillableUa sets the "ua" field if the given value is not nil.
func (lru *LoginRecordUpdate) SetNillableUa(s *string) *LoginRecordUpdate {
	if s != nil {
		lru.SetUa(*s)
	}
	return lru
}

// SetIP sets the "ip" field.
func (lru *LoginRecordUpdate) SetIP(s string) *LoginRecordUpdate {
	lru.mutation.SetIP(s)
	return lru
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (lru *LoginRecordUpdate) SetNillableIP(s *string) *LoginRecordUpdate {
	if s != nil {
		lru.SetIP(*s)
	}
	return lru
}

// SetWay sets the "way" field.
func (lru *LoginRecordUpdate) SetWay(s string) *LoginRecordUpdate {
	lru.mutation.SetWay(s)
	return lru
}

// SetNillableWay sets the "way" field if the given value is not nil.
func (lru *LoginRecordUpdate) SetNillableWay(s *string) *LoginRecordUpdate {
	if s != nil {
		lru.SetWay(*s)
	}
	return lru
}

// SetUserID sets the "user_id" field.
func (lru *LoginRecordUpdate) SetUserID(i int64) *LoginRecordUpdate {
	lru.mutation.SetUserID(i)
	return lru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (lru *LoginRecordUpdate) SetNillableUserID(i *int64) *LoginRecordUpdate {
	if i != nil {
		lru.SetUserID(*i)
	}
	return lru
}

// SetUser sets the "user" edge to the User entity.
func (lru *LoginRecordUpdate) SetUser(u *User) *LoginRecordUpdate {
	return lru.SetUserID(u.ID)
}

// Mutation returns the LoginRecordMutation object of the builder.
func (lru *LoginRecordUpdate) Mutation() *LoginRecordMutation {
	return lru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (lru *LoginRecordUpdate) ClearUser() *LoginRecordUpdate {
	lru.mutation.ClearUser()
	return lru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lru *LoginRecordUpdate) Save(ctx context.Context) (int, error) {
	lru.defaults()
	return withHooks(ctx, lru.sqlSave, lru.mutation, lru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lru *LoginRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := lru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lru *LoginRecordUpdate) Exec(ctx context.Context) error {
	_, err := lru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lru *LoginRecordUpdate) ExecX(ctx context.Context) {
	if err := lru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lru *LoginRecordUpdate) defaults() {
	if _, ok := lru.mutation.UpdatedAt(); !ok {
		v := loginrecord.UpdateDefaultUpdatedAt()
		lru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lru *LoginRecordUpdate) check() error {
	if _, ok := lru.mutation.UserID(); lru.mutation.UserCleared() && !ok {
		return errors.New(`ent_work: clearing a required unique edge "LoginRecord.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lru *LoginRecordUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LoginRecordUpdate {
	lru.modifiers = append(lru.modifiers, modifiers...)
	return lru
}

func (lru *LoginRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(loginrecord.Table, loginrecord.Columns, sqlgraph.NewFieldSpec(loginrecord.FieldID, field.TypeInt64))
	if ps := lru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lru.mutation.CreatedBy(); ok {
		_spec.SetField(loginrecord.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := lru.mutation.AddedCreatedBy(); ok {
		_spec.AddField(loginrecord.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := lru.mutation.UpdatedBy(); ok {
		_spec.SetField(loginrecord.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := lru.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(loginrecord.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := lru.mutation.UpdatedAt(); ok {
		_spec.SetField(loginrecord.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lru.mutation.DeletedAt(); ok {
		_spec.SetField(loginrecord.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := lru.mutation.Ua(); ok {
		_spec.SetField(loginrecord.FieldUa, field.TypeString, value)
	}
	if value, ok := lru.mutation.IP(); ok {
		_spec.SetField(loginrecord.FieldIP, field.TypeString, value)
	}
	if value, ok := lru.mutation.Way(); ok {
		_spec.SetField(loginrecord.FieldWay, field.TypeString, value)
	}
	if lru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loginrecord.UserTable,
			Columns: []string{loginrecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loginrecord.UserTable,
			Columns: []string{loginrecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(lru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, lru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loginrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lru.mutation.done = true
	return n, nil
}

// LoginRecordUpdateOne is the builder for updating a single LoginRecord entity.
type LoginRecordUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LoginRecordMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (lruo *LoginRecordUpdateOne) SetCreatedBy(i int64) *LoginRecordUpdateOne {
	lruo.mutation.ResetCreatedBy()
	lruo.mutation.SetCreatedBy(i)
	return lruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (lruo *LoginRecordUpdateOne) SetNillableCreatedBy(i *int64) *LoginRecordUpdateOne {
	if i != nil {
		lruo.SetCreatedBy(*i)
	}
	return lruo
}

// AddCreatedBy adds i to the "created_by" field.
func (lruo *LoginRecordUpdateOne) AddCreatedBy(i int64) *LoginRecordUpdateOne {
	lruo.mutation.AddCreatedBy(i)
	return lruo
}

// SetUpdatedBy sets the "updated_by" field.
func (lruo *LoginRecordUpdateOne) SetUpdatedBy(i int64) *LoginRecordUpdateOne {
	lruo.mutation.ResetUpdatedBy()
	lruo.mutation.SetUpdatedBy(i)
	return lruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (lruo *LoginRecordUpdateOne) SetNillableUpdatedBy(i *int64) *LoginRecordUpdateOne {
	if i != nil {
		lruo.SetUpdatedBy(*i)
	}
	return lruo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (lruo *LoginRecordUpdateOne) AddUpdatedBy(i int64) *LoginRecordUpdateOne {
	lruo.mutation.AddUpdatedBy(i)
	return lruo
}

// SetUpdatedAt sets the "updated_at" field.
func (lruo *LoginRecordUpdateOne) SetUpdatedAt(t time.Time) *LoginRecordUpdateOne {
	lruo.mutation.SetUpdatedAt(t)
	return lruo
}

// SetDeletedAt sets the "deleted_at" field.
func (lruo *LoginRecordUpdateOne) SetDeletedAt(t time.Time) *LoginRecordUpdateOne {
	lruo.mutation.SetDeletedAt(t)
	return lruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lruo *LoginRecordUpdateOne) SetNillableDeletedAt(t *time.Time) *LoginRecordUpdateOne {
	if t != nil {
		lruo.SetDeletedAt(*t)
	}
	return lruo
}

// SetUa sets the "ua" field.
func (lruo *LoginRecordUpdateOne) SetUa(s string) *LoginRecordUpdateOne {
	lruo.mutation.SetUa(s)
	return lruo
}

// SetNillableUa sets the "ua" field if the given value is not nil.
func (lruo *LoginRecordUpdateOne) SetNillableUa(s *string) *LoginRecordUpdateOne {
	if s != nil {
		lruo.SetUa(*s)
	}
	return lruo
}

// SetIP sets the "ip" field.
func (lruo *LoginRecordUpdateOne) SetIP(s string) *LoginRecordUpdateOne {
	lruo.mutation.SetIP(s)
	return lruo
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (lruo *LoginRecordUpdateOne) SetNillableIP(s *string) *LoginRecordUpdateOne {
	if s != nil {
		lruo.SetIP(*s)
	}
	return lruo
}

// SetWay sets the "way" field.
func (lruo *LoginRecordUpdateOne) SetWay(s string) *LoginRecordUpdateOne {
	lruo.mutation.SetWay(s)
	return lruo
}

// SetNillableWay sets the "way" field if the given value is not nil.
func (lruo *LoginRecordUpdateOne) SetNillableWay(s *string) *LoginRecordUpdateOne {
	if s != nil {
		lruo.SetWay(*s)
	}
	return lruo
}

// SetUserID sets the "user_id" field.
func (lruo *LoginRecordUpdateOne) SetUserID(i int64) *LoginRecordUpdateOne {
	lruo.mutation.SetUserID(i)
	return lruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (lruo *LoginRecordUpdateOne) SetNillableUserID(i *int64) *LoginRecordUpdateOne {
	if i != nil {
		lruo.SetUserID(*i)
	}
	return lruo
}

// SetUser sets the "user" edge to the User entity.
func (lruo *LoginRecordUpdateOne) SetUser(u *User) *LoginRecordUpdateOne {
	return lruo.SetUserID(u.ID)
}

// Mutation returns the LoginRecordMutation object of the builder.
func (lruo *LoginRecordUpdateOne) Mutation() *LoginRecordMutation {
	return lruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (lruo *LoginRecordUpdateOne) ClearUser() *LoginRecordUpdateOne {
	lruo.mutation.ClearUser()
	return lruo
}

// Where appends a list predicates to the LoginRecordUpdate builder.
func (lruo *LoginRecordUpdateOne) Where(ps ...predicate.LoginRecord) *LoginRecordUpdateOne {
	lruo.mutation.Where(ps...)
	return lruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lruo *LoginRecordUpdateOne) Select(field string, fields ...string) *LoginRecordUpdateOne {
	lruo.fields = append([]string{field}, fields...)
	return lruo
}

// Save executes the query and returns the updated LoginRecord entity.
func (lruo *LoginRecordUpdateOne) Save(ctx context.Context) (*LoginRecord, error) {
	lruo.defaults()
	return withHooks(ctx, lruo.sqlSave, lruo.mutation, lruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lruo *LoginRecordUpdateOne) SaveX(ctx context.Context) *LoginRecord {
	node, err := lruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lruo *LoginRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := lruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lruo *LoginRecordUpdateOne) ExecX(ctx context.Context) {
	if err := lruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lruo *LoginRecordUpdateOne) defaults() {
	if _, ok := lruo.mutation.UpdatedAt(); !ok {
		v := loginrecord.UpdateDefaultUpdatedAt()
		lruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lruo *LoginRecordUpdateOne) check() error {
	if _, ok := lruo.mutation.UserID(); lruo.mutation.UserCleared() && !ok {
		return errors.New(`ent_work: clearing a required unique edge "LoginRecord.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lruo *LoginRecordUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LoginRecordUpdateOne {
	lruo.modifiers = append(lruo.modifiers, modifiers...)
	return lruo
}

func (lruo *LoginRecordUpdateOne) sqlSave(ctx context.Context) (_node *LoginRecord, err error) {
	if err := lruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(loginrecord.Table, loginrecord.Columns, sqlgraph.NewFieldSpec(loginrecord.FieldID, field.TypeInt64))
	id, ok := lruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent_work: missing "LoginRecord.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loginrecord.FieldID)
		for _, f := range fields {
			if !loginrecord.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent_work: invalid field %q for query", f)}
			}
			if f != loginrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lruo.mutation.CreatedBy(); ok {
		_spec.SetField(loginrecord.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := lruo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(loginrecord.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := lruo.mutation.UpdatedBy(); ok {
		_spec.SetField(loginrecord.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := lruo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(loginrecord.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := lruo.mutation.UpdatedAt(); ok {
		_spec.SetField(loginrecord.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lruo.mutation.DeletedAt(); ok {
		_spec.SetField(loginrecord.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := lruo.mutation.Ua(); ok {
		_spec.SetField(loginrecord.FieldUa, field.TypeString, value)
	}
	if value, ok := lruo.mutation.IP(); ok {
		_spec.SetField(loginrecord.FieldIP, field.TypeString, value)
	}
	if value, ok := lruo.mutation.Way(); ok {
		_spec.SetField(loginrecord.FieldWay, field.TypeString, value)
	}
	if lruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loginrecord.UserTable,
			Columns: []string{loginrecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loginrecord.UserTable,
			Columns: []string{loginrecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(lruo.modifiers...)
	_node = &LoginRecord{config: lruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loginrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lruo.mutation.done = true
	return _node, nil
}
