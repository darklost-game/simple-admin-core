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
	"github.com/suyuan32/simple-admin-core/rpc/ent/loglogin"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

// LogLoginQuery is the builder for querying LogLogin entities.
type LogLoginQuery struct {
	config
	ctx        *QueryContext
	order      []loglogin.OrderOption
	inters     []Interceptor
	predicates []predicate.LogLogin
	withUser   *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LogLoginQuery builder.
func (llq *LogLoginQuery) Where(ps ...predicate.LogLogin) *LogLoginQuery {
	llq.predicates = append(llq.predicates, ps...)
	return llq
}

// Limit the number of records to be returned by this query.
func (llq *LogLoginQuery) Limit(limit int) *LogLoginQuery {
	llq.ctx.Limit = &limit
	return llq
}

// Offset to start from.
func (llq *LogLoginQuery) Offset(offset int) *LogLoginQuery {
	llq.ctx.Offset = &offset
	return llq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (llq *LogLoginQuery) Unique(unique bool) *LogLoginQuery {
	llq.ctx.Unique = &unique
	return llq
}

// Order specifies how the records should be ordered.
func (llq *LogLoginQuery) Order(o ...loglogin.OrderOption) *LogLoginQuery {
	llq.order = append(llq.order, o...)
	return llq
}

// QueryUser chains the current query on the "user" edge.
func (llq *LogLoginQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: llq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := llq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := llq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(loglogin.Table, loglogin.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, loglogin.UserTable, loglogin.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(llq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LogLogin entity from the query.
// Returns a *NotFoundError when no LogLogin was found.
func (llq *LogLoginQuery) First(ctx context.Context) (*LogLogin, error) {
	nodes, err := llq.Limit(1).All(setContextOp(ctx, llq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loglogin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (llq *LogLoginQuery) FirstX(ctx context.Context) *LogLogin {
	node, err := llq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LogLogin ID from the query.
// Returns a *NotFoundError when no LogLogin ID was found.
func (llq *LogLoginQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = llq.Limit(1).IDs(setContextOp(ctx, llq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loglogin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (llq *LogLoginQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := llq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LogLogin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LogLogin entity is found.
// Returns a *NotFoundError when no LogLogin entities are found.
func (llq *LogLoginQuery) Only(ctx context.Context) (*LogLogin, error) {
	nodes, err := llq.Limit(2).All(setContextOp(ctx, llq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loglogin.Label}
	default:
		return nil, &NotSingularError{loglogin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (llq *LogLoginQuery) OnlyX(ctx context.Context) *LogLogin {
	node, err := llq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LogLogin ID in the query.
// Returns a *NotSingularError when more than one LogLogin ID is found.
// Returns a *NotFoundError when no entities are found.
func (llq *LogLoginQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = llq.Limit(2).IDs(setContextOp(ctx, llq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loglogin.Label}
	default:
		err = &NotSingularError{loglogin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (llq *LogLoginQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := llq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LogLogins.
func (llq *LogLoginQuery) All(ctx context.Context) ([]*LogLogin, error) {
	ctx = setContextOp(ctx, llq.ctx, "All")
	if err := llq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LogLogin, *LogLoginQuery]()
	return withInterceptors[[]*LogLogin](ctx, llq, qr, llq.inters)
}

// AllX is like All, but panics if an error occurs.
func (llq *LogLoginQuery) AllX(ctx context.Context) []*LogLogin {
	nodes, err := llq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LogLogin IDs.
func (llq *LogLoginQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if llq.ctx.Unique == nil && llq.path != nil {
		llq.Unique(true)
	}
	ctx = setContextOp(ctx, llq.ctx, "IDs")
	if err = llq.Select(loglogin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (llq *LogLoginQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := llq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (llq *LogLoginQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, llq.ctx, "Count")
	if err := llq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, llq, querierCount[*LogLoginQuery](), llq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (llq *LogLoginQuery) CountX(ctx context.Context) int {
	count, err := llq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (llq *LogLoginQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, llq.ctx, "Exist")
	switch _, err := llq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (llq *LogLoginQuery) ExistX(ctx context.Context) bool {
	exist, err := llq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LogLoginQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (llq *LogLoginQuery) Clone() *LogLoginQuery {
	if llq == nil {
		return nil
	}
	return &LogLoginQuery{
		config:     llq.config,
		ctx:        llq.ctx.Clone(),
		order:      append([]loglogin.OrderOption{}, llq.order...),
		inters:     append([]Interceptor{}, llq.inters...),
		predicates: append([]predicate.LogLogin{}, llq.predicates...),
		withUser:   llq.withUser.Clone(),
		// clone intermediate query.
		sql:  llq.sql.Clone(),
		path: llq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (llq *LogLoginQuery) WithUser(opts ...func(*UserQuery)) *LogLoginQuery {
	query := (&UserClient{config: llq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	llq.withUser = query
	return llq
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
//	client.LogLogin.Query().
//		GroupBy(loglogin.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (llq *LogLoginQuery) GroupBy(field string, fields ...string) *LogLoginGroupBy {
	llq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LogLoginGroupBy{build: llq}
	grbuild.flds = &llq.ctx.Fields
	grbuild.label = loglogin.Label
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
//	client.LogLogin.Query().
//		Select(loglogin.FieldCreatedAt).
//		Scan(ctx, &v)
func (llq *LogLoginQuery) Select(fields ...string) *LogLoginSelect {
	llq.ctx.Fields = append(llq.ctx.Fields, fields...)
	sbuild := &LogLoginSelect{LogLoginQuery: llq}
	sbuild.label = loglogin.Label
	sbuild.flds, sbuild.scan = &llq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LogLoginSelect configured with the given aggregations.
func (llq *LogLoginQuery) Aggregate(fns ...AggregateFunc) *LogLoginSelect {
	return llq.Select().Aggregate(fns...)
}

func (llq *LogLoginQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range llq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, llq); err != nil {
				return err
			}
		}
	}
	for _, f := range llq.ctx.Fields {
		if !loglogin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if llq.path != nil {
		prev, err := llq.path(ctx)
		if err != nil {
			return err
		}
		llq.sql = prev
	}
	return nil
}

func (llq *LogLoginQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LogLogin, error) {
	var (
		nodes       = []*LogLogin{}
		_spec       = llq.querySpec()
		loadedTypes = [1]bool{
			llq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LogLogin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LogLogin{config: llq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(llq.modifiers) > 0 {
		_spec.Modifiers = llq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, llq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := llq.withUser; query != nil {
		if err := llq.loadUser(ctx, query, nodes, nil,
			func(n *LogLogin, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (llq *LogLoginQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*LogLogin, init func(*LogLogin), assign func(*LogLogin, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*LogLogin)
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

func (llq *LogLoginQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := llq.querySpec()
	if len(llq.modifiers) > 0 {
		_spec.Modifiers = llq.modifiers
	}
	_spec.Node.Columns = llq.ctx.Fields
	if len(llq.ctx.Fields) > 0 {
		_spec.Unique = llq.ctx.Unique != nil && *llq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, llq.driver, _spec)
}

func (llq *LogLoginQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(loglogin.Table, loglogin.Columns, sqlgraph.NewFieldSpec(loglogin.FieldID, field.TypeUint64))
	_spec.From = llq.sql
	if unique := llq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if llq.path != nil {
		_spec.Unique = true
	}
	if fields := llq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loglogin.FieldID)
		for i := range fields {
			if fields[i] != loglogin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if llq.withUser != nil {
			_spec.Node.AddColumnOnce(loglogin.FieldUUID)
		}
	}
	if ps := llq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := llq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := llq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := llq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (llq *LogLoginQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(llq.driver.Dialect())
	t1 := builder.Table(loglogin.Table)
	columns := llq.ctx.Fields
	if len(columns) == 0 {
		columns = loglogin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if llq.sql != nil {
		selector = llq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if llq.ctx.Unique != nil && *llq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range llq.modifiers {
		m(selector)
	}
	for _, p := range llq.predicates {
		p(selector)
	}
	for _, p := range llq.order {
		p(selector)
	}
	if offset := llq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := llq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (llq *LogLoginQuery) Modify(modifiers ...func(s *sql.Selector)) *LogLoginSelect {
	llq.modifiers = append(llq.modifiers, modifiers...)
	return llq.Select()
}

// LogLoginGroupBy is the group-by builder for LogLogin entities.
type LogLoginGroupBy struct {
	selector
	build *LogLoginQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (llgb *LogLoginGroupBy) Aggregate(fns ...AggregateFunc) *LogLoginGroupBy {
	llgb.fns = append(llgb.fns, fns...)
	return llgb
}

// Scan applies the selector query and scans the result into the given value.
func (llgb *LogLoginGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, llgb.build.ctx, "GroupBy")
	if err := llgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogLoginQuery, *LogLoginGroupBy](ctx, llgb.build, llgb, llgb.build.inters, v)
}

func (llgb *LogLoginGroupBy) sqlScan(ctx context.Context, root *LogLoginQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(llgb.fns))
	for _, fn := range llgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*llgb.flds)+len(llgb.fns))
		for _, f := range *llgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*llgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := llgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LogLoginSelect is the builder for selecting fields of LogLogin entities.
type LogLoginSelect struct {
	*LogLoginQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lls *LogLoginSelect) Aggregate(fns ...AggregateFunc) *LogLoginSelect {
	lls.fns = append(lls.fns, fns...)
	return lls
}

// Scan applies the selector query and scans the result into the given value.
func (lls *LogLoginSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lls.ctx, "Select")
	if err := lls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogLoginQuery, *LogLoginSelect](ctx, lls.LogLoginQuery, lls, lls.inters, v)
}

func (lls *LogLoginSelect) sqlScan(ctx context.Context, root *LogLoginQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lls.fns))
	for _, fn := range lls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lls *LogLoginSelect) Modify(modifiers ...func(s *sql.Selector)) *LogLoginSelect {
	lls.modifiers = append(lls.modifiers, modifiers...)
	return lls
}
