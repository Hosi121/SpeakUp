// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/achievements"
	"github.com/Hosi121/SpeakUp/ent/predicate"
	"github.com/Hosi121/SpeakUp/ent/users"
)

// ACHIEVEMENTSQuery is the builder for querying ACHIEVEMENTS entities.
type ACHIEVEMENTSQuery struct {
	config
	ctx         *QueryContext
	order       []achievements.OrderOption
	inters      []Interceptor
	predicates  []predicate.ACHIEVEMENTS
	withGranted *USERSQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ACHIEVEMENTSQuery builder.
func (aq *ACHIEVEMENTSQuery) Where(ps ...predicate.ACHIEVEMENTS) *ACHIEVEMENTSQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *ACHIEVEMENTSQuery) Limit(limit int) *ACHIEVEMENTSQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *ACHIEVEMENTSQuery) Offset(offset int) *ACHIEVEMENTSQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *ACHIEVEMENTSQuery) Unique(unique bool) *ACHIEVEMENTSQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *ACHIEVEMENTSQuery) Order(o ...achievements.OrderOption) *ACHIEVEMENTSQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryGranted chains the current query on the "granted" edge.
func (aq *ACHIEVEMENTSQuery) QueryGranted() *USERSQuery {
	query := (&USERSClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(achievements.Table, achievements.FieldID, selector),
			sqlgraph.To(users.Table, users.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, achievements.GrantedTable, achievements.GrantedColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ACHIEVEMENTS entity from the query.
// Returns a *NotFoundError when no ACHIEVEMENTS was found.
func (aq *ACHIEVEMENTSQuery) First(ctx context.Context) (*ACHIEVEMENTS, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{achievements.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *ACHIEVEMENTSQuery) FirstX(ctx context.Context) *ACHIEVEMENTS {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ACHIEVEMENTS ID from the query.
// Returns a *NotFoundError when no ACHIEVEMENTS ID was found.
func (aq *ACHIEVEMENTSQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{achievements.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *ACHIEVEMENTSQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ACHIEVEMENTS entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ACHIEVEMENTS entity is found.
// Returns a *NotFoundError when no ACHIEVEMENTS entities are found.
func (aq *ACHIEVEMENTSQuery) Only(ctx context.Context) (*ACHIEVEMENTS, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{achievements.Label}
	default:
		return nil, &NotSingularError{achievements.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *ACHIEVEMENTSQuery) OnlyX(ctx context.Context) *ACHIEVEMENTS {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ACHIEVEMENTS ID in the query.
// Returns a *NotSingularError when more than one ACHIEVEMENTS ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *ACHIEVEMENTSQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{achievements.Label}
	default:
		err = &NotSingularError{achievements.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *ACHIEVEMENTSQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ACHIEVEMENTSs.
func (aq *ACHIEVEMENTSQuery) All(ctx context.Context) ([]*ACHIEVEMENTS, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryAll)
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ACHIEVEMENTS, *ACHIEVEMENTSQuery]()
	return withInterceptors[[]*ACHIEVEMENTS](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *ACHIEVEMENTSQuery) AllX(ctx context.Context) []*ACHIEVEMENTS {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ACHIEVEMENTS IDs.
func (aq *ACHIEVEMENTSQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryIDs)
	if err = aq.Select(achievements.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *ACHIEVEMENTSQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *ACHIEVEMENTSQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryCount)
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*ACHIEVEMENTSQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *ACHIEVEMENTSQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *ACHIEVEMENTSQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryExist)
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *ACHIEVEMENTSQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ACHIEVEMENTSQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *ACHIEVEMENTSQuery) Clone() *ACHIEVEMENTSQuery {
	if aq == nil {
		return nil
	}
	return &ACHIEVEMENTSQuery{
		config:      aq.config,
		ctx:         aq.ctx.Clone(),
		order:       append([]achievements.OrderOption{}, aq.order...),
		inters:      append([]Interceptor{}, aq.inters...),
		predicates:  append([]predicate.ACHIEVEMENTS{}, aq.predicates...),
		withGranted: aq.withGranted.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithGranted tells the query-builder to eager-load the nodes that are connected to
// the "granted" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ACHIEVEMENTSQuery) WithGranted(opts ...func(*USERSQuery)) *ACHIEVEMENTSQuery {
	query := (&USERSClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withGranted = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ACHIEVEMENTS.Query().
//		GroupBy(achievements.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *ACHIEVEMENTSQuery) GroupBy(field string, fields ...string) *ACHIEVEMENTSGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ACHIEVEMENTSGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = achievements.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//	}
//
//	client.ACHIEVEMENTS.Query().
//		Select(achievements.FieldUserID).
//		Scan(ctx, &v)
func (aq *ACHIEVEMENTSQuery) Select(fields ...string) *ACHIEVEMENTSSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &ACHIEVEMENTSSelect{ACHIEVEMENTSQuery: aq}
	sbuild.label = achievements.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ACHIEVEMENTSSelect configured with the given aggregations.
func (aq *ACHIEVEMENTSQuery) Aggregate(fns ...AggregateFunc) *ACHIEVEMENTSSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *ACHIEVEMENTSQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !achievements.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *ACHIEVEMENTSQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ACHIEVEMENTS, error) {
	var (
		nodes       = []*ACHIEVEMENTS{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withGranted != nil,
		}
	)
	if aq.withGranted != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, achievements.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ACHIEVEMENTS).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ACHIEVEMENTS{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withGranted; query != nil {
		if err := aq.loadGranted(ctx, query, nodes, nil,
			func(n *ACHIEVEMENTS, e *USERS) { n.Edges.Granted = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *ACHIEVEMENTSQuery) loadGranted(ctx context.Context, query *USERSQuery, nodes []*ACHIEVEMENTS, init func(*ACHIEVEMENTS), assign func(*ACHIEVEMENTS, *USERS)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ACHIEVEMENTS)
	for i := range nodes {
		if nodes[i].users_acquires == nil {
			continue
		}
		fk := *nodes[i].users_acquires
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(users.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "users_acquires" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *ACHIEVEMENTSQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *ACHIEVEMENTSQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(achievements.Table, achievements.Columns, sqlgraph.NewFieldSpec(achievements.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, achievements.FieldID)
		for i := range fields {
			if fields[i] != achievements.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *ACHIEVEMENTSQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(achievements.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = achievements.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ACHIEVEMENTSGroupBy is the group-by builder for ACHIEVEMENTS entities.
type ACHIEVEMENTSGroupBy struct {
	selector
	build *ACHIEVEMENTSQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *ACHIEVEMENTSGroupBy) Aggregate(fns ...AggregateFunc) *ACHIEVEMENTSGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *ACHIEVEMENTSGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, ent.OpQueryGroupBy)
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ACHIEVEMENTSQuery, *ACHIEVEMENTSGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *ACHIEVEMENTSGroupBy) sqlScan(ctx context.Context, root *ACHIEVEMENTSQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ACHIEVEMENTSSelect is the builder for selecting fields of ACHIEVEMENTS entities.
type ACHIEVEMENTSSelect struct {
	*ACHIEVEMENTSQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *ACHIEVEMENTSSelect) Aggregate(fns ...AggregateFunc) *ACHIEVEMENTSSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *ACHIEVEMENTSSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, ent.OpQuerySelect)
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ACHIEVEMENTSQuery, *ACHIEVEMENTSSelect](ctx, as.ACHIEVEMENTSQuery, as, as.inters, v)
}

func (as *ACHIEVEMENTSSelect) sqlScan(ctx context.Context, root *ACHIEVEMENTSQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
