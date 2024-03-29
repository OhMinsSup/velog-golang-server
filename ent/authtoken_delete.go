// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OhMinsSup/story-server/ent/authtoken"
	"github.com/OhMinsSup/story-server/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AuthTokenDelete is the builder for deleting a AuthToken entity.
type AuthTokenDelete struct {
	config
	hooks    []Hook
	mutation *AuthTokenMutation
}

// Where adds a new predicate to the AuthTokenDelete builder.
func (atd *AuthTokenDelete) Where(ps ...predicate.AuthToken) *AuthTokenDelete {
	atd.mutation.predicates = append(atd.mutation.predicates, ps...)
	return atd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atd *AuthTokenDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atd.hooks) == 0 {
		affected, err = atd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atd.mutation = mutation
			affected, err = atd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atd.hooks) - 1; i >= 0; i-- {
			mut = atd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (atd *AuthTokenDelete) ExecX(ctx context.Context) int {
	n, err := atd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atd *AuthTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: authtoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: authtoken.FieldID,
			},
		},
	}
	if ps := atd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, atd.driver, _spec)
}

// AuthTokenDeleteOne is the builder for deleting a single AuthToken entity.
type AuthTokenDeleteOne struct {
	atd *AuthTokenDelete
}

// Exec executes the deletion query.
func (atdo *AuthTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := atdo.atd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authtoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atdo *AuthTokenDeleteOne) ExecX(ctx context.Context) {
	atdo.atd.ExecX(ctx)
}
