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
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/merchant"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/product"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/user"
	"github.com/huoayi/business-center-ent-private/pkg/enum"
)

// MerchantCreate is the builder for creating a Merchant entity.
type MerchantCreate struct {
	config
	mutation *MerchantMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (mc *MerchantCreate) SetCreatedBy(i int64) *MerchantCreate {
	mc.mutation.SetCreatedBy(i)
	return mc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableCreatedBy(i *int64) *MerchantCreate {
	if i != nil {
		mc.SetCreatedBy(*i)
	}
	return mc
}

// SetUpdatedBy sets the "updated_by" field.
func (mc *MerchantCreate) SetUpdatedBy(i int64) *MerchantCreate {
	mc.mutation.SetUpdatedBy(i)
	return mc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableUpdatedBy(i *int64) *MerchantCreate {
	if i != nil {
		mc.SetUpdatedBy(*i)
	}
	return mc
}

// SetCreatedAt sets the "created_at" field.
func (mc *MerchantCreate) SetCreatedAt(t time.Time) *MerchantCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableCreatedAt(t *time.Time) *MerchantCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MerchantCreate) SetUpdatedAt(t time.Time) *MerchantCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableUpdatedAt(t *time.Time) *MerchantCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetDeletedAt sets the "deleted_at" field.
func (mc *MerchantCreate) SetDeletedAt(t time.Time) *MerchantCreate {
	mc.mutation.SetDeletedAt(t)
	return mc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableDeletedAt(t *time.Time) *MerchantCreate {
	if t != nil {
		mc.SetDeletedAt(*t)
	}
	return mc
}

// SetMerchantName sets the "merchant_name" field.
func (mc *MerchantCreate) SetMerchantName(s string) *MerchantCreate {
	mc.mutation.SetMerchantName(s)
	return mc
}

// SetNillableMerchantName sets the "merchant_name" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableMerchantName(s *string) *MerchantCreate {
	if s != nil {
		mc.SetMerchantName(*s)
	}
	return mc
}

// SetJpgURL sets the "jpg_url" field.
func (mc *MerchantCreate) SetJpgURL(s string) *MerchantCreate {
	mc.mutation.SetJpgURL(s)
	return mc
}

// SetNillableJpgURL sets the "jpg_url" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableJpgURL(s *string) *MerchantCreate {
	if s != nil {
		mc.SetJpgURL(*s)
	}
	return mc
}

// SetComment sets the "comment" field.
func (mc *MerchantCreate) SetComment(s string) *MerchantCreate {
	mc.mutation.SetComment(s)
	return mc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableComment(s *string) *MerchantCreate {
	if s != nil {
		mc.SetComment(*s)
	}
	return mc
}

// SetAmount sets the "amount" field.
func (mc *MerchantCreate) SetAmount(i int) *MerchantCreate {
	mc.mutation.SetAmount(i)
	return mc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableAmount(i *int) *MerchantCreate {
	if i != nil {
		mc.SetAmount(*i)
	}
	return mc
}

// SetUserID sets the "user_id" field.
func (mc *MerchantCreate) SetUserID(i int64) *MerchantCreate {
	mc.mutation.SetUserID(i)
	return mc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableUserID(i *int64) *MerchantCreate {
	if i != nil {
		mc.SetUserID(*i)
	}
	return mc
}

// SetProvence sets the "provence" field.
func (mc *MerchantCreate) SetProvence(e enum.Provence) *MerchantCreate {
	mc.mutation.SetProvence(e)
	return mc
}

// SetNillableProvence sets the "provence" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableProvence(e *enum.Provence) *MerchantCreate {
	if e != nil {
		mc.SetProvence(*e)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MerchantCreate) SetID(i int64) *MerchantCreate {
	mc.mutation.SetID(i)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MerchantCreate) SetNillableID(i *int64) *MerchantCreate {
	if i != nil {
		mc.SetID(*i)
	}
	return mc
}

// SetUser sets the "user" edge to the User entity.
func (mc *MerchantCreate) SetUser(u *User) *MerchantCreate {
	return mc.SetUserID(u.ID)
}

// SetProductsID sets the "products" edge to the Product entity by ID.
func (mc *MerchantCreate) SetProductsID(id int64) *MerchantCreate {
	mc.mutation.SetProductsID(id)
	return mc
}

// SetNillableProductsID sets the "products" edge to the Product entity by ID if the given value is not nil.
func (mc *MerchantCreate) SetNillableProductsID(id *int64) *MerchantCreate {
	if id != nil {
		mc = mc.SetProductsID(*id)
	}
	return mc
}

// SetProducts sets the "products" edge to the Product entity.
func (mc *MerchantCreate) SetProducts(p *Product) *MerchantCreate {
	return mc.SetProductsID(p.ID)
}

// Mutation returns the MerchantMutation object of the builder.
func (mc *MerchantCreate) Mutation() *MerchantMutation {
	return mc.mutation
}

// Save creates the Merchant in the database.
func (mc *MerchantCreate) Save(ctx context.Context) (*Merchant, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MerchantCreate) SaveX(ctx context.Context) *Merchant {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MerchantCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MerchantCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MerchantCreate) defaults() {
	if _, ok := mc.mutation.CreatedBy(); !ok {
		v := merchant.DefaultCreatedBy
		mc.mutation.SetCreatedBy(v)
	}
	if _, ok := mc.mutation.UpdatedBy(); !ok {
		v := merchant.DefaultUpdatedBy
		mc.mutation.SetUpdatedBy(v)
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := merchant.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := merchant.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.DeletedAt(); !ok {
		v := merchant.DefaultDeletedAt
		mc.mutation.SetDeletedAt(v)
	}
	if _, ok := mc.mutation.MerchantName(); !ok {
		v := merchant.DefaultMerchantName
		mc.mutation.SetMerchantName(v)
	}
	if _, ok := mc.mutation.JpgURL(); !ok {
		v := merchant.DefaultJpgURL
		mc.mutation.SetJpgURL(v)
	}
	if _, ok := mc.mutation.Comment(); !ok {
		v := merchant.DefaultComment
		mc.mutation.SetComment(v)
	}
	if _, ok := mc.mutation.Amount(); !ok {
		v := merchant.DefaultAmount
		mc.mutation.SetAmount(v)
	}
	if _, ok := mc.mutation.UserID(); !ok {
		v := merchant.DefaultUserID
		mc.mutation.SetUserID(v)
	}
	if _, ok := mc.mutation.Provence(); !ok {
		v := merchant.DefaultProvence
		mc.mutation.SetProvence(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := merchant.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MerchantCreate) check() error {
	if _, ok := mc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent_work: missing required field "Merchant.created_by"`)}
	}
	if _, ok := mc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent_work: missing required field "Merchant.updated_by"`)}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent_work: missing required field "Merchant.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent_work: missing required field "Merchant.updated_at"`)}
	}
	if _, ok := mc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent_work: missing required field "Merchant.deleted_at"`)}
	}
	if _, ok := mc.mutation.MerchantName(); !ok {
		return &ValidationError{Name: "merchant_name", err: errors.New(`ent_work: missing required field "Merchant.merchant_name"`)}
	}
	if _, ok := mc.mutation.JpgURL(); !ok {
		return &ValidationError{Name: "jpg_url", err: errors.New(`ent_work: missing required field "Merchant.jpg_url"`)}
	}
	if _, ok := mc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent_work: missing required field "Merchant.comment"`)}
	}
	if _, ok := mc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent_work: missing required field "Merchant.amount"`)}
	}
	if _, ok := mc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent_work: missing required field "Merchant.user_id"`)}
	}
	if _, ok := mc.mutation.Provence(); !ok {
		return &ValidationError{Name: "provence", err: errors.New(`ent_work: missing required field "Merchant.provence"`)}
	}
	if v, ok := mc.mutation.Provence(); ok {
		if err := merchant.ProvenceValidator(v); err != nil {
			return &ValidationError{Name: "provence", err: fmt.Errorf(`ent_work: validator failed for field "Merchant.provence": %w`, err)}
		}
	}
	if _, ok := mc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent_work: missing required edge "Merchant.user"`)}
	}
	return nil
}

func (mc *MerchantCreate) sqlSave(ctx context.Context) (*Merchant, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MerchantCreate) createSpec() (*Merchant, *sqlgraph.CreateSpec) {
	var (
		_node = &Merchant{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(merchant.Table, sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedBy(); ok {
		_spec.SetField(merchant.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := mc.mutation.UpdatedBy(); ok {
		_spec.SetField(merchant.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(merchant.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(merchant.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.SetField(merchant.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := mc.mutation.MerchantName(); ok {
		_spec.SetField(merchant.FieldMerchantName, field.TypeString, value)
		_node.MerchantName = value
	}
	if value, ok := mc.mutation.JpgURL(); ok {
		_spec.SetField(merchant.FieldJpgURL, field.TypeString, value)
		_node.JpgURL = value
	}
	if value, ok := mc.mutation.Comment(); ok {
		_spec.SetField(merchant.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if value, ok := mc.mutation.Amount(); ok {
		_spec.SetField(merchant.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if value, ok := mc.mutation.Provence(); ok {
		_spec.SetField(merchant.FieldProvence, field.TypeEnum, value)
		_node.Provence = value
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   merchant.UserTable,
			Columns: []string{merchant.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   merchant.ProductsTable,
			Columns: []string{merchant.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Merchant.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MerchantUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (mc *MerchantCreate) OnConflict(opts ...sql.ConflictOption) *MerchantUpsertOne {
	mc.conflict = opts
	return &MerchantUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Merchant.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MerchantCreate) OnConflictColumns(columns ...string) *MerchantUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MerchantUpsertOne{
		create: mc,
	}
}

type (
	// MerchantUpsertOne is the builder for "upsert"-ing
	//  one Merchant node.
	MerchantUpsertOne struct {
		create *MerchantCreate
	}

	// MerchantUpsert is the "OnConflict" setter.
	MerchantUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *MerchantUpsert) SetCreatedBy(v int64) *MerchantUpsert {
	u.Set(merchant.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *MerchantUpsert) UpdateCreatedBy() *MerchantUpsert {
	u.SetExcluded(merchant.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *MerchantUpsert) AddCreatedBy(v int64) *MerchantUpsert {
	u.Add(merchant.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MerchantUpsert) SetUpdatedBy(v int64) *MerchantUpsert {
	u.Set(merchant.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MerchantUpsert) UpdateUpdatedBy() *MerchantUpsert {
	u.SetExcluded(merchant.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MerchantUpsert) AddUpdatedBy(v int64) *MerchantUpsert {
	u.Add(merchant.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MerchantUpsert) SetUpdatedAt(v time.Time) *MerchantUpsert {
	u.Set(merchant.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MerchantUpsert) UpdateUpdatedAt() *MerchantUpsert {
	u.SetExcluded(merchant.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MerchantUpsert) SetDeletedAt(v time.Time) *MerchantUpsert {
	u.Set(merchant.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MerchantUpsert) UpdateDeletedAt() *MerchantUpsert {
	u.SetExcluded(merchant.FieldDeletedAt)
	return u
}

// SetMerchantName sets the "merchant_name" field.
func (u *MerchantUpsert) SetMerchantName(v string) *MerchantUpsert {
	u.Set(merchant.FieldMerchantName, v)
	return u
}

// UpdateMerchantName sets the "merchant_name" field to the value that was provided on create.
func (u *MerchantUpsert) UpdateMerchantName() *MerchantUpsert {
	u.SetExcluded(merchant.FieldMerchantName)
	return u
}

// SetJpgURL sets the "jpg_url" field.
func (u *MerchantUpsert) SetJpgURL(v string) *MerchantUpsert {
	u.Set(merchant.FieldJpgURL, v)
	return u
}

// UpdateJpgURL sets the "jpg_url" field to the value that was provided on create.
func (u *MerchantUpsert) UpdateJpgURL() *MerchantUpsert {
	u.SetExcluded(merchant.FieldJpgURL)
	return u
}

// SetComment sets the "comment" field.
func (u *MerchantUpsert) SetComment(v string) *MerchantUpsert {
	u.Set(merchant.FieldComment, v)
	return u
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *MerchantUpsert) UpdateComment() *MerchantUpsert {
	u.SetExcluded(merchant.FieldComment)
	return u
}

// SetAmount sets the "amount" field.
func (u *MerchantUpsert) SetAmount(v int) *MerchantUpsert {
	u.Set(merchant.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *MerchantUpsert) UpdateAmount() *MerchantUpsert {
	u.SetExcluded(merchant.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *MerchantUpsert) AddAmount(v int) *MerchantUpsert {
	u.Add(merchant.FieldAmount, v)
	return u
}

// SetUserID sets the "user_id" field.
func (u *MerchantUpsert) SetUserID(v int64) *MerchantUpsert {
	u.Set(merchant.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *MerchantUpsert) UpdateUserID() *MerchantUpsert {
	u.SetExcluded(merchant.FieldUserID)
	return u
}

// SetProvence sets the "provence" field.
func (u *MerchantUpsert) SetProvence(v enum.Provence) *MerchantUpsert {
	u.Set(merchant.FieldProvence, v)
	return u
}

// UpdateProvence sets the "provence" field to the value that was provided on create.
func (u *MerchantUpsert) UpdateProvence() *MerchantUpsert {
	u.SetExcluded(merchant.FieldProvence)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Merchant.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(merchant.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MerchantUpsertOne) UpdateNewValues() *MerchantUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(merchant.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(merchant.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Merchant.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MerchantUpsertOne) Ignore() *MerchantUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MerchantUpsertOne) DoNothing() *MerchantUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MerchantCreate.OnConflict
// documentation for more info.
func (u *MerchantUpsertOne) Update(set func(*MerchantUpsert)) *MerchantUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MerchantUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *MerchantUpsertOne) SetCreatedBy(v int64) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *MerchantUpsertOne) AddCreatedBy(v int64) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *MerchantUpsertOne) UpdateCreatedBy() *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MerchantUpsertOne) SetUpdatedBy(v int64) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MerchantUpsertOne) AddUpdatedBy(v int64) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MerchantUpsertOne) UpdateUpdatedBy() *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MerchantUpsertOne) SetUpdatedAt(v time.Time) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MerchantUpsertOne) UpdateUpdatedAt() *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MerchantUpsertOne) SetDeletedAt(v time.Time) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MerchantUpsertOne) UpdateDeletedAt() *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetMerchantName sets the "merchant_name" field.
func (u *MerchantUpsertOne) SetMerchantName(v string) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.SetMerchantName(v)
	})
}

// UpdateMerchantName sets the "merchant_name" field to the value that was provided on create.
func (u *MerchantUpsertOne) UpdateMerchantName() *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateMerchantName()
	})
}

// SetJpgURL sets the "jpg_url" field.
func (u *MerchantUpsertOne) SetJpgURL(v string) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.SetJpgURL(v)
	})
}

// UpdateJpgURL sets the "jpg_url" field to the value that was provided on create.
func (u *MerchantUpsertOne) UpdateJpgURL() *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateJpgURL()
	})
}

// SetComment sets the "comment" field.
func (u *MerchantUpsertOne) SetComment(v string) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.SetComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *MerchantUpsertOne) UpdateComment() *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateComment()
	})
}

// SetAmount sets the "amount" field.
func (u *MerchantUpsertOne) SetAmount(v int) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *MerchantUpsertOne) AddAmount(v int) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *MerchantUpsertOne) UpdateAmount() *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateAmount()
	})
}

// SetUserID sets the "user_id" field.
func (u *MerchantUpsertOne) SetUserID(v int64) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *MerchantUpsertOne) UpdateUserID() *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateUserID()
	})
}

// SetProvence sets the "provence" field.
func (u *MerchantUpsertOne) SetProvence(v enum.Provence) *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.SetProvence(v)
	})
}

// UpdateProvence sets the "provence" field to the value that was provided on create.
func (u *MerchantUpsertOne) UpdateProvence() *MerchantUpsertOne {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateProvence()
	})
}

// Exec executes the query.
func (u *MerchantUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent_work: missing options for MerchantCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MerchantUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MerchantUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MerchantUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MerchantCreateBulk is the builder for creating many Merchant entities in bulk.
type MerchantCreateBulk struct {
	config
	err      error
	builders []*MerchantCreate
	conflict []sql.ConflictOption
}

// Save creates the Merchant entities in the database.
func (mcb *MerchantCreateBulk) Save(ctx context.Context) ([]*Merchant, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Merchant, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MerchantMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MerchantCreateBulk) SaveX(ctx context.Context) []*Merchant {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MerchantCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MerchantCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Merchant.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MerchantUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (mcb *MerchantCreateBulk) OnConflict(opts ...sql.ConflictOption) *MerchantUpsertBulk {
	mcb.conflict = opts
	return &MerchantUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Merchant.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MerchantCreateBulk) OnConflictColumns(columns ...string) *MerchantUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MerchantUpsertBulk{
		create: mcb,
	}
}

// MerchantUpsertBulk is the builder for "upsert"-ing
// a bulk of Merchant nodes.
type MerchantUpsertBulk struct {
	create *MerchantCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Merchant.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(merchant.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MerchantUpsertBulk) UpdateNewValues() *MerchantUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(merchant.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(merchant.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Merchant.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MerchantUpsertBulk) Ignore() *MerchantUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MerchantUpsertBulk) DoNothing() *MerchantUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MerchantCreateBulk.OnConflict
// documentation for more info.
func (u *MerchantUpsertBulk) Update(set func(*MerchantUpsert)) *MerchantUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MerchantUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *MerchantUpsertBulk) SetCreatedBy(v int64) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *MerchantUpsertBulk) AddCreatedBy(v int64) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *MerchantUpsertBulk) UpdateCreatedBy() *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MerchantUpsertBulk) SetUpdatedBy(v int64) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MerchantUpsertBulk) AddUpdatedBy(v int64) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MerchantUpsertBulk) UpdateUpdatedBy() *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MerchantUpsertBulk) SetUpdatedAt(v time.Time) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MerchantUpsertBulk) UpdateUpdatedAt() *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MerchantUpsertBulk) SetDeletedAt(v time.Time) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MerchantUpsertBulk) UpdateDeletedAt() *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetMerchantName sets the "merchant_name" field.
func (u *MerchantUpsertBulk) SetMerchantName(v string) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.SetMerchantName(v)
	})
}

// UpdateMerchantName sets the "merchant_name" field to the value that was provided on create.
func (u *MerchantUpsertBulk) UpdateMerchantName() *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateMerchantName()
	})
}

// SetJpgURL sets the "jpg_url" field.
func (u *MerchantUpsertBulk) SetJpgURL(v string) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.SetJpgURL(v)
	})
}

// UpdateJpgURL sets the "jpg_url" field to the value that was provided on create.
func (u *MerchantUpsertBulk) UpdateJpgURL() *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateJpgURL()
	})
}

// SetComment sets the "comment" field.
func (u *MerchantUpsertBulk) SetComment(v string) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.SetComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *MerchantUpsertBulk) UpdateComment() *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateComment()
	})
}

// SetAmount sets the "amount" field.
func (u *MerchantUpsertBulk) SetAmount(v int) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *MerchantUpsertBulk) AddAmount(v int) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *MerchantUpsertBulk) UpdateAmount() *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateAmount()
	})
}

// SetUserID sets the "user_id" field.
func (u *MerchantUpsertBulk) SetUserID(v int64) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *MerchantUpsertBulk) UpdateUserID() *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateUserID()
	})
}

// SetProvence sets the "provence" field.
func (u *MerchantUpsertBulk) SetProvence(v enum.Provence) *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.SetProvence(v)
	})
}

// UpdateProvence sets the "provence" field to the value that was provided on create.
func (u *MerchantUpsertBulk) UpdateProvence() *MerchantUpsertBulk {
	return u.Update(func(s *MerchantUpsert) {
		s.UpdateProvence()
	})
}

// Exec executes the query.
func (u *MerchantUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent_work: OnConflict was set for builder %d. Set it on the MerchantCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent_work: missing options for MerchantCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MerchantUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
