// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/robotaccount"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflow"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowrun"
	"github.com/google/uuid"
)

// RobotAccountQuery is the builder for querying RobotAccount entities.
type RobotAccountQuery struct {
	config
	ctx              *QueryContext
	order            []OrderFunc
	inters           []Interceptor
	predicates       []predicate.RobotAccount
	withWorkflow     *WorkflowQuery
	withWorkflowruns *WorkflowRunQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RobotAccountQuery builder.
func (raq *RobotAccountQuery) Where(ps ...predicate.RobotAccount) *RobotAccountQuery {
	raq.predicates = append(raq.predicates, ps...)
	return raq
}

// Limit the number of records to be returned by this query.
func (raq *RobotAccountQuery) Limit(limit int) *RobotAccountQuery {
	raq.ctx.Limit = &limit
	return raq
}

// Offset to start from.
func (raq *RobotAccountQuery) Offset(offset int) *RobotAccountQuery {
	raq.ctx.Offset = &offset
	return raq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (raq *RobotAccountQuery) Unique(unique bool) *RobotAccountQuery {
	raq.ctx.Unique = &unique
	return raq
}

// Order specifies how the records should be ordered.
func (raq *RobotAccountQuery) Order(o ...OrderFunc) *RobotAccountQuery {
	raq.order = append(raq.order, o...)
	return raq
}

// QueryWorkflow chains the current query on the "workflow" edge.
func (raq *RobotAccountQuery) QueryWorkflow() *WorkflowQuery {
	query := (&WorkflowClient{config: raq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(robotaccount.Table, robotaccount.FieldID, selector),
			sqlgraph.To(workflow.Table, workflow.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, robotaccount.WorkflowTable, robotaccount.WorkflowColumn),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkflowruns chains the current query on the "workflowruns" edge.
func (raq *RobotAccountQuery) QueryWorkflowruns() *WorkflowRunQuery {
	query := (&WorkflowRunClient{config: raq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(robotaccount.Table, robotaccount.FieldID, selector),
			sqlgraph.To(workflowrun.Table, workflowrun.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, robotaccount.WorkflowrunsTable, robotaccount.WorkflowrunsColumn),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RobotAccount entity from the query.
// Returns a *NotFoundError when no RobotAccount was found.
func (raq *RobotAccountQuery) First(ctx context.Context) (*RobotAccount, error) {
	nodes, err := raq.Limit(1).All(setContextOp(ctx, raq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{robotaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (raq *RobotAccountQuery) FirstX(ctx context.Context) *RobotAccount {
	node, err := raq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RobotAccount ID from the query.
// Returns a *NotFoundError when no RobotAccount ID was found.
func (raq *RobotAccountQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = raq.Limit(1).IDs(setContextOp(ctx, raq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{robotaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (raq *RobotAccountQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := raq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RobotAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RobotAccount entity is found.
// Returns a *NotFoundError when no RobotAccount entities are found.
func (raq *RobotAccountQuery) Only(ctx context.Context) (*RobotAccount, error) {
	nodes, err := raq.Limit(2).All(setContextOp(ctx, raq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{robotaccount.Label}
	default:
		return nil, &NotSingularError{robotaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (raq *RobotAccountQuery) OnlyX(ctx context.Context) *RobotAccount {
	node, err := raq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RobotAccount ID in the query.
// Returns a *NotSingularError when more than one RobotAccount ID is found.
// Returns a *NotFoundError when no entities are found.
func (raq *RobotAccountQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = raq.Limit(2).IDs(setContextOp(ctx, raq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{robotaccount.Label}
	default:
		err = &NotSingularError{robotaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (raq *RobotAccountQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := raq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RobotAccounts.
func (raq *RobotAccountQuery) All(ctx context.Context) ([]*RobotAccount, error) {
	ctx = setContextOp(ctx, raq.ctx, "All")
	if err := raq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RobotAccount, *RobotAccountQuery]()
	return withInterceptors[[]*RobotAccount](ctx, raq, qr, raq.inters)
}

// AllX is like All, but panics if an error occurs.
func (raq *RobotAccountQuery) AllX(ctx context.Context) []*RobotAccount {
	nodes, err := raq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RobotAccount IDs.
func (raq *RobotAccountQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if raq.ctx.Unique == nil && raq.path != nil {
		raq.Unique(true)
	}
	ctx = setContextOp(ctx, raq.ctx, "IDs")
	if err = raq.Select(robotaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (raq *RobotAccountQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := raq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (raq *RobotAccountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, raq.ctx, "Count")
	if err := raq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, raq, querierCount[*RobotAccountQuery](), raq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (raq *RobotAccountQuery) CountX(ctx context.Context) int {
	count, err := raq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (raq *RobotAccountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, raq.ctx, "Exist")
	switch _, err := raq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (raq *RobotAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := raq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RobotAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (raq *RobotAccountQuery) Clone() *RobotAccountQuery {
	if raq == nil {
		return nil
	}
	return &RobotAccountQuery{
		config:           raq.config,
		ctx:              raq.ctx.Clone(),
		order:            append([]OrderFunc{}, raq.order...),
		inters:           append([]Interceptor{}, raq.inters...),
		predicates:       append([]predicate.RobotAccount{}, raq.predicates...),
		withWorkflow:     raq.withWorkflow.Clone(),
		withWorkflowruns: raq.withWorkflowruns.Clone(),
		// clone intermediate query.
		sql:  raq.sql.Clone(),
		path: raq.path,
	}
}

// WithWorkflow tells the query-builder to eager-load the nodes that are connected to
// the "workflow" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *RobotAccountQuery) WithWorkflow(opts ...func(*WorkflowQuery)) *RobotAccountQuery {
	query := (&WorkflowClient{config: raq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	raq.withWorkflow = query
	return raq
}

// WithWorkflowruns tells the query-builder to eager-load the nodes that are connected to
// the "workflowruns" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *RobotAccountQuery) WithWorkflowruns(opts ...func(*WorkflowRunQuery)) *RobotAccountQuery {
	query := (&WorkflowRunClient{config: raq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	raq.withWorkflowruns = query
	return raq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RobotAccount.Query().
//		GroupBy(robotaccount.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (raq *RobotAccountQuery) GroupBy(field string, fields ...string) *RobotAccountGroupBy {
	raq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RobotAccountGroupBy{build: raq}
	grbuild.flds = &raq.ctx.Fields
	grbuild.label = robotaccount.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.RobotAccount.Query().
//		Select(robotaccount.FieldName).
//		Scan(ctx, &v)
func (raq *RobotAccountQuery) Select(fields ...string) *RobotAccountSelect {
	raq.ctx.Fields = append(raq.ctx.Fields, fields...)
	sbuild := &RobotAccountSelect{RobotAccountQuery: raq}
	sbuild.label = robotaccount.Label
	sbuild.flds, sbuild.scan = &raq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RobotAccountSelect configured with the given aggregations.
func (raq *RobotAccountQuery) Aggregate(fns ...AggregateFunc) *RobotAccountSelect {
	return raq.Select().Aggregate(fns...)
}

func (raq *RobotAccountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range raq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, raq); err != nil {
				return err
			}
		}
	}
	for _, f := range raq.ctx.Fields {
		if !robotaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if raq.path != nil {
		prev, err := raq.path(ctx)
		if err != nil {
			return err
		}
		raq.sql = prev
	}
	return nil
}

func (raq *RobotAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RobotAccount, error) {
	var (
		nodes       = []*RobotAccount{}
		withFKs     = raq.withFKs
		_spec       = raq.querySpec()
		loadedTypes = [2]bool{
			raq.withWorkflow != nil,
			raq.withWorkflowruns != nil,
		}
	)
	if raq.withWorkflow != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, robotaccount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RobotAccount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RobotAccount{config: raq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, raq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := raq.withWorkflow; query != nil {
		if err := raq.loadWorkflow(ctx, query, nodes, nil,
			func(n *RobotAccount, e *Workflow) { n.Edges.Workflow = e }); err != nil {
			return nil, err
		}
	}
	if query := raq.withWorkflowruns; query != nil {
		if err := raq.loadWorkflowruns(ctx, query, nodes,
			func(n *RobotAccount) { n.Edges.Workflowruns = []*WorkflowRun{} },
			func(n *RobotAccount, e *WorkflowRun) { n.Edges.Workflowruns = append(n.Edges.Workflowruns, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (raq *RobotAccountQuery) loadWorkflow(ctx context.Context, query *WorkflowQuery, nodes []*RobotAccount, init func(*RobotAccount), assign func(*RobotAccount, *Workflow)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*RobotAccount)
	for i := range nodes {
		if nodes[i].workflow_robotaccounts == nil {
			continue
		}
		fk := *nodes[i].workflow_robotaccounts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workflow.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workflow_robotaccounts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (raq *RobotAccountQuery) loadWorkflowruns(ctx context.Context, query *WorkflowRunQuery, nodes []*RobotAccount, init func(*RobotAccount), assign func(*RobotAccount, *WorkflowRun)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*RobotAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.WorkflowRun(func(s *sql.Selector) {
		s.Where(sql.InValues(robotaccount.WorkflowrunsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.robot_account_workflowruns
		if fk == nil {
			return fmt.Errorf(`foreign-key "robot_account_workflowruns" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "robot_account_workflowruns" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (raq *RobotAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := raq.querySpec()
	_spec.Node.Columns = raq.ctx.Fields
	if len(raq.ctx.Fields) > 0 {
		_spec.Unique = raq.ctx.Unique != nil && *raq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, raq.driver, _spec)
}

func (raq *RobotAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(robotaccount.Table, robotaccount.Columns, sqlgraph.NewFieldSpec(robotaccount.FieldID, field.TypeUUID))
	_spec.From = raq.sql
	if unique := raq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if raq.path != nil {
		_spec.Unique = true
	}
	if fields := raq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, robotaccount.FieldID)
		for i := range fields {
			if fields[i] != robotaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := raq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := raq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := raq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := raq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (raq *RobotAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(raq.driver.Dialect())
	t1 := builder.Table(robotaccount.Table)
	columns := raq.ctx.Fields
	if len(columns) == 0 {
		columns = robotaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if raq.sql != nil {
		selector = raq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if raq.ctx.Unique != nil && *raq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range raq.predicates {
		p(selector)
	}
	for _, p := range raq.order {
		p(selector)
	}
	if offset := raq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := raq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RobotAccountGroupBy is the group-by builder for RobotAccount entities.
type RobotAccountGroupBy struct {
	selector
	build *RobotAccountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ragb *RobotAccountGroupBy) Aggregate(fns ...AggregateFunc) *RobotAccountGroupBy {
	ragb.fns = append(ragb.fns, fns...)
	return ragb
}

// Scan applies the selector query and scans the result into the given value.
func (ragb *RobotAccountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ragb.build.ctx, "GroupBy")
	if err := ragb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RobotAccountQuery, *RobotAccountGroupBy](ctx, ragb.build, ragb, ragb.build.inters, v)
}

func (ragb *RobotAccountGroupBy) sqlScan(ctx context.Context, root *RobotAccountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ragb.fns))
	for _, fn := range ragb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ragb.flds)+len(ragb.fns))
		for _, f := range *ragb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ragb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ragb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RobotAccountSelect is the builder for selecting fields of RobotAccount entities.
type RobotAccountSelect struct {
	*RobotAccountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ras *RobotAccountSelect) Aggregate(fns ...AggregateFunc) *RobotAccountSelect {
	ras.fns = append(ras.fns, fns...)
	return ras
}

// Scan applies the selector query and scans the result into the given value.
func (ras *RobotAccountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ras.ctx, "Select")
	if err := ras.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RobotAccountQuery, *RobotAccountSelect](ctx, ras.RobotAccountQuery, ras, ras.inters, v)
}

func (ras *RobotAccountSelect) sqlScan(ctx context.Context, root *RobotAccountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ras.fns))
	for _, fn := range ras.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ras.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ras.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
