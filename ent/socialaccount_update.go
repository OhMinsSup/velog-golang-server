// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/OhMinsSup/story-server/ent/predicate"
	"github.com/OhMinsSup/story-server/ent/socialaccount"
	"github.com/OhMinsSup/story-server/ent/user"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// SocialAccountUpdate is the builder for updating SocialAccount entities.
type SocialAccountUpdate struct {
	config
	hooks    []Hook
	mutation *SocialAccountMutation
}

// Where adds a new predicate for the SocialAccountUpdate builder.
func (sau *SocialAccountUpdate) Where(ps ...predicate.SocialAccount) *SocialAccountUpdate {
	sau.mutation.predicates = append(sau.mutation.predicates, ps...)
	return sau
}

// SetSocialID sets the "social_id" field.
func (sau *SocialAccountUpdate) SetSocialID(s string) *SocialAccountUpdate {
	sau.mutation.SetSocialID(s)
	return sau
}

// SetAccessToken sets the "access_token" field.
func (sau *SocialAccountUpdate) SetAccessToken(s string) *SocialAccountUpdate {
	sau.mutation.SetAccessToken(s)
	return sau
}

// SetProvider sets the "provider" field.
func (sau *SocialAccountUpdate) SetProvider(s string) *SocialAccountUpdate {
	sau.mutation.SetProvider(s)
	return sau
}

// SetUpdatedAt sets the "updated_at" field.
func (sau *SocialAccountUpdate) SetUpdatedAt(t time.Time) *SocialAccountUpdate {
	sau.mutation.SetUpdatedAt(t)
	return sau
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sau *SocialAccountUpdate) SetUserID(id uuid.UUID) *SocialAccountUpdate {
	sau.mutation.SetUserID(id)
	return sau
}

// SetUser sets the "user" edge to the User entity.
func (sau *SocialAccountUpdate) SetUser(u *User) *SocialAccountUpdate {
	return sau.SetUserID(u.ID)
}

// Mutation returns the SocialAccountMutation object of the builder.
func (sau *SocialAccountUpdate) Mutation() *SocialAccountMutation {
	return sau.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (sau *SocialAccountUpdate) ClearUser() *SocialAccountUpdate {
	sau.mutation.ClearUser()
	return sau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *SocialAccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sau.defaults()
	if len(sau.hooks) == 0 {
		if err = sau.check(); err != nil {
			return 0, err
		}
		affected, err = sau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SocialAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sau.check(); err != nil {
				return 0, err
			}
			sau.mutation = mutation
			affected, err = sau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sau.hooks) - 1; i >= 0; i-- {
			mut = sau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sau *SocialAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *SocialAccountUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *SocialAccountUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sau *SocialAccountUpdate) defaults() {
	if _, ok := sau.mutation.UpdatedAt(); !ok {
		v := socialaccount.UpdateDefaultUpdatedAt()
		sau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sau *SocialAccountUpdate) check() error {
	if v, ok := sau.mutation.SocialID(); ok {
		if err := socialaccount.SocialIDValidator(v); err != nil {
			return &ValidationError{Name: "social_id", err: fmt.Errorf("ent: validator failed for field \"social_id\": %w", err)}
		}
	}
	if v, ok := sau.mutation.AccessToken(); ok {
		if err := socialaccount.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf("ent: validator failed for field \"access_token\": %w", err)}
		}
	}
	if v, ok := sau.mutation.Provider(); ok {
		if err := socialaccount.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("ent: validator failed for field \"provider\": %w", err)}
		}
	}
	if _, ok := sau.mutation.UserID(); sau.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (sau *SocialAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   socialaccount.Table,
			Columns: socialaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: socialaccount.FieldID,
			},
		},
	}
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.SocialID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialaccount.FieldSocialID,
		})
	}
	if value, ok := sau.mutation.AccessToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialaccount.FieldAccessToken,
		})
	}
	if value, ok := sau.mutation.Provider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialaccount.FieldProvider,
		})
	}
	if value, ok := sau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: socialaccount.FieldUpdatedAt,
		})
	}
	if sau.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   socialaccount.UserTable,
			Columns: []string{socialaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   socialaccount.UserTable,
			Columns: []string{socialaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{socialaccount.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SocialAccountUpdateOne is the builder for updating a single SocialAccount entity.
type SocialAccountUpdateOne struct {
	config
	hooks    []Hook
	mutation *SocialAccountMutation
}

// SetSocialID sets the "social_id" field.
func (sauo *SocialAccountUpdateOne) SetSocialID(s string) *SocialAccountUpdateOne {
	sauo.mutation.SetSocialID(s)
	return sauo
}

// SetAccessToken sets the "access_token" field.
func (sauo *SocialAccountUpdateOne) SetAccessToken(s string) *SocialAccountUpdateOne {
	sauo.mutation.SetAccessToken(s)
	return sauo
}

// SetProvider sets the "provider" field.
func (sauo *SocialAccountUpdateOne) SetProvider(s string) *SocialAccountUpdateOne {
	sauo.mutation.SetProvider(s)
	return sauo
}

// SetUpdatedAt sets the "updated_at" field.
func (sauo *SocialAccountUpdateOne) SetUpdatedAt(t time.Time) *SocialAccountUpdateOne {
	sauo.mutation.SetUpdatedAt(t)
	return sauo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sauo *SocialAccountUpdateOne) SetUserID(id uuid.UUID) *SocialAccountUpdateOne {
	sauo.mutation.SetUserID(id)
	return sauo
}

// SetUser sets the "user" edge to the User entity.
func (sauo *SocialAccountUpdateOne) SetUser(u *User) *SocialAccountUpdateOne {
	return sauo.SetUserID(u.ID)
}

// Mutation returns the SocialAccountMutation object of the builder.
func (sauo *SocialAccountUpdateOne) Mutation() *SocialAccountMutation {
	return sauo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (sauo *SocialAccountUpdateOne) ClearUser() *SocialAccountUpdateOne {
	sauo.mutation.ClearUser()
	return sauo
}

// Save executes the query and returns the updated SocialAccount entity.
func (sauo *SocialAccountUpdateOne) Save(ctx context.Context) (*SocialAccount, error) {
	var (
		err  error
		node *SocialAccount
	)
	sauo.defaults()
	if len(sauo.hooks) == 0 {
		if err = sauo.check(); err != nil {
			return nil, err
		}
		node, err = sauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SocialAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sauo.check(); err != nil {
				return nil, err
			}
			sauo.mutation = mutation
			node, err = sauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sauo.hooks) - 1; i >= 0; i-- {
			mut = sauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *SocialAccountUpdateOne) SaveX(ctx context.Context) *SocialAccount {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *SocialAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *SocialAccountUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sauo *SocialAccountUpdateOne) defaults() {
	if _, ok := sauo.mutation.UpdatedAt(); !ok {
		v := socialaccount.UpdateDefaultUpdatedAt()
		sauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sauo *SocialAccountUpdateOne) check() error {
	if v, ok := sauo.mutation.SocialID(); ok {
		if err := socialaccount.SocialIDValidator(v); err != nil {
			return &ValidationError{Name: "social_id", err: fmt.Errorf("ent: validator failed for field \"social_id\": %w", err)}
		}
	}
	if v, ok := sauo.mutation.AccessToken(); ok {
		if err := socialaccount.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf("ent: validator failed for field \"access_token\": %w", err)}
		}
	}
	if v, ok := sauo.mutation.Provider(); ok {
		if err := socialaccount.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("ent: validator failed for field \"provider\": %w", err)}
		}
	}
	if _, ok := sauo.mutation.UserID(); sauo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (sauo *SocialAccountUpdateOne) sqlSave(ctx context.Context) (_node *SocialAccount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   socialaccount.Table,
			Columns: socialaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: socialaccount.FieldID,
			},
		},
	}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SocialAccount.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := sauo.mutation.SocialID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialaccount.FieldSocialID,
		})
	}
	if value, ok := sauo.mutation.AccessToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialaccount.FieldAccessToken,
		})
	}
	if value, ok := sauo.mutation.Provider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialaccount.FieldProvider,
		})
	}
	if value, ok := sauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: socialaccount.FieldUpdatedAt,
		})
	}
	if sauo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   socialaccount.UserTable,
			Columns: []string{socialaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   socialaccount.UserTable,
			Columns: []string{socialaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SocialAccount{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{socialaccount.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
