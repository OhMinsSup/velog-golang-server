// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/OhMinsSup/story-server/ent/predicate"
	"github.com/OhMinsSup/story-server/ent/user"
	"github.com/OhMinsSup/story-server/ent/userprofile"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserProfileUpdate is the builder for updating UserProfile entities.
type UserProfileUpdate struct {
	config
	hooks    []Hook
	mutation *UserProfileMutation
}

// Where adds a new predicate for the UserProfileUpdate builder.
func (upu *UserProfileUpdate) Where(ps ...predicate.UserProfile) *UserProfileUpdate {
	upu.mutation.predicates = append(upu.mutation.predicates, ps...)
	return upu
}

// SetDisplayName sets the "display_name" field.
func (upu *UserProfileUpdate) SetDisplayName(s string) *UserProfileUpdate {
	upu.mutation.SetDisplayName(s)
	return upu
}

// SetShortBio sets the "short_bio" field.
func (upu *UserProfileUpdate) SetShortBio(s string) *UserProfileUpdate {
	upu.mutation.SetShortBio(s)
	return upu
}

// SetAbout sets the "about" field.
func (upu *UserProfileUpdate) SetAbout(s string) *UserProfileUpdate {
	upu.mutation.SetAbout(s)
	return upu
}

// SetNillableAbout sets the "about" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableAbout(s *string) *UserProfileUpdate {
	if s != nil {
		upu.SetAbout(*s)
	}
	return upu
}

// ClearAbout clears the value of the "about" field.
func (upu *UserProfileUpdate) ClearAbout() *UserProfileUpdate {
	upu.mutation.ClearAbout()
	return upu
}

// SetProfileLinks sets the "profile_links" field.
func (upu *UserProfileUpdate) SetProfileLinks(s []string) *UserProfileUpdate {
	upu.mutation.SetProfileLinks(s)
	return upu
}

// ClearProfileLinks clears the value of the "profile_links" field.
func (upu *UserProfileUpdate) ClearProfileLinks() *UserProfileUpdate {
	upu.mutation.ClearProfileLinks()
	return upu
}

// SetThumbnail sets the "thumbnail" field.
func (upu *UserProfileUpdate) SetThumbnail(s string) *UserProfileUpdate {
	upu.mutation.SetThumbnail(s)
	return upu
}

// SetNillableThumbnail sets the "thumbnail" field if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableThumbnail(s *string) *UserProfileUpdate {
	if s != nil {
		upu.SetThumbnail(*s)
	}
	return upu
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (upu *UserProfileUpdate) ClearThumbnail() *UserProfileUpdate {
	upu.mutation.ClearThumbnail()
	return upu
}

// SetUpdatedAt sets the "updated_at" field.
func (upu *UserProfileUpdate) SetUpdatedAt(t time.Time) *UserProfileUpdate {
	upu.mutation.SetUpdatedAt(t)
	return upu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (upu *UserProfileUpdate) SetUserID(id uuid.UUID) *UserProfileUpdate {
	upu.mutation.SetUserID(id)
	return upu
}

// SetUser sets the "user" edge to the User entity.
func (upu *UserProfileUpdate) SetUser(u *User) *UserProfileUpdate {
	return upu.SetUserID(u.ID)
}

// Mutation returns the UserProfileMutation object of the builder.
func (upu *UserProfileUpdate) Mutation() *UserProfileMutation {
	return upu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (upu *UserProfileUpdate) ClearUser() *UserProfileUpdate {
	upu.mutation.ClearUser()
	return upu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upu *UserProfileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	upu.defaults()
	if len(upu.hooks) == 0 {
		if err = upu.check(); err != nil {
			return 0, err
		}
		affected, err = upu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upu.check(); err != nil {
				return 0, err
			}
			upu.mutation = mutation
			affected, err = upu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(upu.hooks) - 1; i >= 0; i-- {
			mut = upu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (upu *UserProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := upu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upu *UserProfileUpdate) Exec(ctx context.Context) error {
	_, err := upu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upu *UserProfileUpdate) ExecX(ctx context.Context) {
	if err := upu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upu *UserProfileUpdate) defaults() {
	if _, ok := upu.mutation.UpdatedAt(); !ok {
		v := userprofile.UpdateDefaultUpdatedAt()
		upu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upu *UserProfileUpdate) check() error {
	if v, ok := upu.mutation.DisplayName(); ok {
		if err := userprofile.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf("ent: validator failed for field \"display_name\": %w", err)}
		}
	}
	if v, ok := upu.mutation.ShortBio(); ok {
		if err := userprofile.ShortBioValidator(v); err != nil {
			return &ValidationError{Name: "short_bio", err: fmt.Errorf("ent: validator failed for field \"short_bio\": %w", err)}
		}
	}
	if _, ok := upu.mutation.UserID(); upu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (upu *UserProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userprofile.Table,
			Columns: userprofile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userprofile.FieldID,
			},
		},
	}
	if ps := upu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upu.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldDisplayName,
		})
	}
	if value, ok := upu.mutation.ShortBio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldShortBio,
		})
	}
	if value, ok := upu.mutation.About(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldAbout,
		})
	}
	if upu.mutation.AboutCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldAbout,
		})
	}
	if value, ok := upu.mutation.ProfileLinks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: userprofile.FieldProfileLinks,
		})
	}
	if upu.mutation.ProfileLinksCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: userprofile.FieldProfileLinks,
		})
	}
	if value, ok := upu.mutation.Thumbnail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldThumbnail,
		})
	}
	if upu.mutation.ThumbnailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldThumbnail,
		})
	}
	if value, ok := upu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userprofile.FieldUpdatedAt,
		})
	}
	if upu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.UserTable,
			Columns: []string{userprofile.UserColumn},
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
	if nodes := upu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.UserTable,
			Columns: []string{userprofile.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, upu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprofile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserProfileUpdateOne is the builder for updating a single UserProfile entity.
type UserProfileUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserProfileMutation
}

// SetDisplayName sets the "display_name" field.
func (upuo *UserProfileUpdateOne) SetDisplayName(s string) *UserProfileUpdateOne {
	upuo.mutation.SetDisplayName(s)
	return upuo
}

// SetShortBio sets the "short_bio" field.
func (upuo *UserProfileUpdateOne) SetShortBio(s string) *UserProfileUpdateOne {
	upuo.mutation.SetShortBio(s)
	return upuo
}

// SetAbout sets the "about" field.
func (upuo *UserProfileUpdateOne) SetAbout(s string) *UserProfileUpdateOne {
	upuo.mutation.SetAbout(s)
	return upuo
}

// SetNillableAbout sets the "about" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableAbout(s *string) *UserProfileUpdateOne {
	if s != nil {
		upuo.SetAbout(*s)
	}
	return upuo
}

// ClearAbout clears the value of the "about" field.
func (upuo *UserProfileUpdateOne) ClearAbout() *UserProfileUpdateOne {
	upuo.mutation.ClearAbout()
	return upuo
}

// SetProfileLinks sets the "profile_links" field.
func (upuo *UserProfileUpdateOne) SetProfileLinks(s []string) *UserProfileUpdateOne {
	upuo.mutation.SetProfileLinks(s)
	return upuo
}

// ClearProfileLinks clears the value of the "profile_links" field.
func (upuo *UserProfileUpdateOne) ClearProfileLinks() *UserProfileUpdateOne {
	upuo.mutation.ClearProfileLinks()
	return upuo
}

// SetThumbnail sets the "thumbnail" field.
func (upuo *UserProfileUpdateOne) SetThumbnail(s string) *UserProfileUpdateOne {
	upuo.mutation.SetThumbnail(s)
	return upuo
}

// SetNillableThumbnail sets the "thumbnail" field if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableThumbnail(s *string) *UserProfileUpdateOne {
	if s != nil {
		upuo.SetThumbnail(*s)
	}
	return upuo
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (upuo *UserProfileUpdateOne) ClearThumbnail() *UserProfileUpdateOne {
	upuo.mutation.ClearThumbnail()
	return upuo
}

// SetUpdatedAt sets the "updated_at" field.
func (upuo *UserProfileUpdateOne) SetUpdatedAt(t time.Time) *UserProfileUpdateOne {
	upuo.mutation.SetUpdatedAt(t)
	return upuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (upuo *UserProfileUpdateOne) SetUserID(id uuid.UUID) *UserProfileUpdateOne {
	upuo.mutation.SetUserID(id)
	return upuo
}

// SetUser sets the "user" edge to the User entity.
func (upuo *UserProfileUpdateOne) SetUser(u *User) *UserProfileUpdateOne {
	return upuo.SetUserID(u.ID)
}

// Mutation returns the UserProfileMutation object of the builder.
func (upuo *UserProfileUpdateOne) Mutation() *UserProfileMutation {
	return upuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (upuo *UserProfileUpdateOne) ClearUser() *UserProfileUpdateOne {
	upuo.mutation.ClearUser()
	return upuo
}

// Save executes the query and returns the updated UserProfile entity.
func (upuo *UserProfileUpdateOne) Save(ctx context.Context) (*UserProfile, error) {
	var (
		err  error
		node *UserProfile
	)
	upuo.defaults()
	if len(upuo.hooks) == 0 {
		if err = upuo.check(); err != nil {
			return nil, err
		}
		node, err = upuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upuo.check(); err != nil {
				return nil, err
			}
			upuo.mutation = mutation
			node, err = upuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(upuo.hooks) - 1; i >= 0; i-- {
			mut = upuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (upuo *UserProfileUpdateOne) SaveX(ctx context.Context) *UserProfile {
	node, err := upuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upuo *UserProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := upuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upuo *UserProfileUpdateOne) ExecX(ctx context.Context) {
	if err := upuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upuo *UserProfileUpdateOne) defaults() {
	if _, ok := upuo.mutation.UpdatedAt(); !ok {
		v := userprofile.UpdateDefaultUpdatedAt()
		upuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upuo *UserProfileUpdateOne) check() error {
	if v, ok := upuo.mutation.DisplayName(); ok {
		if err := userprofile.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf("ent: validator failed for field \"display_name\": %w", err)}
		}
	}
	if v, ok := upuo.mutation.ShortBio(); ok {
		if err := userprofile.ShortBioValidator(v); err != nil {
			return &ValidationError{Name: "short_bio", err: fmt.Errorf("ent: validator failed for field \"short_bio\": %w", err)}
		}
	}
	if _, ok := upuo.mutation.UserID(); upuo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (upuo *UserProfileUpdateOne) sqlSave(ctx context.Context) (_node *UserProfile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userprofile.Table,
			Columns: userprofile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userprofile.FieldID,
			},
		},
	}
	id, ok := upuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserProfile.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := upuo.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldDisplayName,
		})
	}
	if value, ok := upuo.mutation.ShortBio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldShortBio,
		})
	}
	if value, ok := upuo.mutation.About(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldAbout,
		})
	}
	if upuo.mutation.AboutCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldAbout,
		})
	}
	if value, ok := upuo.mutation.ProfileLinks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: userprofile.FieldProfileLinks,
		})
	}
	if upuo.mutation.ProfileLinksCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: userprofile.FieldProfileLinks,
		})
	}
	if value, ok := upuo.mutation.Thumbnail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldThumbnail,
		})
	}
	if upuo.mutation.ThumbnailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldThumbnail,
		})
	}
	if value, ok := upuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userprofile.FieldUpdatedAt,
		})
	}
	if upuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.UserTable,
			Columns: []string{userprofile.UserColumn},
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
	if nodes := upuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.UserTable,
			Columns: []string{userprofile.UserColumn},
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
	_node = &UserProfile{config: upuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprofile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
