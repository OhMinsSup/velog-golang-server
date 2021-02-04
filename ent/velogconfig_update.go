// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/OhMinsSup/story-server/ent/predicate"
	"github.com/OhMinsSup/story-server/ent/user"
	"github.com/OhMinsSup/story-server/ent/velogconfig"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// VelogConfigUpdate is the builder for updating VelogConfig entities.
type VelogConfigUpdate struct {
	config
	hooks    []Hook
	mutation *VelogConfigMutation
}

// Where adds a new predicate for the VelogConfigUpdate builder.
func (vcu *VelogConfigUpdate) Where(ps ...predicate.VelogConfig) *VelogConfigUpdate {
	vcu.mutation.predicates = append(vcu.mutation.predicates, ps...)
	return vcu
}

// SetTitle sets the "title" field.
func (vcu *VelogConfigUpdate) SetTitle(s string) *VelogConfigUpdate {
	vcu.mutation.SetTitle(s)
	return vcu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (vcu *VelogConfigUpdate) SetNillableTitle(s *string) *VelogConfigUpdate {
	if s != nil {
		vcu.SetTitle(*s)
	}
	return vcu
}

// ClearTitle clears the value of the "title" field.
func (vcu *VelogConfigUpdate) ClearTitle() *VelogConfigUpdate {
	vcu.mutation.ClearTitle()
	return vcu
}

// SetLogoTitle sets the "logo_title" field.
func (vcu *VelogConfigUpdate) SetLogoTitle(s string) *VelogConfigUpdate {
	vcu.mutation.SetLogoTitle(s)
	return vcu
}

// SetNillableLogoTitle sets the "logo_title" field if the given value is not nil.
func (vcu *VelogConfigUpdate) SetNillableLogoTitle(s *string) *VelogConfigUpdate {
	if s != nil {
		vcu.SetLogoTitle(*s)
	}
	return vcu
}

// ClearLogoTitle clears the value of the "logo_title" field.
func (vcu *VelogConfigUpdate) ClearLogoTitle() *VelogConfigUpdate {
	vcu.mutation.ClearLogoTitle()
	return vcu
}

// SetUpdatedAt sets the "updated_at" field.
func (vcu *VelogConfigUpdate) SetUpdatedAt(t time.Time) *VelogConfigUpdate {
	vcu.mutation.SetUpdatedAt(t)
	return vcu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vcu *VelogConfigUpdate) SetUserID(id uuid.UUID) *VelogConfigUpdate {
	vcu.mutation.SetUserID(id)
	return vcu
}

// SetUser sets the "user" edge to the User entity.
func (vcu *VelogConfigUpdate) SetUser(u *User) *VelogConfigUpdate {
	return vcu.SetUserID(u.ID)
}

// Mutation returns the VelogConfigMutation object of the builder.
func (vcu *VelogConfigUpdate) Mutation() *VelogConfigMutation {
	return vcu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vcu *VelogConfigUpdate) ClearUser() *VelogConfigUpdate {
	vcu.mutation.ClearUser()
	return vcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vcu *VelogConfigUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	vcu.defaults()
	if len(vcu.hooks) == 0 {
		if err = vcu.check(); err != nil {
			return 0, err
		}
		affected, err = vcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VelogConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vcu.check(); err != nil {
				return 0, err
			}
			vcu.mutation = mutation
			affected, err = vcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vcu.hooks) - 1; i >= 0; i-- {
			mut = vcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vcu *VelogConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := vcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vcu *VelogConfigUpdate) Exec(ctx context.Context) error {
	_, err := vcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcu *VelogConfigUpdate) ExecX(ctx context.Context) {
	if err := vcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vcu *VelogConfigUpdate) defaults() {
	if _, ok := vcu.mutation.UpdatedAt(); !ok {
		v := velogconfig.UpdateDefaultUpdatedAt()
		vcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vcu *VelogConfigUpdate) check() error {
	if _, ok := vcu.mutation.UserID(); vcu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (vcu *VelogConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   velogconfig.Table,
			Columns: velogconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: velogconfig.FieldID,
			},
		},
	}
	if ps := vcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vcu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: velogconfig.FieldTitle,
		})
	}
	if vcu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: velogconfig.FieldTitle,
		})
	}
	if value, ok := vcu.mutation.LogoTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: velogconfig.FieldLogoTitle,
		})
	}
	if vcu.mutation.LogoTitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: velogconfig.FieldLogoTitle,
		})
	}
	if value, ok := vcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: velogconfig.FieldUpdatedAt,
		})
	}
	if vcu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   velogconfig.UserTable,
			Columns: []string{velogconfig.UserColumn},
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
	if nodes := vcu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   velogconfig.UserTable,
			Columns: []string{velogconfig.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, vcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{velogconfig.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// VelogConfigUpdateOne is the builder for updating a single VelogConfig entity.
type VelogConfigUpdateOne struct {
	config
	hooks    []Hook
	mutation *VelogConfigMutation
}

// SetTitle sets the "title" field.
func (vcuo *VelogConfigUpdateOne) SetTitle(s string) *VelogConfigUpdateOne {
	vcuo.mutation.SetTitle(s)
	return vcuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (vcuo *VelogConfigUpdateOne) SetNillableTitle(s *string) *VelogConfigUpdateOne {
	if s != nil {
		vcuo.SetTitle(*s)
	}
	return vcuo
}

// ClearTitle clears the value of the "title" field.
func (vcuo *VelogConfigUpdateOne) ClearTitle() *VelogConfigUpdateOne {
	vcuo.mutation.ClearTitle()
	return vcuo
}

// SetLogoTitle sets the "logo_title" field.
func (vcuo *VelogConfigUpdateOne) SetLogoTitle(s string) *VelogConfigUpdateOne {
	vcuo.mutation.SetLogoTitle(s)
	return vcuo
}

// SetNillableLogoTitle sets the "logo_title" field if the given value is not nil.
func (vcuo *VelogConfigUpdateOne) SetNillableLogoTitle(s *string) *VelogConfigUpdateOne {
	if s != nil {
		vcuo.SetLogoTitle(*s)
	}
	return vcuo
}

// ClearLogoTitle clears the value of the "logo_title" field.
func (vcuo *VelogConfigUpdateOne) ClearLogoTitle() *VelogConfigUpdateOne {
	vcuo.mutation.ClearLogoTitle()
	return vcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (vcuo *VelogConfigUpdateOne) SetUpdatedAt(t time.Time) *VelogConfigUpdateOne {
	vcuo.mutation.SetUpdatedAt(t)
	return vcuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vcuo *VelogConfigUpdateOne) SetUserID(id uuid.UUID) *VelogConfigUpdateOne {
	vcuo.mutation.SetUserID(id)
	return vcuo
}

// SetUser sets the "user" edge to the User entity.
func (vcuo *VelogConfigUpdateOne) SetUser(u *User) *VelogConfigUpdateOne {
	return vcuo.SetUserID(u.ID)
}

// Mutation returns the VelogConfigMutation object of the builder.
func (vcuo *VelogConfigUpdateOne) Mutation() *VelogConfigMutation {
	return vcuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vcuo *VelogConfigUpdateOne) ClearUser() *VelogConfigUpdateOne {
	vcuo.mutation.ClearUser()
	return vcuo
}

// Save executes the query and returns the updated VelogConfig entity.
func (vcuo *VelogConfigUpdateOne) Save(ctx context.Context) (*VelogConfig, error) {
	var (
		err  error
		node *VelogConfig
	)
	vcuo.defaults()
	if len(vcuo.hooks) == 0 {
		if err = vcuo.check(); err != nil {
			return nil, err
		}
		node, err = vcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VelogConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vcuo.check(); err != nil {
				return nil, err
			}
			vcuo.mutation = mutation
			node, err = vcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vcuo.hooks) - 1; i >= 0; i-- {
			mut = vcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vcuo *VelogConfigUpdateOne) SaveX(ctx context.Context) *VelogConfig {
	node, err := vcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vcuo *VelogConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := vcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcuo *VelogConfigUpdateOne) ExecX(ctx context.Context) {
	if err := vcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vcuo *VelogConfigUpdateOne) defaults() {
	if _, ok := vcuo.mutation.UpdatedAt(); !ok {
		v := velogconfig.UpdateDefaultUpdatedAt()
		vcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vcuo *VelogConfigUpdateOne) check() error {
	if _, ok := vcuo.mutation.UserID(); vcuo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (vcuo *VelogConfigUpdateOne) sqlSave(ctx context.Context) (_node *VelogConfig, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   velogconfig.Table,
			Columns: velogconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: velogconfig.FieldID,
			},
		},
	}
	id, ok := vcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing VelogConfig.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := vcuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: velogconfig.FieldTitle,
		})
	}
	if vcuo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: velogconfig.FieldTitle,
		})
	}
	if value, ok := vcuo.mutation.LogoTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: velogconfig.FieldLogoTitle,
		})
	}
	if vcuo.mutation.LogoTitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: velogconfig.FieldLogoTitle,
		})
	}
	if value, ok := vcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: velogconfig.FieldUpdatedAt,
		})
	}
	if vcuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   velogconfig.UserTable,
			Columns: []string{velogconfig.UserColumn},
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
	if nodes := vcuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   velogconfig.UserTable,
			Columns: []string{velogconfig.UserColumn},
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
	_node = &VelogConfig{config: vcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{velogconfig.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
