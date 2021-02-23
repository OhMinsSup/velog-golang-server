// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/OhMinsSup/story-server/ent/predicate"
	"github.com/OhMinsSup/story-server/ent/user"
	"github.com/OhMinsSup/story-server/ent/velogconfig"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// VelogConfigQuery is the builder for querying VelogConfig entities.
type VelogConfigQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.VelogConfig
	// eager-loading edges.
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VelogConfigQuery builder.
func (vcq *VelogConfigQuery) Where(ps ...predicate.VelogConfig) *VelogConfigQuery {
	vcq.predicates = append(vcq.predicates, ps...)
	return vcq
}

// Limit adds a limit step to the query.
func (vcq *VelogConfigQuery) Limit(limit int) *VelogConfigQuery {
	vcq.limit = &limit
	return vcq
}

// Offset adds an offset step to the query.
func (vcq *VelogConfigQuery) Offset(offset int) *VelogConfigQuery {
	vcq.offset = &offset
	return vcq
}

// Order adds an order step to the query.
func (vcq *VelogConfigQuery) Order(o ...OrderFunc) *VelogConfigQuery {
	vcq.order = append(vcq.order, o...)
	return vcq
}

// QueryUser chains the current query on the "user" edge.
func (vcq *VelogConfigQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: vcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vcq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(velogconfig.Table, velogconfig.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, velogconfig.UserTable, velogconfig.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(vcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VelogConfig entity from the query.
// Returns a *NotFoundError when no VelogConfig was found.
func (vcq *VelogConfigQuery) First(ctx context.Context) (*VelogConfig, error) {
	nodes, err := vcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{velogconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vcq *VelogConfigQuery) FirstX(ctx context.Context) *VelogConfig {
	node, err := vcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VelogConfig ID from the query.
// Returns a *NotFoundError when no VelogConfig ID was found.
func (vcq *VelogConfigQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{velogconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vcq *VelogConfigQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VelogConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one VelogConfig entity is not found.
// Returns a *NotFoundError when no VelogConfig entities are found.
func (vcq *VelogConfigQuery) Only(ctx context.Context) (*VelogConfig, error) {
	nodes, err := vcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{velogconfig.Label}
	default:
		return nil, &NotSingularError{velogconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vcq *VelogConfigQuery) OnlyX(ctx context.Context) *VelogConfig {
	node, err := vcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VelogConfig ID in the query.
// Returns a *NotSingularError when exactly one VelogConfig ID is not found.
// Returns a *NotFoundError when no entities are found.
func (vcq *VelogConfigQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{velogconfig.Label}
	default:
		err = &NotSingularError{velogconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vcq *VelogConfigQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VelogConfigs.
func (vcq *VelogConfigQuery) All(ctx context.Context) ([]*VelogConfig, error) {
	if err := vcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vcq *VelogConfigQuery) AllX(ctx context.Context) []*VelogConfig {
	nodes, err := vcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VelogConfig IDs.
func (vcq *VelogConfigQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := vcq.Select(velogconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vcq *VelogConfigQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vcq *VelogConfigQuery) Count(ctx context.Context) (int, error) {
	if err := vcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vcq *VelogConfigQuery) CountX(ctx context.Context) int {
	count, err := vcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vcq *VelogConfigQuery) Exist(ctx context.Context) (bool, error) {
	if err := vcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vcq *VelogConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := vcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VelogConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vcq *VelogConfigQuery) Clone() *VelogConfigQuery {
	if vcq == nil {
		return nil
	}
	return &VelogConfigQuery{
		config:     vcq.config,
		limit:      vcq.limit,
		offset:     vcq.offset,
		order:      append([]OrderFunc{}, vcq.order...),
		predicates: append([]predicate.VelogConfig{}, vcq.predicates...),
		withUser:   vcq.withUser.Clone(),
		// clone intermediate query.
		sql:  vcq.sql.Clone(),
		path: vcq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (vcq *VelogConfigQuery) WithUser(opts ...func(*UserQuery)) *VelogConfigQuery {
	query := &UserQuery{config: vcq.config}
	for _, opt := range opts {
		opt(query)
	}
	vcq.withUser = query
	return vcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VelogConfig.Query().
//		GroupBy(velogconfig.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vcq *VelogConfigQuery) GroupBy(field string, fields ...string) *VelogConfigGroupBy {
	group := &VelogConfigGroupBy{config: vcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vcq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.VelogConfig.Query().
//		Select(velogconfig.FieldTitle).
//		Scan(ctx, &v)
//
func (vcq *VelogConfigQuery) Select(field string, fields ...string) *VelogConfigSelect {
	vcq.fields = append([]string{field}, fields...)
	return &VelogConfigSelect{VelogConfigQuery: vcq}
}

func (vcq *VelogConfigQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vcq.fields {
		if !velogconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vcq.path != nil {
		prev, err := vcq.path(ctx)
		if err != nil {
			return err
		}
		vcq.sql = prev
	}
	return nil
}

func (vcq *VelogConfigQuery) sqlAll(ctx context.Context) ([]*VelogConfig, error) {
	var (
		nodes       = []*VelogConfig{}
		withFKs     = vcq.withFKs
		_spec       = vcq.querySpec()
		loadedTypes = [1]bool{
			vcq.withUser != nil,
		}
	)
	if vcq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, velogconfig.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &VelogConfig{config: vcq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, vcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := vcq.withUser; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*VelogConfig)
		for i := range nodes {
			if fk := nodes[i].fk_user_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "fk_user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (vcq *VelogConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vcq.querySpec()
	return sqlgraph.CountNodes(ctx, vcq.driver, _spec)
}

func (vcq *VelogConfigQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (vcq *VelogConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   velogconfig.Table,
			Columns: velogconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: velogconfig.FieldID,
			},
		},
		From:   vcq.sql,
		Unique: true,
	}
	if fields := vcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, velogconfig.FieldID)
		for i := range fields {
			if fields[i] != velogconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, velogconfig.ValidColumn)
			}
		}
	}
	return _spec
}

func (vcq *VelogConfigQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(vcq.driver.Dialect())
	t1 := builder.Table(velogconfig.Table)
	selector := builder.Select(t1.Columns(velogconfig.Columns...)...).From(t1)
	if vcq.sql != nil {
		selector = vcq.sql
		selector.Select(selector.Columns(velogconfig.Columns...)...)
	}
	for _, p := range vcq.predicates {
		p(selector)
	}
	for _, p := range vcq.order {
		p(selector, velogconfig.ValidColumn)
	}
	if offset := vcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VelogConfigGroupBy is the group-by builder for VelogConfig entities.
type VelogConfigGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vcgb *VelogConfigGroupBy) Aggregate(fns ...AggregateFunc) *VelogConfigGroupBy {
	vcgb.fns = append(vcgb.fns, fns...)
	return vcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vcgb *VelogConfigGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vcgb.path(ctx)
	if err != nil {
		return err
	}
	vcgb.sql = query
	return vcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vcgb *VelogConfigGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := vcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (vcgb *VelogConfigGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(vcgb.fields) > 1 {
		return nil, errors.New("ent: VelogConfigGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := vcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vcgb *VelogConfigGroupBy) StringsX(ctx context.Context) []string {
	v, err := vcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vcgb *VelogConfigGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{velogconfig.Label}
	default:
		err = fmt.Errorf("ent: VelogConfigGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vcgb *VelogConfigGroupBy) StringX(ctx context.Context) string {
	v, err := vcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (vcgb *VelogConfigGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(vcgb.fields) > 1 {
		return nil, errors.New("ent: VelogConfigGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := vcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vcgb *VelogConfigGroupBy) IntsX(ctx context.Context) []int {
	v, err := vcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vcgb *VelogConfigGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{velogconfig.Label}
	default:
		err = fmt.Errorf("ent: VelogConfigGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vcgb *VelogConfigGroupBy) IntX(ctx context.Context) int {
	v, err := vcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (vcgb *VelogConfigGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(vcgb.fields) > 1 {
		return nil, errors.New("ent: VelogConfigGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := vcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vcgb *VelogConfigGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := vcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vcgb *VelogConfigGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{velogconfig.Label}
	default:
		err = fmt.Errorf("ent: VelogConfigGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vcgb *VelogConfigGroupBy) Float64X(ctx context.Context) float64 {
	v, err := vcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (vcgb *VelogConfigGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(vcgb.fields) > 1 {
		return nil, errors.New("ent: VelogConfigGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := vcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vcgb *VelogConfigGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := vcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vcgb *VelogConfigGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{velogconfig.Label}
	default:
		err = fmt.Errorf("ent: VelogConfigGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vcgb *VelogConfigGroupBy) BoolX(ctx context.Context) bool {
	v, err := vcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vcgb *VelogConfigGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vcgb.fields {
		if !velogconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vcgb *VelogConfigGroupBy) sqlQuery() *sql.Selector {
	selector := vcgb.sql
	columns := make([]string, 0, len(vcgb.fields)+len(vcgb.fns))
	columns = append(columns, vcgb.fields...)
	for _, fn := range vcgb.fns {
		columns = append(columns, fn(selector, velogconfig.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(vcgb.fields...)
}

// VelogConfigSelect is the builder for selecting fields of VelogConfig entities.
type VelogConfigSelect struct {
	*VelogConfigQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vcs *VelogConfigSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vcs.prepareQuery(ctx); err != nil {
		return err
	}
	vcs.sql = vcs.VelogConfigQuery.sqlQuery()
	return vcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vcs *VelogConfigSelect) ScanX(ctx context.Context, v interface{}) {
	if err := vcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (vcs *VelogConfigSelect) Strings(ctx context.Context) ([]string, error) {
	if len(vcs.fields) > 1 {
		return nil, errors.New("ent: VelogConfigSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := vcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vcs *VelogConfigSelect) StringsX(ctx context.Context) []string {
	v, err := vcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (vcs *VelogConfigSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{velogconfig.Label}
	default:
		err = fmt.Errorf("ent: VelogConfigSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vcs *VelogConfigSelect) StringX(ctx context.Context) string {
	v, err := vcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (vcs *VelogConfigSelect) Ints(ctx context.Context) ([]int, error) {
	if len(vcs.fields) > 1 {
		return nil, errors.New("ent: VelogConfigSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := vcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vcs *VelogConfigSelect) IntsX(ctx context.Context) []int {
	v, err := vcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (vcs *VelogConfigSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{velogconfig.Label}
	default:
		err = fmt.Errorf("ent: VelogConfigSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vcs *VelogConfigSelect) IntX(ctx context.Context) int {
	v, err := vcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (vcs *VelogConfigSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(vcs.fields) > 1 {
		return nil, errors.New("ent: VelogConfigSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := vcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vcs *VelogConfigSelect) Float64sX(ctx context.Context) []float64 {
	v, err := vcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (vcs *VelogConfigSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{velogconfig.Label}
	default:
		err = fmt.Errorf("ent: VelogConfigSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vcs *VelogConfigSelect) Float64X(ctx context.Context) float64 {
	v, err := vcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (vcs *VelogConfigSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(vcs.fields) > 1 {
		return nil, errors.New("ent: VelogConfigSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := vcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vcs *VelogConfigSelect) BoolsX(ctx context.Context) []bool {
	v, err := vcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (vcs *VelogConfigSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{velogconfig.Label}
	default:
		err = fmt.Errorf("ent: VelogConfigSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vcs *VelogConfigSelect) BoolX(ctx context.Context) bool {
	v, err := vcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vcs *VelogConfigSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vcs.sqlQuery().Query()
	if err := vcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vcs *VelogConfigSelect) sqlQuery() sql.Querier {
	selector := vcs.sql
	selector.Select(selector.Columns(vcs.fields...)...)
	return selector
}