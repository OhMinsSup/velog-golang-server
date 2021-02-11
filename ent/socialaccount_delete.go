// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OhMinsSup/story-server/ent/predicate"
	"github.com/OhMinsSup/story-server/ent/socialaccount"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// SocialAccountDelete is the builder for deleting a SocialAccount entity.
type SocialAccountDelete struct {
	config
	hooks    []Hook
	mutation *SocialAccountMutation
}

// Where adds a new predicate to the SocialAccountDelete builder.
func (sad *SocialAccountDelete) Where(ps ...predicate.SocialAccount) *SocialAccountDelete {
	sad.mutation.predicates = append(sad.mutation.predicates, ps...)
	return sad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sad *SocialAccountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sad.hooks) == 0 {
		affected, err = sad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SocialAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sad.mutation = mutation
			affected, err = sad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sad.hooks) - 1; i >= 0; i-- {
			mut = sad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sad *SocialAccountDelete) ExecX(ctx context.Context) int {
	n, err := sad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sad *SocialAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: socialaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: socialaccount.FieldID,
			},
		},
	}
	if ps := sad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sad.driver, _spec)
}

// SocialAccountDeleteOne is the builder for deleting a single SocialAccount entity.
type SocialAccountDeleteOne struct {
	sad *SocialAccountDelete
}

// Exec executes the deletion query.
func (sado *SocialAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := sado.sad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{socialaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sado *SocialAccountDeleteOne) ExecX(ctx context.Context) {
	sado.sad.ExecX(ctx)
}
