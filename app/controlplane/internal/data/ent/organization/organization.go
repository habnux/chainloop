// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeMemberships holds the string denoting the memberships edge name in mutations.
	EdgeMemberships = "memberships"
	// EdgeWorkflowContracts holds the string denoting the workflow_contracts edge name in mutations.
	EdgeWorkflowContracts = "workflow_contracts"
	// EdgeWorkflows holds the string denoting the workflows edge name in mutations.
	EdgeWorkflows = "workflows"
	// EdgeCasBackends holds the string denoting the cas_backends edge name in mutations.
	EdgeCasBackends = "cas_backends"
	// EdgeIntegrations holds the string denoting the integrations edge name in mutations.
	EdgeIntegrations = "integrations"
	// Table holds the table name of the organization in the database.
	Table = "organizations"
	// MembershipsTable is the table that holds the memberships relation/edge.
	MembershipsTable = "memberships"
	// MembershipsInverseTable is the table name for the Membership entity.
	// It exists in this package in order to avoid circular dependency with the "membership" package.
	MembershipsInverseTable = "memberships"
	// MembershipsColumn is the table column denoting the memberships relation/edge.
	MembershipsColumn = "organization_memberships"
	// WorkflowContractsTable is the table that holds the workflow_contracts relation/edge.
	WorkflowContractsTable = "workflow_contracts"
	// WorkflowContractsInverseTable is the table name for the WorkflowContract entity.
	// It exists in this package in order to avoid circular dependency with the "workflowcontract" package.
	WorkflowContractsInverseTable = "workflow_contracts"
	// WorkflowContractsColumn is the table column denoting the workflow_contracts relation/edge.
	WorkflowContractsColumn = "organization_workflow_contracts"
	// WorkflowsTable is the table that holds the workflows relation/edge.
	WorkflowsTable = "workflows"
	// WorkflowsInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowsInverseTable = "workflows"
	// WorkflowsColumn is the table column denoting the workflows relation/edge.
	WorkflowsColumn = "organization_id"
	// CasBackendsTable is the table that holds the cas_backends relation/edge.
	CasBackendsTable = "cas_backends"
	// CasBackendsInverseTable is the table name for the CASBackend entity.
	// It exists in this package in order to avoid circular dependency with the "casbackend" package.
	CasBackendsInverseTable = "cas_backends"
	// CasBackendsColumn is the table column denoting the cas_backends relation/edge.
	CasBackendsColumn = "organization_cas_backends"
	// IntegrationsTable is the table that holds the integrations relation/edge.
	IntegrationsTable = "integrations"
	// IntegrationsInverseTable is the table name for the Integration entity.
	// It exists in this package in order to avoid circular dependency with the "integration" package.
	IntegrationsInverseTable = "integrations"
	// IntegrationsColumn is the table column denoting the integrations relation/edge.
	IntegrationsColumn = "organization_integrations"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Organization queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByMembershipsCount orders the results by memberships count.
func ByMembershipsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMembershipsStep(), opts...)
	}
}

// ByMemberships orders the results by memberships terms.
func ByMemberships(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMembershipsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkflowContractsCount orders the results by workflow_contracts count.
func ByWorkflowContractsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowContractsStep(), opts...)
	}
}

// ByWorkflowContracts orders the results by workflow_contracts terms.
func ByWorkflowContracts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowContractsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkflowsCount orders the results by workflows count.
func ByWorkflowsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowsStep(), opts...)
	}
}

// ByWorkflows orders the results by workflows terms.
func ByWorkflows(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCasBackendsCount orders the results by cas_backends count.
func ByCasBackendsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCasBackendsStep(), opts...)
	}
}

// ByCasBackends orders the results by cas_backends terms.
func ByCasBackends(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCasBackendsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIntegrationsCount orders the results by integrations count.
func ByIntegrationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIntegrationsStep(), opts...)
	}
}

// ByIntegrations orders the results by integrations terms.
func ByIntegrations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIntegrationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMembershipsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MembershipsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MembershipsTable, MembershipsColumn),
	)
}
func newWorkflowContractsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowContractsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkflowContractsTable, WorkflowContractsColumn),
	)
}
func newWorkflowsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkflowsTable, WorkflowsColumn),
	)
}
func newCasBackendsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CasBackendsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CasBackendsTable, CasBackendsColumn),
	)
}
func newIntegrationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IntegrationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IntegrationsTable, IntegrationsColumn),
	)
}
