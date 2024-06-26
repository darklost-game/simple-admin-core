// Code generated by ent, DO NOT EDIT.

package logoperation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldUpdatedAt, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldUUID, v))
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldMethod, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldPath, v))
}

// Headers applies equality check predicate on the "headers" field. It's identical to HeadersEQ.
func Headers(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldHeaders, v))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldBody, v))
}

// StatusCode applies equality check predicate on the "status_code" field. It's identical to StatusCodeEQ.
func StatusCode(v int) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldStatusCode, v))
}

// ReqTime applies equality check predicate on the "req_time" field. It's identical to ReqTimeEQ.
func ReqTime(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldReqTime, v))
}

// ResTime applies equality check predicate on the "res_time" field. It's identical to ResTimeEQ.
func ResTime(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldResTime, v))
}

// CostTime applies equality check predicate on the "cost_time" field. It's identical to CostTimeEQ.
func CostTime(v uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldCostTime, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldUpdatedAt, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldUUID, vs...))
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldMethod, v))
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldMethod, v))
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldMethod, vs...))
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldMethod, vs...))
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldMethod, v))
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldMethod, v))
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldMethod, v))
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldMethod, v))
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldContains(FieldMethod, v))
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldHasPrefix(FieldMethod, v))
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldHasSuffix(FieldMethod, v))
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEqualFold(FieldMethod, v))
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldContainsFold(FieldMethod, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldContainsFold(FieldPath, v))
}

// HeadersEQ applies the EQ predicate on the "headers" field.
func HeadersEQ(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldHeaders, v))
}

// HeadersNEQ applies the NEQ predicate on the "headers" field.
func HeadersNEQ(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldHeaders, v))
}

// HeadersIn applies the In predicate on the "headers" field.
func HeadersIn(vs ...string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldHeaders, vs...))
}

// HeadersNotIn applies the NotIn predicate on the "headers" field.
func HeadersNotIn(vs ...string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldHeaders, vs...))
}

// HeadersGT applies the GT predicate on the "headers" field.
func HeadersGT(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldHeaders, v))
}

// HeadersGTE applies the GTE predicate on the "headers" field.
func HeadersGTE(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldHeaders, v))
}

// HeadersLT applies the LT predicate on the "headers" field.
func HeadersLT(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldHeaders, v))
}

// HeadersLTE applies the LTE predicate on the "headers" field.
func HeadersLTE(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldHeaders, v))
}

// HeadersContains applies the Contains predicate on the "headers" field.
func HeadersContains(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldContains(FieldHeaders, v))
}

// HeadersHasPrefix applies the HasPrefix predicate on the "headers" field.
func HeadersHasPrefix(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldHasPrefix(FieldHeaders, v))
}

// HeadersHasSuffix applies the HasSuffix predicate on the "headers" field.
func HeadersHasSuffix(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldHasSuffix(FieldHeaders, v))
}

// HeadersEqualFold applies the EqualFold predicate on the "headers" field.
func HeadersEqualFold(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEqualFold(FieldHeaders, v))
}

// HeadersContainsFold applies the ContainsFold predicate on the "headers" field.
func HeadersContainsFold(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldContainsFold(FieldHeaders, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldHasSuffix(FieldBody, v))
}

// BodyIsNil applies the IsNil predicate on the "body" field.
func BodyIsNil() predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIsNull(FieldBody))
}

// BodyNotNil applies the NotNil predicate on the "body" field.
func BodyNotNil() predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotNull(FieldBody))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldContainsFold(FieldBody, v))
}

// StatusCodeEQ applies the EQ predicate on the "status_code" field.
func StatusCodeEQ(v int) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldStatusCode, v))
}

// StatusCodeNEQ applies the NEQ predicate on the "status_code" field.
func StatusCodeNEQ(v int) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldStatusCode, v))
}

// StatusCodeIn applies the In predicate on the "status_code" field.
func StatusCodeIn(vs ...int) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldStatusCode, vs...))
}

// StatusCodeNotIn applies the NotIn predicate on the "status_code" field.
func StatusCodeNotIn(vs ...int) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldStatusCode, vs...))
}

// StatusCodeGT applies the GT predicate on the "status_code" field.
func StatusCodeGT(v int) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldStatusCode, v))
}

// StatusCodeGTE applies the GTE predicate on the "status_code" field.
func StatusCodeGTE(v int) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldStatusCode, v))
}

// StatusCodeLT applies the LT predicate on the "status_code" field.
func StatusCodeLT(v int) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldStatusCode, v))
}

// StatusCodeLTE applies the LTE predicate on the "status_code" field.
func StatusCodeLTE(v int) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldStatusCode, v))
}

// ReqTimeEQ applies the EQ predicate on the "req_time" field.
func ReqTimeEQ(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldReqTime, v))
}

// ReqTimeNEQ applies the NEQ predicate on the "req_time" field.
func ReqTimeNEQ(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldReqTime, v))
}

// ReqTimeIn applies the In predicate on the "req_time" field.
func ReqTimeIn(vs ...time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldReqTime, vs...))
}

// ReqTimeNotIn applies the NotIn predicate on the "req_time" field.
func ReqTimeNotIn(vs ...time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldReqTime, vs...))
}

// ReqTimeGT applies the GT predicate on the "req_time" field.
func ReqTimeGT(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldReqTime, v))
}

// ReqTimeGTE applies the GTE predicate on the "req_time" field.
func ReqTimeGTE(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldReqTime, v))
}

// ReqTimeLT applies the LT predicate on the "req_time" field.
func ReqTimeLT(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldReqTime, v))
}

// ReqTimeLTE applies the LTE predicate on the "req_time" field.
func ReqTimeLTE(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldReqTime, v))
}

// ResTimeEQ applies the EQ predicate on the "res_time" field.
func ResTimeEQ(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldResTime, v))
}

// ResTimeNEQ applies the NEQ predicate on the "res_time" field.
func ResTimeNEQ(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldResTime, v))
}

// ResTimeIn applies the In predicate on the "res_time" field.
func ResTimeIn(vs ...time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldResTime, vs...))
}

// ResTimeNotIn applies the NotIn predicate on the "res_time" field.
func ResTimeNotIn(vs ...time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldResTime, vs...))
}

// ResTimeGT applies the GT predicate on the "res_time" field.
func ResTimeGT(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldResTime, v))
}

// ResTimeGTE applies the GTE predicate on the "res_time" field.
func ResTimeGTE(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldResTime, v))
}

// ResTimeLT applies the LT predicate on the "res_time" field.
func ResTimeLT(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldResTime, v))
}

// ResTimeLTE applies the LTE predicate on the "res_time" field.
func ResTimeLTE(v time.Time) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldResTime, v))
}

// CostTimeEQ applies the EQ predicate on the "cost_time" field.
func CostTimeEQ(v uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldEQ(FieldCostTime, v))
}

// CostTimeNEQ applies the NEQ predicate on the "cost_time" field.
func CostTimeNEQ(v uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNEQ(FieldCostTime, v))
}

// CostTimeIn applies the In predicate on the "cost_time" field.
func CostTimeIn(vs ...uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldIn(FieldCostTime, vs...))
}

// CostTimeNotIn applies the NotIn predicate on the "cost_time" field.
func CostTimeNotIn(vs ...uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldNotIn(FieldCostTime, vs...))
}

// CostTimeGT applies the GT predicate on the "cost_time" field.
func CostTimeGT(v uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGT(FieldCostTime, v))
}

// CostTimeGTE applies the GTE predicate on the "cost_time" field.
func CostTimeGTE(v uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldGTE(FieldCostTime, v))
}

// CostTimeLT applies the LT predicate on the "cost_time" field.
func CostTimeLT(v uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLT(FieldCostTime, v))
}

// CostTimeLTE applies the LTE predicate on the "cost_time" field.
func CostTimeLTE(v uint64) predicate.LogOperation {
	return predicate.LogOperation(sql.FieldLTE(FieldCostTime, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.LogOperation {
	return predicate.LogOperation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.LogOperation {
	return predicate.LogOperation(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LogOperation) predicate.LogOperation {
	return predicate.LogOperation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LogOperation) predicate.LogOperation {
	return predicate.LogOperation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LogOperation) predicate.LogOperation {
	return predicate.LogOperation(sql.NotPredicates(p))
}
