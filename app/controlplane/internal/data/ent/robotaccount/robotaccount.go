// Code generated by ent, DO NOT EDIT.

package robotaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the robotaccount type in the database.
	Label = "robot_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldRevokedAt holds the string denoting the revoked_at field in the database.
	FieldRevokedAt = "revoked_at"
	// EdgeWorkflow holds the string denoting the workflow edge name in mutations.
	EdgeWorkflow = "workflow"
	// Table holds the table name of the robotaccount in the database.
	Table = "robot_accounts"
	// WorkflowTable is the table that holds the workflow relation/edge.
	WorkflowTable = "robot_accounts"
	// WorkflowInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowInverseTable = "workflows"
	// WorkflowColumn is the table column denoting the workflow relation/edge.
	WorkflowColumn = "workflow_robotaccounts"
)

// Columns holds all SQL columns for robotaccount fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
	FieldRevokedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "robot_accounts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"workflow_robotaccounts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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

// OrderOption defines the ordering options for the RobotAccount queries.
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

// ByRevokedAt orders the results by the revoked_at field.
func ByRevokedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRevokedAt, opts...).ToFunc()
}

// ByWorkflowField orders the results by workflow field.
func ByWorkflowField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowStep(), sql.OrderByField(field, opts...))
	}
}
func newWorkflowStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
	)
}
