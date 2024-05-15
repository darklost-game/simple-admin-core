// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/rpc/ent/loglogin"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
)

// LogLoginDelete is the builder for deleting a LogLogin entity.
type LogLoginDelete struct {
	config
	hooks    []Hook
	mutation *LogLoginMutation
}

// Where appends a list predicates to the LogLoginDelete builder.
func (lld *LogLoginDelete) Where(ps ...predicate.LogLogin) *LogLoginDelete {
	lld.mutation.Where(ps...)
	return lld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lld *LogLoginDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lld.sqlExec, lld.mutation, lld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lld *LogLoginDelete) ExecX(ctx context.Context) int {
	n, err := lld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lld *LogLoginDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(loglogin.Table, sqlgraph.NewFieldSpec(loglogin.FieldID, field.TypeUint64))
	if ps := lld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lld.mutation.done = true
	return affected, err
}

// LogLoginDeleteOne is the builder for deleting a single LogLogin entity.
type LogLoginDeleteOne struct {
	lld *LogLoginDelete
}

// Where appends a list predicates to the LogLoginDelete builder.
func (lldo *LogLoginDeleteOne) Where(ps ...predicate.LogLogin) *LogLoginDeleteOne {
	lldo.lld.mutation.Where(ps...)
	return lldo
}

// Exec executes the deletion query.
func (lldo *LogLoginDeleteOne) Exec(ctx context.Context) error {
	n, err := lldo.lld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{loglogin.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lldo *LogLoginDeleteOne) ExecX(ctx context.Context) {
	if err := lldo.Exec(ctx); err != nil {
		panic(err)
	}
}
