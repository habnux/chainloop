// Code generated by ent, DO NOT EDIT.

package workflowrun

import (
	"fmt"
	"time"

	"github.com/chainloop-dev/chainloop/app/controlplane/internal/biz"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the workflowrun type in the database.
	Label = "workflow_run"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldReason holds the string denoting the reason field in the database.
	FieldReason = "reason"
	// FieldRunURL holds the string denoting the run_url field in the database.
	FieldRunURL = "run_url"
	// FieldRunnerType holds the string denoting the runner_type field in the database.
	FieldRunnerType = "runner_type"
	// FieldAttestationRef holds the string denoting the attestation_ref field in the database.
	FieldAttestationRef = "attestation_ref"
	// EdgeWorkflow holds the string denoting the workflow edge name in mutations.
	EdgeWorkflow = "workflow"
	// EdgeRobotaccount holds the string denoting the robotaccount edge name in mutations.
	EdgeRobotaccount = "robotaccount"
	// EdgeContractVersion holds the string denoting the contract_version edge name in mutations.
	EdgeContractVersion = "contract_version"
	// Table holds the table name of the workflowrun in the database.
	Table = "workflow_runs"
	// WorkflowTable is the table that holds the workflow relation/edge.
	WorkflowTable = "workflow_runs"
	// WorkflowInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowInverseTable = "workflows"
	// WorkflowColumn is the table column denoting the workflow relation/edge.
	WorkflowColumn = "workflow_workflowruns"
	// RobotaccountTable is the table that holds the robotaccount relation/edge.
	RobotaccountTable = "workflow_runs"
	// RobotaccountInverseTable is the table name for the RobotAccount entity.
	// It exists in this package in order to avoid circular dependency with the "robotaccount" package.
	RobotaccountInverseTable = "robot_accounts"
	// RobotaccountColumn is the table column denoting the robotaccount relation/edge.
	RobotaccountColumn = "robot_account_workflowruns"
	// ContractVersionTable is the table that holds the contract_version relation/edge.
	ContractVersionTable = "workflow_runs"
	// ContractVersionInverseTable is the table name for the WorkflowContractVersion entity.
	// It exists in this package in order to avoid circular dependency with the "workflowcontractversion" package.
	ContractVersionInverseTable = "workflow_contract_versions"
	// ContractVersionColumn is the table column denoting the contract_version relation/edge.
	ContractVersionColumn = "workflow_run_contract_version"
)

// Columns holds all SQL columns for workflowrun fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldFinishedAt,
	FieldState,
	FieldReason,
	FieldRunURL,
	FieldRunnerType,
	FieldAttestationRef,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "workflow_runs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"robot_account_workflowruns",
	"workflow_workflowruns",
	"workflow_run_contract_version",
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

const DefaultState biz.WorkflowRunStatus = "initialized"

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s biz.WorkflowRunStatus) error {
	switch s {
	case "initialized", "success", "error", "expired", "canceled":
		return nil
	default:
		return fmt.Errorf("workflowrun: invalid enum value for state field: %q", s)
	}
}
