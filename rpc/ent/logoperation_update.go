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
	"github.com/suyuan32/simple-admin-core/rpc/ent/logoperation"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

// LogOperationUpdate is the builder for updating LogOperation entities.
type LogOperationUpdate struct {
	config
	hooks     []Hook
	mutation  *LogOperationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LogOperationUpdate builder.
func (lou *LogOperationUpdate) Where(ps ...predicate.LogOperation) *LogOperationUpdate {
	lou.mutation.Where(ps...)
	return lou
}

// SetUpdatedAt sets the "updated_at" field.
func (lou *LogOperationUpdate) SetUpdatedAt(t time.Time) *LogOperationUpdate {
	lou.mutation.SetUpdatedAt(t)
	return lou
}

// SetUUID sets the "uuid" field.
func (lou *LogOperationUpdate) SetUUID(u uuid.UUID) *LogOperationUpdate {
	lou.mutation.SetUUID(u)
	return lou
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (lou *LogOperationUpdate) SetNillableUUID(u *uuid.UUID) *LogOperationUpdate {
	if u != nil {
		lou.SetUUID(*u)
	}
	return lou
}

// SetMethod sets the "method" field.
func (lou *LogOperationUpdate) SetMethod(s string) *LogOperationUpdate {
	lou.mutation.SetMethod(s)
	return lou
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (lou *LogOperationUpdate) SetNillableMethod(s *string) *LogOperationUpdate {
	if s != nil {
		lou.SetMethod(*s)
	}
	return lou
}

// SetPath sets the "path" field.
func (lou *LogOperationUpdate) SetPath(s string) *LogOperationUpdate {
	lou.mutation.SetPath(s)
	return lou
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (lou *LogOperationUpdate) SetNillablePath(s *string) *LogOperationUpdate {
	if s != nil {
		lou.SetPath(*s)
	}
	return lou
}

// SetHeaders sets the "headers" field.
func (lou *LogOperationUpdate) SetHeaders(s string) *LogOperationUpdate {
	lou.mutation.SetHeaders(s)
	return lou
}

// SetNillableHeaders sets the "headers" field if the given value is not nil.
func (lou *LogOperationUpdate) SetNillableHeaders(s *string) *LogOperationUpdate {
	if s != nil {
		lou.SetHeaders(*s)
	}
	return lou
}

// SetBody sets the "body" field.
func (lou *LogOperationUpdate) SetBody(s string) *LogOperationUpdate {
	lou.mutation.SetBody(s)
	return lou
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (lou *LogOperationUpdate) SetNillableBody(s *string) *LogOperationUpdate {
	if s != nil {
		lou.SetBody(*s)
	}
	return lou
}

// ClearBody clears the value of the "body" field.
func (lou *LogOperationUpdate) ClearBody() *LogOperationUpdate {
	lou.mutation.ClearBody()
	return lou
}

// SetStatusCode sets the "status_code" field.
func (lou *LogOperationUpdate) SetStatusCode(i int) *LogOperationUpdate {
	lou.mutation.ResetStatusCode()
	lou.mutation.SetStatusCode(i)
	return lou
}

// SetNillableStatusCode sets the "status_code" field if the given value is not nil.
func (lou *LogOperationUpdate) SetNillableStatusCode(i *int) *LogOperationUpdate {
	if i != nil {
		lou.SetStatusCode(*i)
	}
	return lou
}

// AddStatusCode adds i to the "status_code" field.
func (lou *LogOperationUpdate) AddStatusCode(i int) *LogOperationUpdate {
	lou.mutation.AddStatusCode(i)
	return lou
}

// SetReqTime sets the "req_time" field.
func (lou *LogOperationUpdate) SetReqTime(t time.Time) *LogOperationUpdate {
	lou.mutation.SetReqTime(t)
	return lou
}

// SetNillableReqTime sets the "req_time" field if the given value is not nil.
func (lou *LogOperationUpdate) SetNillableReqTime(t *time.Time) *LogOperationUpdate {
	if t != nil {
		lou.SetReqTime(*t)
	}
	return lou
}

// SetResTime sets the "res_time" field.
func (lou *LogOperationUpdate) SetResTime(t time.Time) *LogOperationUpdate {
	lou.mutation.SetResTime(t)
	return lou
}

// SetNillableResTime sets the "res_time" field if the given value is not nil.
func (lou *LogOperationUpdate) SetNillableResTime(t *time.Time) *LogOperationUpdate {
	if t != nil {
		lou.SetResTime(*t)
	}
	return lou
}

// SetCostTime sets the "cost_time" field.
func (lou *LogOperationUpdate) SetCostTime(u uint64) *LogOperationUpdate {
	lou.mutation.ResetCostTime()
	lou.mutation.SetCostTime(u)
	return lou
}

// SetNillableCostTime sets the "cost_time" field if the given value is not nil.
func (lou *LogOperationUpdate) SetNillableCostTime(u *uint64) *LogOperationUpdate {
	if u != nil {
		lou.SetCostTime(*u)
	}
	return lou
}

// AddCostTime adds u to the "cost_time" field.
func (lou *LogOperationUpdate) AddCostTime(u int64) *LogOperationUpdate {
	lou.mutation.AddCostTime(u)
	return lou
}

// SetUserID sets the "user" edge to the User entity by ID.
func (lou *LogOperationUpdate) SetUserID(id uuid.UUID) *LogOperationUpdate {
	lou.mutation.SetUserID(id)
	return lou
}

// SetUser sets the "user" edge to the User entity.
func (lou *LogOperationUpdate) SetUser(u *User) *LogOperationUpdate {
	return lou.SetUserID(u.ID)
}

// Mutation returns the LogOperationMutation object of the builder.
func (lou *LogOperationUpdate) Mutation() *LogOperationMutation {
	return lou.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (lou *LogOperationUpdate) ClearUser() *LogOperationUpdate {
	lou.mutation.ClearUser()
	return lou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lou *LogOperationUpdate) Save(ctx context.Context) (int, error) {
	lou.defaults()
	return withHooks(ctx, lou.sqlSave, lou.mutation, lou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lou *LogOperationUpdate) SaveX(ctx context.Context) int {
	affected, err := lou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lou *LogOperationUpdate) Exec(ctx context.Context) error {
	_, err := lou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lou *LogOperationUpdate) ExecX(ctx context.Context) {
	if err := lou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lou *LogOperationUpdate) defaults() {
	if _, ok := lou.mutation.UpdatedAt(); !ok {
		v := logoperation.UpdateDefaultUpdatedAt()
		lou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lou *LogOperationUpdate) check() error {
	if v, ok := lou.mutation.Headers(); ok {
		if err := logoperation.HeadersValidator(v); err != nil {
			return &ValidationError{Name: "headers", err: fmt.Errorf(`ent: validator failed for field "LogOperation.headers": %w`, err)}
		}
	}
	if _, ok := lou.mutation.UserID(); lou.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LogOperation.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lou *LogOperationUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LogOperationUpdate {
	lou.modifiers = append(lou.modifiers, modifiers...)
	return lou
}

func (lou *LogOperationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(logoperation.Table, logoperation.Columns, sqlgraph.NewFieldSpec(logoperation.FieldID, field.TypeUint64))
	if ps := lou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lou.mutation.UpdatedAt(); ok {
		_spec.SetField(logoperation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lou.mutation.Method(); ok {
		_spec.SetField(logoperation.FieldMethod, field.TypeString, value)
	}
	if value, ok := lou.mutation.Path(); ok {
		_spec.SetField(logoperation.FieldPath, field.TypeString, value)
	}
	if value, ok := lou.mutation.Headers(); ok {
		_spec.SetField(logoperation.FieldHeaders, field.TypeString, value)
	}
	if value, ok := lou.mutation.Body(); ok {
		_spec.SetField(logoperation.FieldBody, field.TypeString, value)
	}
	if lou.mutation.BodyCleared() {
		_spec.ClearField(logoperation.FieldBody, field.TypeString)
	}
	if value, ok := lou.mutation.StatusCode(); ok {
		_spec.SetField(logoperation.FieldStatusCode, field.TypeInt, value)
	}
	if value, ok := lou.mutation.AddedStatusCode(); ok {
		_spec.AddField(logoperation.FieldStatusCode, field.TypeInt, value)
	}
	if value, ok := lou.mutation.ReqTime(); ok {
		_spec.SetField(logoperation.FieldReqTime, field.TypeTime, value)
	}
	if value, ok := lou.mutation.ResTime(); ok {
		_spec.SetField(logoperation.FieldResTime, field.TypeTime, value)
	}
	if value, ok := lou.mutation.CostTime(); ok {
		_spec.SetField(logoperation.FieldCostTime, field.TypeUint64, value)
	}
	if value, ok := lou.mutation.AddedCostTime(); ok {
		_spec.AddField(logoperation.FieldCostTime, field.TypeUint64, value)
	}
	if lou.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logoperation.UserTable,
			Columns: []string{logoperation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lou.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logoperation.UserTable,
			Columns: []string{logoperation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(lou.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, lou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logoperation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lou.mutation.done = true
	return n, nil
}

// LogOperationUpdateOne is the builder for updating a single LogOperation entity.
type LogOperationUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LogOperationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (louo *LogOperationUpdateOne) SetUpdatedAt(t time.Time) *LogOperationUpdateOne {
	louo.mutation.SetUpdatedAt(t)
	return louo
}

// SetUUID sets the "uuid" field.
func (louo *LogOperationUpdateOne) SetUUID(u uuid.UUID) *LogOperationUpdateOne {
	louo.mutation.SetUUID(u)
	return louo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (louo *LogOperationUpdateOne) SetNillableUUID(u *uuid.UUID) *LogOperationUpdateOne {
	if u != nil {
		louo.SetUUID(*u)
	}
	return louo
}

// SetMethod sets the "method" field.
func (louo *LogOperationUpdateOne) SetMethod(s string) *LogOperationUpdateOne {
	louo.mutation.SetMethod(s)
	return louo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (louo *LogOperationUpdateOne) SetNillableMethod(s *string) *LogOperationUpdateOne {
	if s != nil {
		louo.SetMethod(*s)
	}
	return louo
}

// SetPath sets the "path" field.
func (louo *LogOperationUpdateOne) SetPath(s string) *LogOperationUpdateOne {
	louo.mutation.SetPath(s)
	return louo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (louo *LogOperationUpdateOne) SetNillablePath(s *string) *LogOperationUpdateOne {
	if s != nil {
		louo.SetPath(*s)
	}
	return louo
}

// SetHeaders sets the "headers" field.
func (louo *LogOperationUpdateOne) SetHeaders(s string) *LogOperationUpdateOne {
	louo.mutation.SetHeaders(s)
	return louo
}

// SetNillableHeaders sets the "headers" field if the given value is not nil.
func (louo *LogOperationUpdateOne) SetNillableHeaders(s *string) *LogOperationUpdateOne {
	if s != nil {
		louo.SetHeaders(*s)
	}
	return louo
}

// SetBody sets the "body" field.
func (louo *LogOperationUpdateOne) SetBody(s string) *LogOperationUpdateOne {
	louo.mutation.SetBody(s)
	return louo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (louo *LogOperationUpdateOne) SetNillableBody(s *string) *LogOperationUpdateOne {
	if s != nil {
		louo.SetBody(*s)
	}
	return louo
}

// ClearBody clears the value of the "body" field.
func (louo *LogOperationUpdateOne) ClearBody() *LogOperationUpdateOne {
	louo.mutation.ClearBody()
	return louo
}

// SetStatusCode sets the "status_code" field.
func (louo *LogOperationUpdateOne) SetStatusCode(i int) *LogOperationUpdateOne {
	louo.mutation.ResetStatusCode()
	louo.mutation.SetStatusCode(i)
	return louo
}

// SetNillableStatusCode sets the "status_code" field if the given value is not nil.
func (louo *LogOperationUpdateOne) SetNillableStatusCode(i *int) *LogOperationUpdateOne {
	if i != nil {
		louo.SetStatusCode(*i)
	}
	return louo
}

// AddStatusCode adds i to the "status_code" field.
func (louo *LogOperationUpdateOne) AddStatusCode(i int) *LogOperationUpdateOne {
	louo.mutation.AddStatusCode(i)
	return louo
}

// SetReqTime sets the "req_time" field.
func (louo *LogOperationUpdateOne) SetReqTime(t time.Time) *LogOperationUpdateOne {
	louo.mutation.SetReqTime(t)
	return louo
}

// SetNillableReqTime sets the "req_time" field if the given value is not nil.
func (louo *LogOperationUpdateOne) SetNillableReqTime(t *time.Time) *LogOperationUpdateOne {
	if t != nil {
		louo.SetReqTime(*t)
	}
	return louo
}

// SetResTime sets the "res_time" field.
func (louo *LogOperationUpdateOne) SetResTime(t time.Time) *LogOperationUpdateOne {
	louo.mutation.SetResTime(t)
	return louo
}

// SetNillableResTime sets the "res_time" field if the given value is not nil.
func (louo *LogOperationUpdateOne) SetNillableResTime(t *time.Time) *LogOperationUpdateOne {
	if t != nil {
		louo.SetResTime(*t)
	}
	return louo
}

// SetCostTime sets the "cost_time" field.
func (louo *LogOperationUpdateOne) SetCostTime(u uint64) *LogOperationUpdateOne {
	louo.mutation.ResetCostTime()
	louo.mutation.SetCostTime(u)
	return louo
}

// SetNillableCostTime sets the "cost_time" field if the given value is not nil.
func (louo *LogOperationUpdateOne) SetNillableCostTime(u *uint64) *LogOperationUpdateOne {
	if u != nil {
		louo.SetCostTime(*u)
	}
	return louo
}

// AddCostTime adds u to the "cost_time" field.
func (louo *LogOperationUpdateOne) AddCostTime(u int64) *LogOperationUpdateOne {
	louo.mutation.AddCostTime(u)
	return louo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (louo *LogOperationUpdateOne) SetUserID(id uuid.UUID) *LogOperationUpdateOne {
	louo.mutation.SetUserID(id)
	return louo
}

// SetUser sets the "user" edge to the User entity.
func (louo *LogOperationUpdateOne) SetUser(u *User) *LogOperationUpdateOne {
	return louo.SetUserID(u.ID)
}

// Mutation returns the LogOperationMutation object of the builder.
func (louo *LogOperationUpdateOne) Mutation() *LogOperationMutation {
	return louo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (louo *LogOperationUpdateOne) ClearUser() *LogOperationUpdateOne {
	louo.mutation.ClearUser()
	return louo
}

// Where appends a list predicates to the LogOperationUpdate builder.
func (louo *LogOperationUpdateOne) Where(ps ...predicate.LogOperation) *LogOperationUpdateOne {
	louo.mutation.Where(ps...)
	return louo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (louo *LogOperationUpdateOne) Select(field string, fields ...string) *LogOperationUpdateOne {
	louo.fields = append([]string{field}, fields...)
	return louo
}

// Save executes the query and returns the updated LogOperation entity.
func (louo *LogOperationUpdateOne) Save(ctx context.Context) (*LogOperation, error) {
	louo.defaults()
	return withHooks(ctx, louo.sqlSave, louo.mutation, louo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (louo *LogOperationUpdateOne) SaveX(ctx context.Context) *LogOperation {
	node, err := louo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (louo *LogOperationUpdateOne) Exec(ctx context.Context) error {
	_, err := louo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (louo *LogOperationUpdateOne) ExecX(ctx context.Context) {
	if err := louo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (louo *LogOperationUpdateOne) defaults() {
	if _, ok := louo.mutation.UpdatedAt(); !ok {
		v := logoperation.UpdateDefaultUpdatedAt()
		louo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (louo *LogOperationUpdateOne) check() error {
	if v, ok := louo.mutation.Headers(); ok {
		if err := logoperation.HeadersValidator(v); err != nil {
			return &ValidationError{Name: "headers", err: fmt.Errorf(`ent: validator failed for field "LogOperation.headers": %w`, err)}
		}
	}
	if _, ok := louo.mutation.UserID(); louo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LogOperation.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (louo *LogOperationUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LogOperationUpdateOne {
	louo.modifiers = append(louo.modifiers, modifiers...)
	return louo
}

func (louo *LogOperationUpdateOne) sqlSave(ctx context.Context) (_node *LogOperation, err error) {
	if err := louo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(logoperation.Table, logoperation.Columns, sqlgraph.NewFieldSpec(logoperation.FieldID, field.TypeUint64))
	id, ok := louo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LogOperation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := louo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logoperation.FieldID)
		for _, f := range fields {
			if !logoperation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != logoperation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := louo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := louo.mutation.UpdatedAt(); ok {
		_spec.SetField(logoperation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := louo.mutation.Method(); ok {
		_spec.SetField(logoperation.FieldMethod, field.TypeString, value)
	}
	if value, ok := louo.mutation.Path(); ok {
		_spec.SetField(logoperation.FieldPath, field.TypeString, value)
	}
	if value, ok := louo.mutation.Headers(); ok {
		_spec.SetField(logoperation.FieldHeaders, field.TypeString, value)
	}
	if value, ok := louo.mutation.Body(); ok {
		_spec.SetField(logoperation.FieldBody, field.TypeString, value)
	}
	if louo.mutation.BodyCleared() {
		_spec.ClearField(logoperation.FieldBody, field.TypeString)
	}
	if value, ok := louo.mutation.StatusCode(); ok {
		_spec.SetField(logoperation.FieldStatusCode, field.TypeInt, value)
	}
	if value, ok := louo.mutation.AddedStatusCode(); ok {
		_spec.AddField(logoperation.FieldStatusCode, field.TypeInt, value)
	}
	if value, ok := louo.mutation.ReqTime(); ok {
		_spec.SetField(logoperation.FieldReqTime, field.TypeTime, value)
	}
	if value, ok := louo.mutation.ResTime(); ok {
		_spec.SetField(logoperation.FieldResTime, field.TypeTime, value)
	}
	if value, ok := louo.mutation.CostTime(); ok {
		_spec.SetField(logoperation.FieldCostTime, field.TypeUint64, value)
	}
	if value, ok := louo.mutation.AddedCostTime(); ok {
		_spec.AddField(logoperation.FieldCostTime, field.TypeUint64, value)
	}
	if louo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logoperation.UserTable,
			Columns: []string{logoperation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := louo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logoperation.UserTable,
			Columns: []string{logoperation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(louo.modifiers...)
	_node = &LogOperation{config: louo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, louo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logoperation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	louo.mutation.done = true
	return _node, nil
}
