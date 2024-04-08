// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/predicate"
	"github.com/huoayi/business-center-ent-private/pkg/enum"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDeletedAt, v))
}

// ProductName applies equality check predicate on the "product_name" field. It's identical to ProductNameEQ.
func ProductName(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldProductName, v))
}

// JpgURL applies equality check predicate on the "jpg_url" field. It's identical to JpgURLEQ.
func JpgURL(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldJpgURL, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldComment, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// Unit applies equality check predicate on the "unit" field. It's identical to UnitEQ.
func Unit(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUnit, v))
}

// BusinessID applies equality check predicate on the "business_id" field. It's identical to BusinessIDEQ.
func BusinessID(v int64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldBusinessID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldDeletedAt, v))
}

// ProductNameEQ applies the EQ predicate on the "product_name" field.
func ProductNameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldProductName, v))
}

// ProductNameNEQ applies the NEQ predicate on the "product_name" field.
func ProductNameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldProductName, v))
}

// ProductNameIn applies the In predicate on the "product_name" field.
func ProductNameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldProductName, vs...))
}

// ProductNameNotIn applies the NotIn predicate on the "product_name" field.
func ProductNameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldProductName, vs...))
}

// ProductNameGT applies the GT predicate on the "product_name" field.
func ProductNameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldProductName, v))
}

// ProductNameGTE applies the GTE predicate on the "product_name" field.
func ProductNameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldProductName, v))
}

// ProductNameLT applies the LT predicate on the "product_name" field.
func ProductNameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldProductName, v))
}

// ProductNameLTE applies the LTE predicate on the "product_name" field.
func ProductNameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldProductName, v))
}

// ProductNameContains applies the Contains predicate on the "product_name" field.
func ProductNameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldProductName, v))
}

// ProductNameHasPrefix applies the HasPrefix predicate on the "product_name" field.
func ProductNameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldProductName, v))
}

// ProductNameHasSuffix applies the HasSuffix predicate on the "product_name" field.
func ProductNameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldProductName, v))
}

// ProductNameEqualFold applies the EqualFold predicate on the "product_name" field.
func ProductNameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldProductName, v))
}

// ProductNameContainsFold applies the ContainsFold predicate on the "product_name" field.
func ProductNameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldProductName, v))
}

// JpgURLEQ applies the EQ predicate on the "jpg_url" field.
func JpgURLEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldJpgURL, v))
}

// JpgURLNEQ applies the NEQ predicate on the "jpg_url" field.
func JpgURLNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldJpgURL, v))
}

// JpgURLIn applies the In predicate on the "jpg_url" field.
func JpgURLIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldJpgURL, vs...))
}

// JpgURLNotIn applies the NotIn predicate on the "jpg_url" field.
func JpgURLNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldJpgURL, vs...))
}

// JpgURLGT applies the GT predicate on the "jpg_url" field.
func JpgURLGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldJpgURL, v))
}

// JpgURLGTE applies the GTE predicate on the "jpg_url" field.
func JpgURLGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldJpgURL, v))
}

// JpgURLLT applies the LT predicate on the "jpg_url" field.
func JpgURLLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldJpgURL, v))
}

// JpgURLLTE applies the LTE predicate on the "jpg_url" field.
func JpgURLLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldJpgURL, v))
}

// JpgURLContains applies the Contains predicate on the "jpg_url" field.
func JpgURLContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldJpgURL, v))
}

// JpgURLHasPrefix applies the HasPrefix predicate on the "jpg_url" field.
func JpgURLHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldJpgURL, v))
}

// JpgURLHasSuffix applies the HasSuffix predicate on the "jpg_url" field.
func JpgURLHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldJpgURL, v))
}

// JpgURLEqualFold applies the EqualFold predicate on the "jpg_url" field.
func JpgURLEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldJpgURL, v))
}

// JpgURLContainsFold applies the ContainsFold predicate on the "jpg_url" field.
func JpgURLContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldJpgURL, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldComment, v))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldComment, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int64) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int64) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int64) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int64) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int64) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int64) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int64) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldPrice, v))
}

// UnitEQ applies the EQ predicate on the "unit" field.
func UnitEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUnit, v))
}

// UnitNEQ applies the NEQ predicate on the "unit" field.
func UnitNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUnit, v))
}

// UnitIn applies the In predicate on the "unit" field.
func UnitIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUnit, vs...))
}

// UnitNotIn applies the NotIn predicate on the "unit" field.
func UnitNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUnit, vs...))
}

// UnitGT applies the GT predicate on the "unit" field.
func UnitGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUnit, v))
}

// UnitGTE applies the GTE predicate on the "unit" field.
func UnitGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUnit, v))
}

// UnitLT applies the LT predicate on the "unit" field.
func UnitLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUnit, v))
}

// UnitLTE applies the LTE predicate on the "unit" field.
func UnitLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUnit, v))
}

// UnitContains applies the Contains predicate on the "unit" field.
func UnitContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldUnit, v))
}

// UnitHasPrefix applies the HasPrefix predicate on the "unit" field.
func UnitHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldUnit, v))
}

// UnitHasSuffix applies the HasSuffix predicate on the "unit" field.
func UnitHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldUnit, v))
}

// UnitEqualFold applies the EqualFold predicate on the "unit" field.
func UnitEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldUnit, v))
}

// UnitContainsFold applies the ContainsFold predicate on the "unit" field.
func UnitContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldUnit, v))
}

// BusinessIDEQ applies the EQ predicate on the "business_id" field.
func BusinessIDEQ(v int64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldBusinessID, v))
}

// BusinessIDNEQ applies the NEQ predicate on the "business_id" field.
func BusinessIDNEQ(v int64) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldBusinessID, v))
}

// BusinessIDIn applies the In predicate on the "business_id" field.
func BusinessIDIn(vs ...int64) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldBusinessID, vs...))
}

// BusinessIDNotIn applies the NotIn predicate on the "business_id" field.
func BusinessIDNotIn(vs ...int64) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldBusinessID, vs...))
}

// ProduceTypeEQ applies the EQ predicate on the "produce_type" field.
func ProduceTypeEQ(v enum.ProduceType) predicate.Product {
	vc := v
	return predicate.Product(sql.FieldEQ(FieldProduceType, vc))
}

// ProduceTypeNEQ applies the NEQ predicate on the "produce_type" field.
func ProduceTypeNEQ(v enum.ProduceType) predicate.Product {
	vc := v
	return predicate.Product(sql.FieldNEQ(FieldProduceType, vc))
}

// ProduceTypeIn applies the In predicate on the "produce_type" field.
func ProduceTypeIn(vs ...enum.ProduceType) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(sql.FieldIn(FieldProduceType, v...))
}

// ProduceTypeNotIn applies the NotIn predicate on the "produce_type" field.
func ProduceTypeNotIn(vs ...enum.ProduceType) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(sql.FieldNotIn(FieldProduceType, v...))
}

// HasMerchant applies the HasEdge predicate on the "merchant" edge.
func HasMerchant() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MerchantTable, MerchantColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMerchantWith applies the HasEdge predicate on the "merchant" edge with a given conditions (other predicates).
func HasMerchantWith(preds ...predicate.Merchant) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newMerchantStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(sql.NotPredicates(p))
}
