// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/OhMinsSup/story-server/ent/authtoken"
	"github.com/OhMinsSup/story-server/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// AuthTokenQuery is the builder for querying AuthToken entities.
type AuthTokenQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.AuthToken
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthTokenQuery builder.
func (atq *AuthTokenQuery) Where(ps ...predicate.AuthToken) *AuthTokenQuery {
	atq.predicates = append(atq.predicates, ps...)
	return atq
}

// Limit adds a limit step to the query.
func (atq *AuthTokenQuery) Limit(limit int) *AuthTokenQuery {
	atq.limit = &limit
	return atq
}

// Offset adds an offset step to the query.
func (atq *AuthTokenQuery) Offset(offset int) *AuthTokenQuery {
	atq.offset = &offset
	return atq
}

// Order adds an order step to the query.
func (atq *AuthTokenQuery) Order(o ...OrderFunc) *AuthTokenQuery {
	atq.order = append(atq.order, o...)
	return atq
}

// First returns the first AuthToken entity from the query.
// Returns a *NotFoundError when no AuthToken was found.
func (atq *AuthTokenQuery) First(ctx context.Context) (*AuthToken, error) {
	nodes, err := atq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authtoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atq *AuthTokenQuery) FirstX(ctx context.Context) *AuthToken {
	node, err := atq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthToken ID from the query.
// Returns a *NotFoundError when no AuthToken ID was found.
func (atq *AuthTokenQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authtoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atq *AuthTokenQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := atq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AuthToken entity is not found.
// Returns a *NotFoundError when no AuthToken entities are found.
func (atq *AuthTokenQuery) Only(ctx context.Context) (*AuthToken, error) {
	nodes, err := atq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authtoken.Label}
	default:
		return nil, &NotSingularError{authtoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atq *AuthTokenQuery) OnlyX(ctx context.Context) *AuthToken {
	node, err := atq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthToken ID in the query.
// Returns a *NotSingularError when exactly one AuthToken ID is not found.
// Returns a *NotFoundError when no entities are found.
func (atq *AuthTokenQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authtoken.Label}
	default:
		err = &NotSingularError{authtoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atq *AuthTokenQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := atq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthTokens.
func (atq *AuthTokenQuery) All(ctx context.Context) ([]*AuthToken, error) {
	if err := atq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return atq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (atq *AuthTokenQuery) AllX(ctx context.Context) []*AuthToken {
	nodes, err := atq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthToken IDs.
func (atq *AuthTokenQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := atq.Select(authtoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atq *AuthTokenQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := atq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atq *AuthTokenQuery) Count(ctx context.Context) (int, error) {
	if err := atq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return atq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (atq *AuthTokenQuery) CountX(ctx context.Context) int {
	count, err := atq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atq *AuthTokenQuery) Exist(ctx context.Context) (bool, error) {
	if err := atq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return atq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (atq *AuthTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := atq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atq *AuthTokenQuery) Clone() *AuthTokenQuery {
	if atq == nil {
		return nil
	}
	return &AuthTokenQuery{
		config:     atq.config,
		limit:      atq.limit,
		offset:     atq.offset,
		order:      append([]OrderFunc{}, atq.order...),
		predicates: append([]predicate.AuthToken{}, atq.predicates...),
		// clone intermediate query.
		sql:  atq.sql.Clone(),
		path: atq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Disabled bool `json:"disabled,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthToken.Query().
//		GroupBy(authtoken.FieldDisabled).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (atq *AuthTokenQuery) GroupBy(field string, fields ...string) *AuthTokenGroupBy {
	group := &AuthTokenGroupBy{config: atq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := atq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return atq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Disabled bool `json:"disabled,omitempty"`
//	}
//
//	client.AuthToken.Query().
//		Select(authtoken.FieldDisabled).
//		Scan(ctx, &v)
//
func (atq *AuthTokenQuery) Select(field string, fields ...string) *AuthTokenSelect {
	atq.fields = append([]string{field}, fields...)
	return &AuthTokenSelect{AuthTokenQuery: atq}
}

func (atq *AuthTokenQuery) prepareQuery(ctx context.Context) error {
	for _, f := range atq.fields {
		if !authtoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if atq.path != nil {
		prev, err := atq.path(ctx)
		if err != nil {
			return err
		}
		atq.sql = prev
	}
	return nil
}

func (atq *AuthTokenQuery) sqlAll(ctx context.Context) ([]*AuthToken, error) {
	var (
		nodes = []*AuthToken{}
		_spec = atq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AuthToken{config: atq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, atq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (atq *AuthTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atq.querySpec()
	return sqlgraph.CountNodes(ctx, atq.driver, _spec)
}

func (atq *AuthTokenQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := atq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (atq *AuthTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authtoken.Table,
			Columns: authtoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: authtoken.FieldID,
			},
		},
		From:   atq.sql,
		Unique: true,
	}
	if fields := atq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authtoken.FieldID)
		for i := range fields {
			if fields[i] != authtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := atq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, authtoken.ValidColumn)
			}
		}
	}
	return _spec
}

func (atq *AuthTokenQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(atq.driver.Dialect())
	t1 := builder.Table(authtoken.Table)
	selector := builder.Select(t1.Columns(authtoken.Columns...)...).From(t1)
	if atq.sql != nil {
		selector = atq.sql
		selector.Select(selector.Columns(authtoken.Columns...)...)
	}
	for _, p := range atq.predicates {
		p(selector)
	}
	for _, p := range atq.order {
		p(selector, authtoken.ValidColumn)
	}
	if offset := atq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthTokenGroupBy is the group-by builder for AuthToken entities.
type AuthTokenGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atgb *AuthTokenGroupBy) Aggregate(fns ...AggregateFunc) *AuthTokenGroupBy {
	atgb.fns = append(atgb.fns, fns...)
	return atgb
}

// Scan applies the group-by query and scans the result into the given value.
func (atgb *AuthTokenGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := atgb.path(ctx)
	if err != nil {
		return err
	}
	atgb.sql = query
	return atgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (atgb *AuthTokenGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := atgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (atgb *AuthTokenGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(atgb.fields) > 1 {
		return nil, errors.New("ent: AuthTokenGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := atgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (atgb *AuthTokenGroupBy) StringsX(ctx context.Context) []string {
	v, err := atgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atgb *AuthTokenGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = atgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authtoken.Label}
	default:
		err = fmt.Errorf("ent: AuthTokenGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (atgb *AuthTokenGroupBy) StringX(ctx context.Context) string {
	v, err := atgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (atgb *AuthTokenGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(atgb.fields) > 1 {
		return nil, errors.New("ent: AuthTokenGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := atgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (atgb *AuthTokenGroupBy) IntsX(ctx context.Context) []int {
	v, err := atgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atgb *AuthTokenGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = atgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authtoken.Label}
	default:
		err = fmt.Errorf("ent: AuthTokenGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (atgb *AuthTokenGroupBy) IntX(ctx context.Context) int {
	v, err := atgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (atgb *AuthTokenGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(atgb.fields) > 1 {
		return nil, errors.New("ent: AuthTokenGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := atgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (atgb *AuthTokenGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := atgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atgb *AuthTokenGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = atgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authtoken.Label}
	default:
		err = fmt.Errorf("ent: AuthTokenGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (atgb *AuthTokenGroupBy) Float64X(ctx context.Context) float64 {
	v, err := atgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (atgb *AuthTokenGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(atgb.fields) > 1 {
		return nil, errors.New("ent: AuthTokenGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := atgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (atgb *AuthTokenGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := atgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atgb *AuthTokenGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = atgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authtoken.Label}
	default:
		err = fmt.Errorf("ent: AuthTokenGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (atgb *AuthTokenGroupBy) BoolX(ctx context.Context) bool {
	v, err := atgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (atgb *AuthTokenGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range atgb.fields {
		if !authtoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := atgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (atgb *AuthTokenGroupBy) sqlQuery() *sql.Selector {
	selector := atgb.sql
	columns := make([]string, 0, len(atgb.fields)+len(atgb.fns))
	columns = append(columns, atgb.fields...)
	for _, fn := range atgb.fns {
		columns = append(columns, fn(selector, authtoken.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(atgb.fields...)
}

// AuthTokenSelect is the builder for selecting fields of AuthToken entities.
type AuthTokenSelect struct {
	*AuthTokenQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ats *AuthTokenSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ats.prepareQuery(ctx); err != nil {
		return err
	}
	ats.sql = ats.AuthTokenQuery.sqlQuery()
	return ats.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ats *AuthTokenSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ats.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ats *AuthTokenSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ats.fields) > 1 {
		return nil, errors.New("ent: AuthTokenSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ats *AuthTokenSelect) StringsX(ctx context.Context) []string {
	v, err := ats.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ats *AuthTokenSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ats.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authtoken.Label}
	default:
		err = fmt.Errorf("ent: AuthTokenSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ats *AuthTokenSelect) StringX(ctx context.Context) string {
	v, err := ats.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ats *AuthTokenSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ats.fields) > 1 {
		return nil, errors.New("ent: AuthTokenSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ats *AuthTokenSelect) IntsX(ctx context.Context) []int {
	v, err := ats.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ats *AuthTokenSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ats.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authtoken.Label}
	default:
		err = fmt.Errorf("ent: AuthTokenSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ats *AuthTokenSelect) IntX(ctx context.Context) int {
	v, err := ats.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ats *AuthTokenSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ats.fields) > 1 {
		return nil, errors.New("ent: AuthTokenSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ats *AuthTokenSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ats.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ats *AuthTokenSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ats.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authtoken.Label}
	default:
		err = fmt.Errorf("ent: AuthTokenSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ats *AuthTokenSelect) Float64X(ctx context.Context) float64 {
	v, err := ats.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ats *AuthTokenSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ats.fields) > 1 {
		return nil, errors.New("ent: AuthTokenSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ats *AuthTokenSelect) BoolsX(ctx context.Context) []bool {
	v, err := ats.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ats *AuthTokenSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ats.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authtoken.Label}
	default:
		err = fmt.Errorf("ent: AuthTokenSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ats *AuthTokenSelect) BoolX(ctx context.Context) bool {
	v, err := ats.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ats *AuthTokenSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ats.sqlQuery().Query()
	if err := ats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ats *AuthTokenSelect) sqlQuery() sql.Querier {
	selector := ats.sql
	selector.Select(selector.Columns(ats.fields...)...)
	return selector
}
