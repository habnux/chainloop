// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflow"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowcontract"
	"github.com/google/uuid"
)

// Workflow is the model entity for the Workflow schema.
type Workflow struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Project holds the value of the "project" field.
	Project string `json:"project,omitempty"`
	// Team holds the value of the "team" field.
	Team string `json:"team,omitempty"`
	// RunsCount holds the value of the "runs_count" field.
	RunsCount int `json:"runs_count,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowQuery when eager-loading is set.
	Edges             WorkflowEdges `json:"edges"`
	organization_id   *uuid.UUID
	workflow_contract *uuid.UUID
}

// WorkflowEdges holds the relations/edges for other nodes in the graph.
type WorkflowEdges struct {
	// Robotaccounts holds the value of the robotaccounts edge.
	Robotaccounts []*RobotAccount `json:"robotaccounts,omitempty"`
	// Workflowruns holds the value of the workflowruns edge.
	Workflowruns []*WorkflowRun `json:"workflowruns,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// Contract holds the value of the contract edge.
	Contract *WorkflowContract `json:"contract,omitempty"`
	// IntegrationAttachments holds the value of the integration_attachments edge.
	IntegrationAttachments []*IntegrationAttachment `json:"integration_attachments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// RobotaccountsOrErr returns the Robotaccounts value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowEdges) RobotaccountsOrErr() ([]*RobotAccount, error) {
	if e.loadedTypes[0] {
		return e.Robotaccounts, nil
	}
	return nil, &NotLoadedError{edge: "robotaccounts"}
}

// WorkflowrunsOrErr returns the Workflowruns value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowEdges) WorkflowrunsOrErr() ([]*WorkflowRun, error) {
	if e.loadedTypes[1] {
		return e.Workflowruns, nil
	}
	return nil, &NotLoadedError{edge: "workflowruns"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowEdges) OrganizationOrErr() (*Organization, error) {
	if e.loadedTypes[2] {
		if e.Organization == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// ContractOrErr returns the Contract value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowEdges) ContractOrErr() (*WorkflowContract, error) {
	if e.loadedTypes[3] {
		if e.Contract == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: workflowcontract.Label}
		}
		return e.Contract, nil
	}
	return nil, &NotLoadedError{edge: "contract"}
}

// IntegrationAttachmentsOrErr returns the IntegrationAttachments value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowEdges) IntegrationAttachmentsOrErr() ([]*IntegrationAttachment, error) {
	if e.loadedTypes[4] {
		return e.IntegrationAttachments, nil
	}
	return nil, &NotLoadedError{edge: "integration_attachments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Workflow) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workflow.FieldRunsCount:
			values[i] = new(sql.NullInt64)
		case workflow.FieldName, workflow.FieldProject, workflow.FieldTeam:
			values[i] = new(sql.NullString)
		case workflow.FieldCreatedAt, workflow.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case workflow.FieldID:
			values[i] = new(uuid.UUID)
		case workflow.ForeignKeys[0]: // organization_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case workflow.ForeignKeys[1]: // workflow_contract
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Workflow", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Workflow fields.
func (w *Workflow) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workflow.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				w.ID = *value
			}
		case workflow.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				w.Name = value.String
			}
		case workflow.FieldProject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field project", values[i])
			} else if value.Valid {
				w.Project = value.String
			}
		case workflow.FieldTeam:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team", values[i])
			} else if value.Valid {
				w.Team = value.String
			}
		case workflow.FieldRunsCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field runs_count", values[i])
			} else if value.Valid {
				w.RunsCount = int(value.Int64)
			}
		case workflow.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case workflow.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				w.DeletedAt = value.Time
			}
		case workflow.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				w.organization_id = new(uuid.UUID)
				*w.organization_id = *value.S.(*uuid.UUID)
			}
		case workflow.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_contract", values[i])
			} else if value.Valid {
				w.workflow_contract = new(uuid.UUID)
				*w.workflow_contract = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryRobotaccounts queries the "robotaccounts" edge of the Workflow entity.
func (w *Workflow) QueryRobotaccounts() *RobotAccountQuery {
	return NewWorkflowClient(w.config).QueryRobotaccounts(w)
}

// QueryWorkflowruns queries the "workflowruns" edge of the Workflow entity.
func (w *Workflow) QueryWorkflowruns() *WorkflowRunQuery {
	return NewWorkflowClient(w.config).QueryWorkflowruns(w)
}

// QueryOrganization queries the "organization" edge of the Workflow entity.
func (w *Workflow) QueryOrganization() *OrganizationQuery {
	return NewWorkflowClient(w.config).QueryOrganization(w)
}

// QueryContract queries the "contract" edge of the Workflow entity.
func (w *Workflow) QueryContract() *WorkflowContractQuery {
	return NewWorkflowClient(w.config).QueryContract(w)
}

// QueryIntegrationAttachments queries the "integration_attachments" edge of the Workflow entity.
func (w *Workflow) QueryIntegrationAttachments() *IntegrationAttachmentQuery {
	return NewWorkflowClient(w.config).QueryIntegrationAttachments(w)
}

// Update returns a builder for updating this Workflow.
// Note that you need to call Workflow.Unwrap() before calling this method if this Workflow
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Workflow) Update() *WorkflowUpdateOne {
	return NewWorkflowClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Workflow entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Workflow) Unwrap() *Workflow {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Workflow is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Workflow) String() string {
	var builder strings.Builder
	builder.WriteString("Workflow(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("name=")
	builder.WriteString(w.Name)
	builder.WriteString(", ")
	builder.WriteString("project=")
	builder.WriteString(w.Project)
	builder.WriteString(", ")
	builder.WriteString("team=")
	builder.WriteString(w.Team)
	builder.WriteString(", ")
	builder.WriteString("runs_count=")
	builder.WriteString(fmt.Sprintf("%v", w.RunsCount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(w.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Workflows is a parsable slice of Workflow.
type Workflows []*Workflow
