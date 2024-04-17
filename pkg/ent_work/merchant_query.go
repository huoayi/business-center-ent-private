// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/merchant"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/predicate"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/product"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/user"
)

// MerchantQuery is the builder for querying Merchant entities.
type MerchantQuery struct {
	config
	ctx          *QueryContext
	order        []merchant.OrderOption
	inters       []Interceptor
	predicates   []predicate.Merchant
	withUser     *UserQuery
	withProducts *ProductQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MerchantQuery builder.
func (mq *MerchantQuery) Where(ps ...predicate.Merchant) *MerchantQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MerchantQuery) Limit(limit int) *MerchantQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MerchantQuery) Offset(offset int) *MerchantQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MerchantQuery) Unique(unique bool) *MerchantQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MerchantQuery) Order(o ...merchant.OrderOption) *MerchantQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryUser chains the current query on the "user" edge.
func (mq *MerchantQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(merchant.Table, merchant.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, merchant.UserTable, merchant.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProducts chains the current query on the "products" edge.
func (mq *MerchantQuery) QueryProducts() *ProductQuery {
	query := (&ProductClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(merchant.Table, merchant.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, merchant.ProductsTable, merchant.ProductsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Merchant entity from the query.
// Returns a *NotFoundError when no Merchant was found.
func (mq *MerchantQuery) First(ctx context.Context) (*Merchant, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{merchant.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MerchantQuery) FirstX(ctx context.Context) *Merchant {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Merchant ID from the query.
// Returns a *NotFoundError when no Merchant ID was found.
func (mq *MerchantQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{merchant.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MerchantQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Merchant entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Merchant entity is found.
// Returns a *NotFoundError when no Merchant entities are found.
func (mq *MerchantQuery) Only(ctx context.Context) (*Merchant, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{merchant.Label}
	default:
		return nil, &NotSingularError{merchant.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MerchantQuery) OnlyX(ctx context.Context) *Merchant {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Merchant ID in the query.
// Returns a *NotSingularError when more than one Merchant ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MerchantQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{merchant.Label}
	default:
		err = &NotSingularError{merchant.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MerchantQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Merchants.
func (mq *MerchantQuery) All(ctx context.Context) ([]*Merchant, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Merchant, *MerchantQuery]()
	return withInterceptors[[]*Merchant](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MerchantQuery) AllX(ctx context.Context) []*Merchant {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Merchant IDs.
func (mq *MerchantQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(merchant.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MerchantQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MerchantQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MerchantQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MerchantQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MerchantQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent_work: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MerchantQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MerchantQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MerchantQuery) Clone() *MerchantQuery {
	if mq == nil {
		return nil
	}
	return &MerchantQuery{
		config:       mq.config,
		ctx:          mq.ctx.Clone(),
		order:        append([]merchant.OrderOption{}, mq.order...),
		inters:       append([]Interceptor{}, mq.inters...),
		predicates:   append([]predicate.Merchant{}, mq.predicates...),
		withUser:     mq.withUser.Clone(),
		withProducts: mq.withProducts.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MerchantQuery) WithUser(opts ...func(*UserQuery)) *MerchantQuery {
	query := (&UserClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withUser = query
	return mq
}

// WithProducts tells the query-builder to eager-load the nodes that are connected to
// the "products" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MerchantQuery) WithProducts(opts ...func(*ProductQuery)) *MerchantQuery {
	query := (&ProductClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withProducts = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Merchant.Query().
//		GroupBy(merchant.FieldCreatedBy).
//		Aggregate(ent_work.Count()).
//		Scan(ctx, &v)
func (mq *MerchantQuery) GroupBy(field string, fields ...string) *MerchantGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MerchantGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = merchant.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//	}
//
//	client.Merchant.Query().
//		Select(merchant.FieldCreatedBy).
//		Scan(ctx, &v)
func (mq *MerchantQuery) Select(fields ...string) *MerchantSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MerchantSelect{MerchantQuery: mq}
	sbuild.label = merchant.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MerchantSelect configured with the given aggregations.
func (mq *MerchantQuery) Aggregate(fns ...AggregateFunc) *MerchantSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MerchantQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent_work: uninitialized interceptor (forgotten import ent_work/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !merchant.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent_work: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MerchantQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Merchant, error) {
	var (
		nodes       = []*Merchant{}
		_spec       = mq.querySpec()
		loadedTypes = [2]bool{
			mq.withUser != nil,
			mq.withProducts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Merchant).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Merchant{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withUser; query != nil {
		if err := mq.loadUser(ctx, query, nodes, nil,
			func(n *Merchant, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withProducts; query != nil {
		if err := mq.loadProducts(ctx, query, nodes, nil,
			func(n *Merchant, e *Product) { n.Edges.Products = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MerchantQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Merchant, init func(*Merchant), assign func(*Merchant, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Merchant)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MerchantQuery) loadProducts(ctx context.Context, query *ProductQuery, nodes []*Merchant, init func(*Merchant), assign func(*Merchant, *Product)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Merchant)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(product.FieldBusinessID)
	}
	query.Where(predicate.Product(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(merchant.ProductsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BusinessID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "business_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MerchantQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MerchantQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(merchant.Table, merchant.Columns, sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeInt64))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchant.FieldID)
		for i := range fields {
			if fields[i] != merchant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mq.withUser != nil {
			_spec.Node.AddColumnOnce(merchant.FieldUserID)
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MerchantQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(merchant.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = merchant.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range mq.modifiers {
		m(selector)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mq *MerchantQuery) Modify(modifiers ...func(s *sql.Selector)) *MerchantSelect {
	mq.modifiers = append(mq.modifiers, modifiers...)
	return mq.Select()
}

// MerchantGroupBy is the group-by builder for Merchant entities.
type MerchantGroupBy struct {
	selector
	build *MerchantQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MerchantGroupBy) Aggregate(fns ...AggregateFunc) *MerchantGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MerchantGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MerchantQuery, *MerchantGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MerchantGroupBy) sqlScan(ctx context.Context, root *MerchantQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MerchantSelect is the builder for selecting fields of Merchant entities.
type MerchantSelect struct {
	*MerchantQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MerchantSelect) Aggregate(fns ...AggregateFunc) *MerchantSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MerchantSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MerchantQuery, *MerchantSelect](ctx, ms.MerchantQuery, ms, ms.inters, v)
}

func (ms *MerchantSelect) sqlScan(ctx context.Context, root *MerchantQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ms *MerchantSelect) Modify(modifiers ...func(s *sql.Selector)) *MerchantSelect {
	ms.modifiers = append(ms.modifiers, modifiers...)
	return ms
}
