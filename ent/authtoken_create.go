// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/OhMinsSup/story-server/ent/authtoken"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// AuthTokenCreate is the builder for creating a AuthToken entity.
type AuthTokenCreate struct {
	config
	mutation *AuthTokenMutation
	hooks    []Hook
}

// SetDisabled sets the "disabled" field.
func (atc *AuthTokenCreate) SetDisabled(b bool) *AuthTokenCreate {
	atc.mutation.SetDisabled(b)
	return atc
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (atc *AuthTokenCreate) SetNillableDisabled(b *bool) *AuthTokenCreate {
	if b != nil {
		atc.SetDisabled(*b)
	}
	return atc
}

// SetCreatedAt sets the "created_at" field.
func (atc *AuthTokenCreate) SetCreatedAt(t time.Time) *AuthTokenCreate {
	atc.mutation.SetCreatedAt(t)
	return atc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atc *AuthTokenCreate) SetNillableCreatedAt(t *time.Time) *AuthTokenCreate {
	if t != nil {
		atc.SetCreatedAt(*t)
	}
	return atc
}

// SetUpdatedAt sets the "updated_at" field.
func (atc *AuthTokenCreate) SetUpdatedAt(t time.Time) *AuthTokenCreate {
	atc.mutation.SetUpdatedAt(t)
	return atc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atc *AuthTokenCreate) SetNillableUpdatedAt(t *time.Time) *AuthTokenCreate {
	if t != nil {
		atc.SetUpdatedAt(*t)
	}
	return atc
}

// SetFkUserID sets the "fk_user_id" field.
func (atc *AuthTokenCreate) SetFkUserID(u uuid.UUID) *AuthTokenCreate {
	atc.mutation.SetFkUserID(u)
	return atc
}

// SetID sets the "id" field.
func (atc *AuthTokenCreate) SetID(u uuid.UUID) *AuthTokenCreate {
	atc.mutation.SetID(u)
	return atc
}

// Mutation returns the AuthTokenMutation object of the builder.
func (atc *AuthTokenCreate) Mutation() *AuthTokenMutation {
	return atc.mutation
}

// Save creates the AuthToken in the database.
func (atc *AuthTokenCreate) Save(ctx context.Context) (*AuthToken, error) {
	var (
		err  error
		node *AuthToken
	)
	atc.defaults()
	if len(atc.hooks) == 0 {
		if err = atc.check(); err != nil {
			return nil, err
		}
		node, err = atc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atc.check(); err != nil {
				return nil, err
			}
			atc.mutation = mutation
			node, err = atc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atc.hooks) - 1; i >= 0; i-- {
			mut = atc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AuthTokenCreate) SaveX(ctx context.Context) *AuthToken {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (atc *AuthTokenCreate) defaults() {
	if _, ok := atc.mutation.Disabled(); !ok {
		v := authtoken.DefaultDisabled
		atc.mutation.SetDisabled(v)
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		v := authtoken.DefaultCreatedAt()
		atc.mutation.SetCreatedAt(v)
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		v := authtoken.DefaultUpdatedAt()
		atc.mutation.SetUpdatedAt(v)
	}
	if _, ok := atc.mutation.ID(); !ok {
		v := authtoken.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AuthTokenCreate) check() error {
	if _, ok := atc.mutation.Disabled(); !ok {
		return &ValidationError{Name: "disabled", err: errors.New("ent: missing required field \"disabled\"")}
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := atc.mutation.FkUserID(); !ok {
		return &ValidationError{Name: "fk_user_id", err: errors.New("ent: missing required field \"fk_user_id\"")}
	}
	return nil
}

func (atc *AuthTokenCreate) sqlSave(ctx context.Context) (*AuthToken, error) {
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (atc *AuthTokenCreate) createSpec() (*AuthToken, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthToken{config: atc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: authtoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: authtoken.FieldID,
			},
		}
	)
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := atc.mutation.Disabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: authtoken.FieldDisabled,
		})
		_node.Disabled = value
	}
	if value, ok := atc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authtoken.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := atc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authtoken.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := atc.mutation.FkUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: authtoken.FieldFkUserID,
		})
		_node.FkUserID = value
	}
	return _node, _spec
}

// AuthTokenCreateBulk is the builder for creating many AuthToken entities in bulk.
type AuthTokenCreateBulk struct {
	config
	builders []*AuthTokenCreate
}

// Save creates the AuthToken entities in the database.
func (atcb *AuthTokenCreateBulk) Save(ctx context.Context) ([]*AuthToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AuthToken, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthTokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AuthTokenCreateBulk) SaveX(ctx context.Context) []*AuthToken {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}