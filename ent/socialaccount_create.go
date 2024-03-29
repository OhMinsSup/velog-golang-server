// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/OhMinsSup/story-server/ent/socialaccount"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// SocialAccountCreate is the builder for creating a SocialAccount entity.
type SocialAccountCreate struct {
	config
	mutation *SocialAccountMutation
	hooks    []Hook
}

// SetSocialID sets the "social_id" field.
func (sac *SocialAccountCreate) SetSocialID(s string) *SocialAccountCreate {
	sac.mutation.SetSocialID(s)
	return sac
}

// SetAccessToken sets the "access_token" field.
func (sac *SocialAccountCreate) SetAccessToken(s string) *SocialAccountCreate {
	sac.mutation.SetAccessToken(s)
	return sac
}

// SetProvider sets the "provider" field.
func (sac *SocialAccountCreate) SetProvider(s string) *SocialAccountCreate {
	sac.mutation.SetProvider(s)
	return sac
}

// SetFkUserID sets the "fk_user_id" field.
func (sac *SocialAccountCreate) SetFkUserID(u uuid.UUID) *SocialAccountCreate {
	sac.mutation.SetFkUserID(u)
	return sac
}

// SetCreatedAt sets the "created_at" field.
func (sac *SocialAccountCreate) SetCreatedAt(t time.Time) *SocialAccountCreate {
	sac.mutation.SetCreatedAt(t)
	return sac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sac *SocialAccountCreate) SetNillableCreatedAt(t *time.Time) *SocialAccountCreate {
	if t != nil {
		sac.SetCreatedAt(*t)
	}
	return sac
}

// SetUpdatedAt sets the "updated_at" field.
func (sac *SocialAccountCreate) SetUpdatedAt(t time.Time) *SocialAccountCreate {
	sac.mutation.SetUpdatedAt(t)
	return sac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sac *SocialAccountCreate) SetNillableUpdatedAt(t *time.Time) *SocialAccountCreate {
	if t != nil {
		sac.SetUpdatedAt(*t)
	}
	return sac
}

// SetID sets the "id" field.
func (sac *SocialAccountCreate) SetID(u uuid.UUID) *SocialAccountCreate {
	sac.mutation.SetID(u)
	return sac
}

// Mutation returns the SocialAccountMutation object of the builder.
func (sac *SocialAccountCreate) Mutation() *SocialAccountMutation {
	return sac.mutation
}

// Save creates the SocialAccount in the database.
func (sac *SocialAccountCreate) Save(ctx context.Context) (*SocialAccount, error) {
	var (
		err  error
		node *SocialAccount
	)
	sac.defaults()
	if len(sac.hooks) == 0 {
		if err = sac.check(); err != nil {
			return nil, err
		}
		node, err = sac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SocialAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sac.check(); err != nil {
				return nil, err
			}
			sac.mutation = mutation
			node, err = sac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sac.hooks) - 1; i >= 0; i-- {
			mut = sac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sac *SocialAccountCreate) SaveX(ctx context.Context) *SocialAccount {
	v, err := sac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sac *SocialAccountCreate) defaults() {
	if _, ok := sac.mutation.CreatedAt(); !ok {
		v := socialaccount.DefaultCreatedAt()
		sac.mutation.SetCreatedAt(v)
	}
	if _, ok := sac.mutation.UpdatedAt(); !ok {
		v := socialaccount.DefaultUpdatedAt()
		sac.mutation.SetUpdatedAt(v)
	}
	if _, ok := sac.mutation.ID(); !ok {
		v := socialaccount.DefaultID()
		sac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sac *SocialAccountCreate) check() error {
	if _, ok := sac.mutation.SocialID(); !ok {
		return &ValidationError{Name: "social_id", err: errors.New("ent: missing required field \"social_id\"")}
	}
	if v, ok := sac.mutation.SocialID(); ok {
		if err := socialaccount.SocialIDValidator(v); err != nil {
			return &ValidationError{Name: "social_id", err: fmt.Errorf("ent: validator failed for field \"social_id\": %w", err)}
		}
	}
	if _, ok := sac.mutation.AccessToken(); !ok {
		return &ValidationError{Name: "access_token", err: errors.New("ent: missing required field \"access_token\"")}
	}
	if v, ok := sac.mutation.AccessToken(); ok {
		if err := socialaccount.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf("ent: validator failed for field \"access_token\": %w", err)}
		}
	}
	if _, ok := sac.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New("ent: missing required field \"provider\"")}
	}
	if v, ok := sac.mutation.Provider(); ok {
		if err := socialaccount.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("ent: validator failed for field \"provider\": %w", err)}
		}
	}
	if _, ok := sac.mutation.FkUserID(); !ok {
		return &ValidationError{Name: "fk_user_id", err: errors.New("ent: missing required field \"fk_user_id\"")}
	}
	if _, ok := sac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := sac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (sac *SocialAccountCreate) sqlSave(ctx context.Context) (*SocialAccount, error) {
	_node, _spec := sac.createSpec()
	if err := sqlgraph.CreateNode(ctx, sac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (sac *SocialAccountCreate) createSpec() (*SocialAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &SocialAccount{config: sac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: socialaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: socialaccount.FieldID,
			},
		}
	)
	if id, ok := sac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sac.mutation.SocialID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialaccount.FieldSocialID,
		})
		_node.SocialID = value
	}
	if value, ok := sac.mutation.AccessToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialaccount.FieldAccessToken,
		})
		_node.AccessToken = value
	}
	if value, ok := sac.mutation.Provider(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialaccount.FieldProvider,
		})
		_node.Provider = value
	}
	if value, ok := sac.mutation.FkUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: socialaccount.FieldFkUserID,
		})
		_node.FkUserID = value
	}
	if value, ok := sac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: socialaccount.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: socialaccount.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// SocialAccountCreateBulk is the builder for creating many SocialAccount entities in bulk.
type SocialAccountCreateBulk struct {
	config
	builders []*SocialAccountCreate
}

// Save creates the SocialAccount entities in the database.
func (sacb *SocialAccountCreateBulk) Save(ctx context.Context) ([]*SocialAccount, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sacb.builders))
	nodes := make([]*SocialAccount, len(sacb.builders))
	mutators := make([]Mutator, len(sacb.builders))
	for i := range sacb.builders {
		func(i int, root context.Context) {
			builder := sacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SocialAccountMutation)
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
					_, err = mutators[i+1].Mutate(root, sacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sacb *SocialAccountCreateBulk) SaveX(ctx context.Context) []*SocialAccount {
	v, err := sacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
