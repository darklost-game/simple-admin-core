// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/logoperation"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

// LogOperationQuery is the builder for querying LogOperation entities.
type LogOperationQuery struct {
	config
	ctx        *QueryContext
	order      []logoperation.OrderOption
	inters     []Interceptor
	predicates []predicate.LogOperation
	withUser   *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LogOperationQuery builder.
func (loq *LogOperationQuery) Where(ps ...predicate.LogOperation) *LogOperationQuery {
	loq.predicates = append(loq.predicates, ps...)
	return loq
}

// Limit the number of records to be returned by this query.
func (loq *LogOperationQuery) Limit(limit int) *LogOperationQuery {
	loq.ctx.Limit = &limit
	return loq
}

// Offset to start from.
func (loq *LogOperationQuery) Offset(offset int) *LogOperationQuery {
	loq.ctx.Offset = &offset
	return loq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (loq *LogOperationQuery) Unique(unique bool) *LogOperationQuery {
	loq.ctx.Unique = &unique
	return loq
}

// Order specifies how the records should be ordered.
func (loq *LogOperationQuery) Order(o ...logoperation.OrderOption) *LogOperationQuery {
	loq.order = append(loq.order, o...)
	return loq
}

// QueryUser chains the current query on the "user" edge.
func (loq *LogOperationQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: loq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := loq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := loq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(logoperation.Table, logoperation.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, logoperation.UserTable, logoperation.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(loq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LogOperation entity from the query.
// Returns a *NotFoundError when no LogOperation was found.
func (loq *LogOperationQuery) First(ctx context.Context) (*LogOperation, error) {
	nodes, err := loq.Limit(1).All(setContextOp(ctx, loq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{logoperation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (loq *LogOperationQuery) FirstX(ctx context.Context) *LogOperation {
	node, err := loq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LogOperation ID from the query.
// Returns a *NotFoundError when no LogOperation ID was found.
func (loq *LogOperationQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = loq.Limit(1).IDs(setContextOp(ctx, loq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{logoperation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (loq *LogOperationQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := loq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LogOperation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LogOperation entity is found.
// Returns a *NotFoundError when no LogOperation entities are found.
func (loq *LogOperationQuery) Only(ctx context.Context) (*LogOperation, error) {
	nodes, err := loq.Limit(2).All(setContextOp(ctx, loq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{logoperation.Label}
	default:
		return nil, &NotSingularError{logoperation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (loq *LogOperationQuery) OnlyX(ctx context.Context) *LogOperation {
	node, err := loq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LogOperation ID in the query.
// Returns a *NotSingularError when more than one LogOperation ID is found.
// Returns a *NotFoundError when no entities are found.
func (loq *LogOperationQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = loq.Limit(2).IDs(setContextOp(ctx, loq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{logoperation.Label}
	default:
		err = &NotSingularError{logoperation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (loq *LogOperationQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := loq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LogOperations.
func (loq *LogOperationQuery) All(ctx context.Context) ([]*LogOperation, error) {
	ctx = setContextOp(ctx, loq.ctx, "All")
	if err := loq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LogOperation, *LogOperationQuery]()
	return withInterceptors[[]*LogOperation](ctx, loq, qr, loq.inters)
}

// AllX is like All, but panics if an error occurs.
func (loq *LogOperationQuery) AllX(ctx context.Context) []*LogOperation {
	nodes, err := loq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LogOperation IDs.
func (loq *LogOperationQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if loq.ctx.Unique == nil && loq.path != nil {
		loq.Unique(true)
	}
	ctx = setContextOp(ctx, loq.ctx, "IDs")
	if err = loq.Select(logoperation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (loq *LogOperationQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := loq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (loq *LogOperationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, loq.ctx, "Count")
	if err := loq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, loq, querierCount[*LogOperationQuery](), loq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (loq *LogOperationQuery) CountX(ctx context.Context) int {
	count, err := loq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (loq *LogOperationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, loq.ctx, "Exist")
	switch _, err := loq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (loq *LogOperationQuery) ExistX(ctx context.Context) bool {
	exist, err := loq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LogOperationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (loq *LogOperationQuery) Clone() *LogOperationQuery {
	if loq == nil {
		return nil
	}
	return &LogOperationQuery{
		config:     loq.config,
		ctx:        loq.ctx.Clone(),
		order:      append([]logoperation.OrderOption{}, loq.order...),
		inters:     append([]Interceptor{}, loq.inters...),
		predicates: append([]predicate.LogOperation{}, loq.predicates...),
		withUser:   loq.withUser.Clone(),
		// clone intermediate query.
		sql:  loq.sql.Clone(),
		path: loq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (loq *LogOperationQuery) WithUser(opts ...func(*UserQuery)) *LogOperationQuery {
	query := (&UserClient{config: loq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	loq.withUser = query
	return loq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LogOperation.Query().
//		GroupBy(logoperation.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (loq *LogOperationQuery) GroupBy(field string, fields ...string) *LogOperationGroupBy {
	loq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LogOperationGroupBy{build: loq}
	grbuild.flds = &loq.ctx.Fields
	grbuild.label = logoperation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.LogOperation.Query().
//		Select(logoperation.FieldCreatedAt).
//		Scan(ctx, &v)
func (loq *LogOperationQuery) Select(fields ...string) *LogOperationSelect {
	loq.ctx.Fields = append(loq.ctx.Fields, fields...)
	sbuild := &LogOperationSelect{LogOperationQuery: loq}
	sbuild.label = logoperation.Label
	sbuild.flds, sbuild.scan = &loq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LogOperationSelect configured with the given aggregations.
func (loq *LogOperationQuery) Aggregate(fns ...AggregateFunc) *LogOperationSelect {
	return loq.Select().Aggregate(fns...)
}

func (loq *LogOperationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range loq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, loq); err != nil {
				return err
			}
		}
	}
	for _, f := range loq.ctx.Fields {
		if !logoperation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if loq.path != nil {
		prev, err := loq.path(ctx)
		if err != nil {
			return err
		}
		loq.sql = prev
	}
	return nil
}

func (loq *LogOperationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LogOperation, error) {
	var (
		nodes       = []*LogOperation{}
		_spec       = loq.querySpec()
		loadedTypes = [1]bool{
			loq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LogOperation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LogOperation{config: loq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(loq.modifiers) > 0 {
		_spec.Modifiers = loq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, loq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := loq.withUser; query != nil {
		if err := loq.loadUser(ctx, query, nodes, nil,
			func(n *LogOperation, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (loq *LogOperationQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*LogOperation, init func(*LogOperation), assign func(*LogOperation, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*LogOperation)
	for i := range nodes {
		fk := nodes[i].UUID
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
			return fmt.Errorf(`unexpected foreign-key "uuid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (loq *LogOperationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := loq.querySpec()
	if len(loq.modifiers) > 0 {
		_spec.Modifiers = loq.modifiers
	}
	_spec.Node.Columns = loq.ctx.Fields
	if len(loq.ctx.Fields) > 0 {
		_spec.Unique = loq.ctx.Unique != nil && *loq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, loq.driver, _spec)
}

func (loq *LogOperationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(logoperation.Table, logoperation.Columns, sqlgraph.NewFieldSpec(logoperation.FieldID, field.TypeUint64))
	_spec.From = loq.sql
	if unique := loq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if loq.path != nil {
		_spec.Unique = true
	}
	if fields := loq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logoperation.FieldID)
		for i := range fields {
			if fields[i] != logoperation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if loq.withUser != nil {
			_spec.Node.AddColumnOnce(logoperation.FieldUUID)
		}
	}
	if ps := loq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := loq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := loq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := loq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (loq *LogOperationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(loq.driver.Dialect())
	t1 := builder.Table(logoperation.Table)
	columns := loq.ctx.Fields
	if len(columns) == 0 {
		columns = logoperation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if loq.sql != nil {
		selector = loq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if loq.ctx.Unique != nil && *loq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range loq.modifiers {
		m(selector)
	}
	for _, p := range loq.predicates {
		p(selector)
	}
	for _, p := range loq.order {
		p(selector)
	}
	if offset := loq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := loq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (loq *LogOperationQuery) Modify(modifiers ...func(s *sql.Selector)) *LogOperationSelect {
	loq.modifiers = append(loq.modifiers, modifiers...)
	return loq.Select()
}

// LogOperationGroupBy is the group-by builder for LogOperation entities.
type LogOperationGroupBy struct {
	selector
	build *LogOperationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (logb *LogOperationGroupBy) Aggregate(fns ...AggregateFunc) *LogOperationGroupBy {
	logb.fns = append(logb.fns, fns...)
	return logb
}

// Scan applies the selector query and scans the result into the given value.
func (logb *LogOperationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, logb.build.ctx, "GroupBy")
	if err := logb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogOperationQuery, *LogOperationGroupBy](ctx, logb.build, logb, logb.build.inters, v)
}

func (logb *LogOperationGroupBy) sqlScan(ctx context.Context, root *LogOperationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(logb.fns))
	for _, fn := range logb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*logb.flds)+len(logb.fns))
		for _, f := range *logb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*logb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := logb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LogOperationSelect is the builder for selecting fields of LogOperation entities.
type LogOperationSelect struct {
	*LogOperationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (los *LogOperationSelect) Aggregate(fns ...AggregateFunc) *LogOperationSelect {
	los.fns = append(los.fns, fns...)
	return los
}

// Scan applies the selector query and scans the result into the given value.
func (los *LogOperationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, los.ctx, "Select")
	if err := los.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogOperationQuery, *LogOperationSelect](ctx, los.LogOperationQuery, los, los.inters, v)
}

func (los *LogOperationSelect) sqlScan(ctx context.Context, root *LogOperationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(los.fns))
	for _, fn := range los.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*los.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := los.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (los *LogOperationSelect) Modify(modifiers ...func(s *sql.Selector)) *LogOperationSelect {
	los.modifiers = append(los.modifiers, modifiers...)
	return los
}
