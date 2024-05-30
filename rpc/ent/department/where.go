// Code generated by ent, DO NOT EDIT.

package department

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldStatus, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v uint32) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldSort, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldName, v))
}

// Ancestors applies equality check predicate on the "ancestors" field. It's identical to AncestorsEQ.
func Ancestors(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldAncestors, v))
}

// Leader applies equality check predicate on the "leader" field. It's identical to LeaderEQ.
func Leader(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldLeader, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldPhone, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldEmail, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldRemark, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v uint64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldParentID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldStatus))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v uint32) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v uint32) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...uint32) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...uint32) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v uint32) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v uint32) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v uint32) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v uint32) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldSort, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldName, v))
}

// AncestorsEQ applies the EQ predicate on the "ancestors" field.
func AncestorsEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldAncestors, v))
}

// AncestorsNEQ applies the NEQ predicate on the "ancestors" field.
func AncestorsNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldAncestors, v))
}

// AncestorsIn applies the In predicate on the "ancestors" field.
func AncestorsIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldAncestors, vs...))
}

// AncestorsNotIn applies the NotIn predicate on the "ancestors" field.
func AncestorsNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldAncestors, vs...))
}

// AncestorsGT applies the GT predicate on the "ancestors" field.
func AncestorsGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldAncestors, v))
}

// AncestorsGTE applies the GTE predicate on the "ancestors" field.
func AncestorsGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldAncestors, v))
}

// AncestorsLT applies the LT predicate on the "ancestors" field.
func AncestorsLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldAncestors, v))
}

// AncestorsLTE applies the LTE predicate on the "ancestors" field.
func AncestorsLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldAncestors, v))
}

// AncestorsContains applies the Contains predicate on the "ancestors" field.
func AncestorsContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldAncestors, v))
}

// AncestorsHasPrefix applies the HasPrefix predicate on the "ancestors" field.
func AncestorsHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldAncestors, v))
}

// AncestorsHasSuffix applies the HasSuffix predicate on the "ancestors" field.
func AncestorsHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldAncestors, v))
}

// AncestorsIsNil applies the IsNil predicate on the "ancestors" field.
func AncestorsIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldAncestors))
}

// AncestorsNotNil applies the NotNil predicate on the "ancestors" field.
func AncestorsNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldAncestors))
}

// AncestorsEqualFold applies the EqualFold predicate on the "ancestors" field.
func AncestorsEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldAncestors, v))
}

// AncestorsContainsFold applies the ContainsFold predicate on the "ancestors" field.
func AncestorsContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldAncestors, v))
}

// LeaderEQ applies the EQ predicate on the "leader" field.
func LeaderEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldLeader, v))
}

// LeaderNEQ applies the NEQ predicate on the "leader" field.
func LeaderNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldLeader, v))
}

// LeaderIn applies the In predicate on the "leader" field.
func LeaderIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldLeader, vs...))
}

// LeaderNotIn applies the NotIn predicate on the "leader" field.
func LeaderNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldLeader, vs...))
}

// LeaderGT applies the GT predicate on the "leader" field.
func LeaderGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldLeader, v))
}

// LeaderGTE applies the GTE predicate on the "leader" field.
func LeaderGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldLeader, v))
}

// LeaderLT applies the LT predicate on the "leader" field.
func LeaderLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldLeader, v))
}

// LeaderLTE applies the LTE predicate on the "leader" field.
func LeaderLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldLeader, v))
}

// LeaderContains applies the Contains predicate on the "leader" field.
func LeaderContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldLeader, v))
}

// LeaderHasPrefix applies the HasPrefix predicate on the "leader" field.
func LeaderHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldLeader, v))
}

// LeaderHasSuffix applies the HasSuffix predicate on the "leader" field.
func LeaderHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldLeader, v))
}

// LeaderIsNil applies the IsNil predicate on the "leader" field.
func LeaderIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldLeader))
}

// LeaderNotNil applies the NotNil predicate on the "leader" field.
func LeaderNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldLeader))
}

// LeaderEqualFold applies the EqualFold predicate on the "leader" field.
func LeaderEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldLeader, v))
}

// LeaderContainsFold applies the ContainsFold predicate on the "leader" field.
func LeaderContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldLeader, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldPhone, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldEmail, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldRemark, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v uint64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v uint64) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...uint64) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...uint64) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldParentID))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Department) predicate.Department {
	return predicate.Department(sql.NotPredicates(p))
}
