// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/OhMinsSup/story-server/ent/post"
	"github.com/OhMinsSup/story-server/ent/tag"
	"github.com/OhMinsSup/story-server/ent/user"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetBody sets the "body" field.
func (pc *PostCreate) SetBody(s string) *PostCreate {
	pc.mutation.SetBody(s)
	return pc
}

// SetThumbnail sets the "thumbnail" field.
func (pc *PostCreate) SetThumbnail(s string) *PostCreate {
	pc.mutation.SetThumbnail(s)
	return pc
}

// SetNillableThumbnail sets the "thumbnail" field if the given value is not nil.
func (pc *PostCreate) SetNillableThumbnail(s *string) *PostCreate {
	if s != nil {
		pc.SetThumbnail(*s)
	}
	return pc
}

// SetIsTemp sets the "is_temp" field.
func (pc *PostCreate) SetIsTemp(b bool) *PostCreate {
	pc.mutation.SetIsTemp(b)
	return pc
}

// SetIsMarkdown sets the "is_markdown" field.
func (pc *PostCreate) SetIsMarkdown(b bool) *PostCreate {
	pc.mutation.SetIsMarkdown(b)
	return pc
}

// SetIsPrivate sets the "is_private" field.
func (pc *PostCreate) SetIsPrivate(b bool) *PostCreate {
	pc.mutation.SetIsPrivate(b)
	return pc
}

// SetNillableIsPrivate sets the "is_private" field if the given value is not nil.
func (pc *PostCreate) SetNillableIsPrivate(b *bool) *PostCreate {
	if b != nil {
		pc.SetIsPrivate(*b)
	}
	return pc
}

// SetURLSlug sets the "url_slug" field.
func (pc *PostCreate) SetURLSlug(s string) *PostCreate {
	pc.mutation.SetURLSlug(s)
	return pc
}

// SetLikes sets the "likes" field.
func (pc *PostCreate) SetLikes(i int64) *PostCreate {
	pc.mutation.SetLikes(i)
	return pc
}

// SetNillableLikes sets the "likes" field if the given value is not nil.
func (pc *PostCreate) SetNillableLikes(i *int64) *PostCreate {
	if i != nil {
		pc.SetLikes(*i)
	}
	return pc
}

// SetViews sets the "views" field.
func (pc *PostCreate) SetViews(i int64) *PostCreate {
	pc.mutation.SetViews(i)
	return pc
}

// SetNillableViews sets the "views" field if the given value is not nil.
func (pc *PostCreate) SetNillableViews(i *int64) *PostCreate {
	if i != nil {
		pc.SetViews(*i)
	}
	return pc
}

// SetMeta sets the "meta" field.
func (pc *PostCreate) SetMeta(m map[string]interface{}) *PostCreate {
	pc.mutation.SetMeta(m)
	return pc
}

// SetReleasedAt sets the "released_at" field.
func (pc *PostCreate) SetReleasedAt(t time.Time) *PostCreate {
	pc.mutation.SetReleasedAt(t)
	return pc
}

// SetNillableReleasedAt sets the "released_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableReleasedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetReleasedAt(*t)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PostCreate) SetCreatedAt(t time.Time) *PostCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PostCreate) SetUpdatedAt(t time.Time) *PostCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableUpdatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PostCreate) SetID(u uuid.UUID) *PostCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *PostCreate) SetUserID(id uuid.UUID) *PostCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableUserID(id *uuid.UUID) *PostCreate {
	if id != nil {
		pc = pc.SetUserID(*id)
	}
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PostCreate) SetUser(u *User) *PostCreate {
	return pc.SetUserID(u.ID)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (pc *PostCreate) AddTagIDs(ids ...uuid.UUID) *PostCreate {
	pc.mutation.AddTagIDs(ids...)
	return pc
}

// AddTags adds the "tags" edges to the Tag entity.
func (pc *PostCreate) AddTags(t ...*Tag) *PostCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTagIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	var (
		err  error
		node *Post
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.IsPrivate(); !ok {
		v := post.DefaultIsPrivate
		pc.mutation.SetIsPrivate(v)
	}
	if _, ok := pc.mutation.Likes(); !ok {
		v := post.DefaultLikes
		pc.mutation.SetLikes(v)
	}
	if _, ok := pc.mutation.Views(); !ok {
		v := post.DefaultViews
		pc.mutation.SetViews(v)
	}
	if _, ok := pc.mutation.ReleasedAt(); !ok {
		v := post.DefaultReleasedAt()
		pc.mutation.SetReleasedAt(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := post.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := post.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := post.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if v, ok := pc.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New("ent: missing required field \"body\"")}
	}
	if v, ok := pc.mutation.Thumbnail(); ok {
		if err := post.ThumbnailValidator(v); err != nil {
			return &ValidationError{Name: "thumbnail", err: fmt.Errorf("ent: validator failed for field \"thumbnail\": %w", err)}
		}
	}
	if _, ok := pc.mutation.IsTemp(); !ok {
		return &ValidationError{Name: "is_temp", err: errors.New("ent: missing required field \"is_temp\"")}
	}
	if _, ok := pc.mutation.IsMarkdown(); !ok {
		return &ValidationError{Name: "is_markdown", err: errors.New("ent: missing required field \"is_markdown\"")}
	}
	if _, ok := pc.mutation.IsPrivate(); !ok {
		return &ValidationError{Name: "is_private", err: errors.New("ent: missing required field \"is_private\"")}
	}
	if _, ok := pc.mutation.URLSlug(); !ok {
		return &ValidationError{Name: "url_slug", err: errors.New("ent: missing required field \"url_slug\"")}
	}
	if v, ok := pc.mutation.URLSlug(); ok {
		if err := post.URLSlugValidator(v); err != nil {
			return &ValidationError{Name: "url_slug", err: fmt.Errorf("ent: validator failed for field \"url_slug\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Likes(); !ok {
		return &ValidationError{Name: "likes", err: errors.New("ent: missing required field \"likes\"")}
	}
	if _, ok := pc.mutation.Views(); !ok {
		return &ValidationError{Name: "views", err: errors.New("ent: missing required field \"views\"")}
	}
	if _, ok := pc.mutation.Meta(); !ok {
		return &ValidationError{Name: "meta", err: errors.New("ent: missing required field \"meta\"")}
	}
	if _, ok := pc.mutation.ReleasedAt(); !ok {
		return &ValidationError{Name: "released_at", err: errors.New("ent: missing required field \"released_at\"")}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: post.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: post.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := pc.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldBody,
		})
		_node.Body = value
	}
	if value, ok := pc.mutation.Thumbnail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldThumbnail,
		})
		_node.Thumbnail = &value
	}
	if value, ok := pc.mutation.IsTemp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: post.FieldIsTemp,
		})
		_node.IsTemp = value
	}
	if value, ok := pc.mutation.IsMarkdown(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: post.FieldIsMarkdown,
		})
		_node.IsMarkdown = value
	}
	if value, ok := pc.mutation.IsPrivate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: post.FieldIsPrivate,
		})
		_node.IsPrivate = value
	}
	if value, ok := pc.mutation.URLSlug(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldURLSlug,
		})
		_node.URLSlug = value
	}
	if value, ok := pc.mutation.Likes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldLikes,
		})
		_node.Likes = value
	}
	if value, ok := pc.mutation.Views(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldViews,
		})
		_node.Views = value
	}
	if value, ok := pc.mutation.Meta(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: post.FieldMeta,
		})
		_node.Meta = value
	}
	if value, ok := pc.mutation.ReleasedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldReleasedAt,
		})
		_node.ReleasedAt = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.UserTable,
			Columns: []string{post.UserColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.TagsTable,
			Columns: post.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	builders []*PostCreate
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}