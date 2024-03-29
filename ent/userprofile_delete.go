// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OhMinsSup/story-server/ent/predicate"
	"github.com/OhMinsSup/story-server/ent/userprofile"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// UserProfileDelete is the builder for deleting a UserProfile entity.
type UserProfileDelete struct {
	config
	hooks    []Hook
	mutation *UserProfileMutation
}

// Where adds a new predicate to the UserProfileDelete builder.
func (upd *UserProfileDelete) Where(ps ...predicate.UserProfile) *UserProfileDelete {
	upd.mutation.predicates = append(upd.mutation.predicates, ps...)
	return upd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (upd *UserProfileDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(upd.hooks) == 0 {
		affected, err = upd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			upd.mutation = mutation
			affected, err = upd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(upd.hooks) - 1; i >= 0; i-- {
			mut = upd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (upd *UserProfileDelete) ExecX(ctx context.Context) int {
	n, err := upd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (upd *UserProfileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userprofile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userprofile.FieldID,
			},
		},
	}
	if ps := upd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, upd.driver, _spec)
}

// UserProfileDeleteOne is the builder for deleting a single UserProfile entity.
type UserProfileDeleteOne struct {
	upd *UserProfileDelete
}

// Exec executes the deletion query.
func (updo *UserProfileDeleteOne) Exec(ctx context.Context) error {
	n, err := updo.upd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userprofile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (updo *UserProfileDeleteOne) ExecX(ctx context.Context) {
	updo.upd.ExecX(ctx)
}
