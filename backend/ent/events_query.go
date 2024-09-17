// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/ai_themes"
	"github.com/Hosi121/SpeakUp/ent/event_records"
	"github.com/Hosi121/SpeakUp/ent/events"
	"github.com/Hosi121/SpeakUp/ent/predicate"
)

// EVENTSQuery is the builder for querying EVENTS entities.
type EVENTSQuery struct {
	config
	ctx              *QueryContext
	order            []events.OrderOption
	inters           []Interceptor
	predicates       []predicate.EVENTS
	withParticipated *EVENTRECORDSQuery
	withUses         *AITHEMESQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EVENTSQuery builder.
func (eq *EVENTSQuery) Where(ps ...predicate.EVENTS) *EVENTSQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EVENTSQuery) Limit(limit int) *EVENTSQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EVENTSQuery) Offset(offset int) *EVENTSQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EVENTSQuery) Unique(unique bool) *EVENTSQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EVENTSQuery) Order(o ...events.OrderOption) *EVENTSQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryParticipated chains the current query on the "participated" edge.
func (eq *EVENTSQuery) QueryParticipated() *EVENTRECORDSQuery {
	query := (&EVENTRECORDSClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(events.Table, events.FieldID, selector),
			sqlgraph.To(event_records.Table, event_records.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, events.ParticipatedTable, events.ParticipatedColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUses chains the current query on the "uses" edge.
func (eq *EVENTSQuery) QueryUses() *AITHEMESQuery {
	query := (&AITHEMESClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(events.Table, events.FieldID, selector),
			sqlgraph.To(ai_themes.Table, ai_themes.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, events.UsesTable, events.UsesColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EVENTS entity from the query.
// Returns a *NotFoundError when no EVENTS was found.
func (eq *EVENTSQuery) First(ctx context.Context) (*EVENTS, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{events.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EVENTSQuery) FirstX(ctx context.Context) *EVENTS {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EVENTS ID from the query.
// Returns a *NotFoundError when no EVENTS ID was found.
func (eq *EVENTSQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{events.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EVENTSQuery) FirstIDX(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EVENTS entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EVENTS entity is found.
// Returns a *NotFoundError when no EVENTS entities are found.
func (eq *EVENTSQuery) Only(ctx context.Context) (*EVENTS, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{events.Label}
	default:
		return nil, &NotSingularError{events.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EVENTSQuery) OnlyX(ctx context.Context) *EVENTS {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EVENTS ID in the query.
// Returns a *NotSingularError when more than one EVENTS ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EVENTSQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{events.Label}
	default:
		err = &NotSingularError{events.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EVENTSQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EVENTSs.
func (eq *EVENTSQuery) All(ctx context.Context) ([]*EVENTS, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryAll)
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EVENTS, *EVENTSQuery]()
	return withInterceptors[[]*EVENTS](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EVENTSQuery) AllX(ctx context.Context) []*EVENTS {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EVENTS IDs.
func (eq *EVENTSQuery) IDs(ctx context.Context) (ids []int, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryIDs)
	if err = eq.Select(events.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EVENTSQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EVENTSQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryCount)
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EVENTSQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EVENTSQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EVENTSQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryExist)
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EVENTSQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EVENTSQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EVENTSQuery) Clone() *EVENTSQuery {
	if eq == nil {
		return nil
	}
	return &EVENTSQuery{
		config:           eq.config,
		ctx:              eq.ctx.Clone(),
		order:            append([]events.OrderOption{}, eq.order...),
		inters:           append([]Interceptor{}, eq.inters...),
		predicates:       append([]predicate.EVENTS{}, eq.predicates...),
		withParticipated: eq.withParticipated.Clone(),
		withUses:         eq.withUses.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithParticipated tells the query-builder to eager-load the nodes that are connected to
// the "participated" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EVENTSQuery) WithParticipated(opts ...func(*EVENTRECORDSQuery)) *EVENTSQuery {
	query := (&EVENTRECORDSClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withParticipated = query
	return eq
}

// WithUses tells the query-builder to eager-load the nodes that are connected to
// the "uses" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EVENTSQuery) WithUses(opts ...func(*AITHEMESQuery)) *EVENTSQuery {
	query := (&AITHEMESClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withUses = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EventStart time.Time `json:"event_start,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EVENTS.Query().
//		GroupBy(events.FieldEventStart).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *EVENTSQuery) GroupBy(field string, fields ...string) *EVENTSGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EVENTSGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = events.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EventStart time.Time `json:"event_start,omitempty"`
//	}
//
//	client.EVENTS.Query().
//		Select(events.FieldEventStart).
//		Scan(ctx, &v)
func (eq *EVENTSQuery) Select(fields ...string) *EVENTSSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EVENTSSelect{EVENTSQuery: eq}
	sbuild.label = events.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EVENTSSelect configured with the given aggregations.
func (eq *EVENTSQuery) Aggregate(fns ...AggregateFunc) *EVENTSSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EVENTSQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !events.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EVENTSQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EVENTS, error) {
	var (
		nodes       = []*EVENTS{}
		withFKs     = eq.withFKs
		_spec       = eq.querySpec()
		loadedTypes = [2]bool{
			eq.withParticipated != nil,
			eq.withUses != nil,
		}
	)
	if eq.withUses != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, events.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EVENTS).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EVENTS{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withParticipated; query != nil {
		if err := eq.loadParticipated(ctx, query, nodes,
			func(n *EVENTS) { n.Edges.Participated = []*EVENT_RECORDS{} },
			func(n *EVENTS, e *EVENT_RECORDS) { n.Edges.Participated = append(n.Edges.Participated, e) }); err != nil {
			return nil, err
		}
	}
	if query := eq.withUses; query != nil {
		if err := eq.loadUses(ctx, query, nodes, nil,
			func(n *EVENTS, e *AI_THEMES) { n.Edges.Uses = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EVENTSQuery) loadParticipated(ctx context.Context, query *EVENTRECORDSQuery, nodes []*EVENTS, init func(*EVENTS), assign func(*EVENTS, *EVENT_RECORDS)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*EVENTS)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.EVENT_RECORDS(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(events.ParticipatedColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.events_participated
		if fk == nil {
			return fmt.Errorf(`foreign-key "events_participated" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "events_participated" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eq *EVENTSQuery) loadUses(ctx context.Context, query *AITHEMESQuery, nodes []*EVENTS, init func(*EVENTS), assign func(*EVENTS, *AI_THEMES)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*EVENTS)
	for i := range nodes {
		if nodes[i].events_uses == nil {
			continue
		}
		fk := *nodes[i].events_uses
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(ai_themes.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "events_uses" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eq *EVENTSQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EVENTSQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(events.Table, events.Columns, sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, events.FieldID)
		for i := range fields {
			if fields[i] != events.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EVENTSQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(events.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = events.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EVENTSGroupBy is the group-by builder for EVENTS entities.
type EVENTSGroupBy struct {
	selector
	build *EVENTSQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EVENTSGroupBy) Aggregate(fns ...AggregateFunc) *EVENTSGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EVENTSGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, ent.OpQueryGroupBy)
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EVENTSQuery, *EVENTSGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EVENTSGroupBy) sqlScan(ctx context.Context, root *EVENTSQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EVENTSSelect is the builder for selecting fields of EVENTS entities.
type EVENTSSelect struct {
	*EVENTSQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EVENTSSelect) Aggregate(fns ...AggregateFunc) *EVENTSSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EVENTSSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, ent.OpQuerySelect)
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EVENTSQuery, *EVENTSSelect](ctx, es.EVENTSQuery, es, es.inters, v)
}

func (es *EVENTSSelect) sqlScan(ctx context.Context, root *EVENTSQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}