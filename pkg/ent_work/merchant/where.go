// Code generated by ent, DO NOT EDIT.

package merchant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/predicate"
	"github.com/huoayi/business-center-ent-private/pkg/enum"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldDeletedAt, v))
}

// MerchantName applies equality check predicate on the "merchant_name" field. It's identical to MerchantNameEQ.
func MerchantName(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldMerchantName, v))
}

// JpgURL applies equality check predicate on the "jpg_url" field. It's identical to JpgURLEQ.
func JpgURL(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldJpgURL, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldComment, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldAmount, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldUserID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Merchant {
	return predicate.Merchant(sql.FieldLTE(FieldDeletedAt, v))
}

// MerchantNameEQ applies the EQ predicate on the "merchant_name" field.
func MerchantNameEQ(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldMerchantName, v))
}

// MerchantNameNEQ applies the NEQ predicate on the "merchant_name" field.
func MerchantNameNEQ(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldMerchantName, v))
}

// MerchantNameIn applies the In predicate on the "merchant_name" field.
func MerchantNameIn(vs ...string) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldMerchantName, vs...))
}

// MerchantNameNotIn applies the NotIn predicate on the "merchant_name" field.
func MerchantNameNotIn(vs ...string) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldMerchantName, vs...))
}

// MerchantNameGT applies the GT predicate on the "merchant_name" field.
func MerchantNameGT(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldGT(FieldMerchantName, v))
}

// MerchantNameGTE applies the GTE predicate on the "merchant_name" field.
func MerchantNameGTE(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldGTE(FieldMerchantName, v))
}

// MerchantNameLT applies the LT predicate on the "merchant_name" field.
func MerchantNameLT(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldLT(FieldMerchantName, v))
}

// MerchantNameLTE applies the LTE predicate on the "merchant_name" field.
func MerchantNameLTE(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldLTE(FieldMerchantName, v))
}

// MerchantNameContains applies the Contains predicate on the "merchant_name" field.
func MerchantNameContains(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldContains(FieldMerchantName, v))
}

// MerchantNameHasPrefix applies the HasPrefix predicate on the "merchant_name" field.
func MerchantNameHasPrefix(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldHasPrefix(FieldMerchantName, v))
}

// MerchantNameHasSuffix applies the HasSuffix predicate on the "merchant_name" field.
func MerchantNameHasSuffix(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldHasSuffix(FieldMerchantName, v))
}

// MerchantNameEqualFold applies the EqualFold predicate on the "merchant_name" field.
func MerchantNameEqualFold(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldEqualFold(FieldMerchantName, v))
}

// MerchantNameContainsFold applies the ContainsFold predicate on the "merchant_name" field.
func MerchantNameContainsFold(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldContainsFold(FieldMerchantName, v))
}

// JpgURLEQ applies the EQ predicate on the "jpg_url" field.
func JpgURLEQ(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldJpgURL, v))
}

// JpgURLNEQ applies the NEQ predicate on the "jpg_url" field.
func JpgURLNEQ(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldJpgURL, v))
}

// JpgURLIn applies the In predicate on the "jpg_url" field.
func JpgURLIn(vs ...string) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldJpgURL, vs...))
}

// JpgURLNotIn applies the NotIn predicate on the "jpg_url" field.
func JpgURLNotIn(vs ...string) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldJpgURL, vs...))
}

// JpgURLGT applies the GT predicate on the "jpg_url" field.
func JpgURLGT(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldGT(FieldJpgURL, v))
}

// JpgURLGTE applies the GTE predicate on the "jpg_url" field.
func JpgURLGTE(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldGTE(FieldJpgURL, v))
}

// JpgURLLT applies the LT predicate on the "jpg_url" field.
func JpgURLLT(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldLT(FieldJpgURL, v))
}

// JpgURLLTE applies the LTE predicate on the "jpg_url" field.
func JpgURLLTE(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldLTE(FieldJpgURL, v))
}

// JpgURLContains applies the Contains predicate on the "jpg_url" field.
func JpgURLContains(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldContains(FieldJpgURL, v))
}

// JpgURLHasPrefix applies the HasPrefix predicate on the "jpg_url" field.
func JpgURLHasPrefix(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldHasPrefix(FieldJpgURL, v))
}

// JpgURLHasSuffix applies the HasSuffix predicate on the "jpg_url" field.
func JpgURLHasSuffix(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldHasSuffix(FieldJpgURL, v))
}

// JpgURLEqualFold applies the EqualFold predicate on the "jpg_url" field.
func JpgURLEqualFold(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldEqualFold(FieldJpgURL, v))
}

// JpgURLContainsFold applies the ContainsFold predicate on the "jpg_url" field.
func JpgURLContainsFold(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldContainsFold(FieldJpgURL, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldHasSuffix(FieldComment, v))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Merchant {
	return predicate.Merchant(sql.FieldContainsFold(FieldComment, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int) predicate.Merchant {
	return predicate.Merchant(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int) predicate.Merchant {
	return predicate.Merchant(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int) predicate.Merchant {
	return predicate.Merchant(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int) predicate.Merchant {
	return predicate.Merchant(sql.FieldLTE(FieldAmount, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Merchant {
	return predicate.Merchant(sql.FieldNotIn(FieldUserID, vs...))
}

// ProvenceEQ applies the EQ predicate on the "provence" field.
func ProvenceEQ(v enum.Provence) predicate.Merchant {
	vc := v
	return predicate.Merchant(sql.FieldEQ(FieldProvence, vc))
}

// ProvenceNEQ applies the NEQ predicate on the "provence" field.
func ProvenceNEQ(v enum.Provence) predicate.Merchant {
	vc := v
	return predicate.Merchant(sql.FieldNEQ(FieldProvence, vc))
}

// ProvenceIn applies the In predicate on the "provence" field.
func ProvenceIn(vs ...enum.Provence) predicate.Merchant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Merchant(sql.FieldIn(FieldProvence, v...))
}

// ProvenceNotIn applies the NotIn predicate on the "provence" field.
func ProvenceNotIn(vs ...enum.Provence) predicate.Merchant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Merchant(sql.FieldNotIn(FieldProvence, v...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Merchant {
	return predicate.Merchant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Merchant {
	return predicate.Merchant(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProducts applies the HasEdge predicate on the "products" edge.
func HasProducts() predicate.Merchant {
	return predicate.Merchant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsWith applies the HasEdge predicate on the "products" edge with a given conditions (other predicates).
func HasProductsWith(preds ...predicate.Product) predicate.Merchant {
	return predicate.Merchant(func(s *sql.Selector) {
		step := newProductsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Merchant) predicate.Merchant {
	return predicate.Merchant(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Merchant) predicate.Merchant {
	return predicate.Merchant(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Merchant) predicate.Merchant {
	return predicate.Merchant(sql.NotPredicates(p))
}
