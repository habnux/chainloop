// Code generated by ent, DO NOT EDIT.

package workflowcontractversion

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldLTE(FieldID, id))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v []byte) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldEQ(FieldBody, v))
}

// Revision applies equality check predicate on the "revision" field. It's identical to RevisionEQ.
func Revision(v int) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldEQ(FieldRevision, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldEQ(FieldCreatedAt, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v []byte) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v []byte) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...[]byte) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...[]byte) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v []byte) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v []byte) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v []byte) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v []byte) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldLTE(FieldBody, v))
}

// RevisionEQ applies the EQ predicate on the "revision" field.
func RevisionEQ(v int) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldEQ(FieldRevision, v))
}

// RevisionNEQ applies the NEQ predicate on the "revision" field.
func RevisionNEQ(v int) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldNEQ(FieldRevision, v))
}

// RevisionIn applies the In predicate on the "revision" field.
func RevisionIn(vs ...int) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldIn(FieldRevision, vs...))
}

// RevisionNotIn applies the NotIn predicate on the "revision" field.
func RevisionNotIn(vs ...int) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldNotIn(FieldRevision, vs...))
}

// RevisionGT applies the GT predicate on the "revision" field.
func RevisionGT(v int) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldGT(FieldRevision, v))
}

// RevisionGTE applies the GTE predicate on the "revision" field.
func RevisionGTE(v int) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldGTE(FieldRevision, v))
}

// RevisionLT applies the LT predicate on the "revision" field.
func RevisionLT(v int) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldLT(FieldRevision, v))
}

// RevisionLTE applies the LTE predicate on the "revision" field.
func RevisionLTE(v int) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldLTE(FieldRevision, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(sql.FieldLTE(FieldCreatedAt, v))
}

// HasContract applies the HasEdge predicate on the "contract" edge.
func HasContract() predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ContractTable, ContractColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContractWith applies the HasEdge predicate on the "contract" edge with a given conditions (other predicates).
func HasContractWith(preds ...predicate.WorkflowContract) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ContractInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ContractTable, ContractColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkflowContractVersion) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkflowContractVersion) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WorkflowContractVersion) predicate.WorkflowContractVersion {
	return predicate.WorkflowContractVersion(func(s *sql.Selector) {
		p(s.Not())
	})
}
