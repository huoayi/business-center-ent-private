// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"time"

	"github.com/huoayi/business-center-ent-private/pkg/ent_work/loginrecord"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/merchant"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/order"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/product"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/schema"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/user"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/vxsocial"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	loginrecordMixin := schema.LoginRecord{}.Mixin()
	loginrecordMixinFields0 := loginrecordMixin[0].Fields()
	_ = loginrecordMixinFields0
	loginrecordFields := schema.LoginRecord{}.Fields()
	_ = loginrecordFields
	// loginrecordDescCreatedBy is the schema descriptor for created_by field.
	loginrecordDescCreatedBy := loginrecordMixinFields0[1].Descriptor()
	// loginrecord.DefaultCreatedBy holds the default value on creation for the created_by field.
	loginrecord.DefaultCreatedBy = loginrecordDescCreatedBy.Default.(int64)
	// loginrecordDescUpdatedBy is the schema descriptor for updated_by field.
	loginrecordDescUpdatedBy := loginrecordMixinFields0[2].Descriptor()
	// loginrecord.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	loginrecord.DefaultUpdatedBy = loginrecordDescUpdatedBy.Default.(int64)
	// loginrecordDescCreatedAt is the schema descriptor for created_at field.
	loginrecordDescCreatedAt := loginrecordMixinFields0[3].Descriptor()
	// loginrecord.DefaultCreatedAt holds the default value on creation for the created_at field.
	loginrecord.DefaultCreatedAt = loginrecordDescCreatedAt.Default.(func() time.Time)
	// loginrecordDescUpdatedAt is the schema descriptor for updated_at field.
	loginrecordDescUpdatedAt := loginrecordMixinFields0[4].Descriptor()
	// loginrecord.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	loginrecord.DefaultUpdatedAt = loginrecordDescUpdatedAt.Default.(func() time.Time)
	// loginrecord.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	loginrecord.UpdateDefaultUpdatedAt = loginrecordDescUpdatedAt.UpdateDefault.(func() time.Time)
	// loginrecordDescDeletedAt is the schema descriptor for deleted_at field.
	loginrecordDescDeletedAt := loginrecordMixinFields0[5].Descriptor()
	// loginrecord.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	loginrecord.DefaultDeletedAt = loginrecordDescDeletedAt.Default.(time.Time)
	// loginrecordDescUa is the schema descriptor for ua field.
	loginrecordDescUa := loginrecordFields[0].Descriptor()
	// loginrecord.DefaultUa holds the default value on creation for the ua field.
	loginrecord.DefaultUa = loginrecordDescUa.Default.(string)
	// loginrecordDescIP is the schema descriptor for ip field.
	loginrecordDescIP := loginrecordFields[1].Descriptor()
	// loginrecord.DefaultIP holds the default value on creation for the ip field.
	loginrecord.DefaultIP = loginrecordDescIP.Default.(string)
	// loginrecordDescWay is the schema descriptor for way field.
	loginrecordDescWay := loginrecordFields[2].Descriptor()
	// loginrecord.DefaultWay holds the default value on creation for the way field.
	loginrecord.DefaultWay = loginrecordDescWay.Default.(string)
	// loginrecordDescUserID is the schema descriptor for user_id field.
	loginrecordDescUserID := loginrecordFields[3].Descriptor()
	// loginrecord.DefaultUserID holds the default value on creation for the user_id field.
	loginrecord.DefaultUserID = loginrecordDescUserID.Default.(int64)
	// loginrecordDescID is the schema descriptor for id field.
	loginrecordDescID := loginrecordMixinFields0[0].Descriptor()
	// loginrecord.DefaultID holds the default value on creation for the id field.
	loginrecord.DefaultID = loginrecordDescID.Default.(func() int64)
	merchantMixin := schema.Merchant{}.Mixin()
	merchantMixinFields0 := merchantMixin[0].Fields()
	_ = merchantMixinFields0
	merchantFields := schema.Merchant{}.Fields()
	_ = merchantFields
	// merchantDescCreatedBy is the schema descriptor for created_by field.
	merchantDescCreatedBy := merchantMixinFields0[1].Descriptor()
	// merchant.DefaultCreatedBy holds the default value on creation for the created_by field.
	merchant.DefaultCreatedBy = merchantDescCreatedBy.Default.(int64)
	// merchantDescUpdatedBy is the schema descriptor for updated_by field.
	merchantDescUpdatedBy := merchantMixinFields0[2].Descriptor()
	// merchant.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	merchant.DefaultUpdatedBy = merchantDescUpdatedBy.Default.(int64)
	// merchantDescCreatedAt is the schema descriptor for created_at field.
	merchantDescCreatedAt := merchantMixinFields0[3].Descriptor()
	// merchant.DefaultCreatedAt holds the default value on creation for the created_at field.
	merchant.DefaultCreatedAt = merchantDescCreatedAt.Default.(func() time.Time)
	// merchantDescUpdatedAt is the schema descriptor for updated_at field.
	merchantDescUpdatedAt := merchantMixinFields0[4].Descriptor()
	// merchant.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	merchant.DefaultUpdatedAt = merchantDescUpdatedAt.Default.(func() time.Time)
	// merchant.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	merchant.UpdateDefaultUpdatedAt = merchantDescUpdatedAt.UpdateDefault.(func() time.Time)
	// merchantDescDeletedAt is the schema descriptor for deleted_at field.
	merchantDescDeletedAt := merchantMixinFields0[5].Descriptor()
	// merchant.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	merchant.DefaultDeletedAt = merchantDescDeletedAt.Default.(time.Time)
	// merchantDescMerchantName is the schema descriptor for merchant_name field.
	merchantDescMerchantName := merchantFields[0].Descriptor()
	// merchant.DefaultMerchantName holds the default value on creation for the merchant_name field.
	merchant.DefaultMerchantName = merchantDescMerchantName.Default.(string)
	// merchantDescJpgURL is the schema descriptor for jpg_url field.
	merchantDescJpgURL := merchantFields[1].Descriptor()
	// merchant.DefaultJpgURL holds the default value on creation for the jpg_url field.
	merchant.DefaultJpgURL = merchantDescJpgURL.Default.(string)
	// merchantDescComment is the schema descriptor for comment field.
	merchantDescComment := merchantFields[2].Descriptor()
	// merchant.DefaultComment holds the default value on creation for the comment field.
	merchant.DefaultComment = merchantDescComment.Default.(string)
	// merchantDescAmount is the schema descriptor for amount field.
	merchantDescAmount := merchantFields[3].Descriptor()
	// merchant.DefaultAmount holds the default value on creation for the amount field.
	merchant.DefaultAmount = merchantDescAmount.Default.(int)
	// merchantDescUserID is the schema descriptor for user_id field.
	merchantDescUserID := merchantFields[4].Descriptor()
	// merchant.DefaultUserID holds the default value on creation for the user_id field.
	merchant.DefaultUserID = merchantDescUserID.Default.(int64)
	// merchantDescID is the schema descriptor for id field.
	merchantDescID := merchantMixinFields0[0].Descriptor()
	// merchant.DefaultID holds the default value on creation for the id field.
	merchant.DefaultID = merchantDescID.Default.(func() int64)
	orderMixin := schema.Order{}.Mixin()
	orderMixinFields0 := orderMixin[0].Fields()
	_ = orderMixinFields0
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreatedBy is the schema descriptor for created_by field.
	orderDescCreatedBy := orderMixinFields0[1].Descriptor()
	// order.DefaultCreatedBy holds the default value on creation for the created_by field.
	order.DefaultCreatedBy = orderDescCreatedBy.Default.(int64)
	// orderDescUpdatedBy is the schema descriptor for updated_by field.
	orderDescUpdatedBy := orderMixinFields0[2].Descriptor()
	// order.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	order.DefaultUpdatedBy = orderDescUpdatedBy.Default.(int64)
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderMixinFields0[3].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() time.Time)
	// orderDescUpdatedAt is the schema descriptor for updated_at field.
	orderDescUpdatedAt := orderMixinFields0[4].Descriptor()
	// order.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	order.DefaultUpdatedAt = orderDescUpdatedAt.Default.(func() time.Time)
	// order.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	order.UpdateDefaultUpdatedAt = orderDescUpdatedAt.UpdateDefault.(func() time.Time)
	// orderDescDeletedAt is the schema descriptor for deleted_at field.
	orderDescDeletedAt := orderMixinFields0[5].Descriptor()
	// order.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	order.DefaultDeletedAt = orderDescDeletedAt.Default.(time.Time)
	// orderDescCount is the schema descriptor for count field.
	orderDescCount := orderFields[0].Descriptor()
	// order.DefaultCount holds the default value on creation for the count field.
	order.DefaultCount = orderDescCount.Default.(int64)
	// orderDescAmount is the schema descriptor for amount field.
	orderDescAmount := orderFields[1].Descriptor()
	// order.DefaultAmount holds the default value on creation for the amount field.
	order.DefaultAmount = orderDescAmount.Default.(int64)
	// orderDescAddress is the schema descriptor for address field.
	orderDescAddress := orderFields[2].Descriptor()
	// order.DefaultAddress holds the default value on creation for the address field.
	order.DefaultAddress = orderDescAddress.Default.(string)
	// orderDescProductsID is the schema descriptor for products_id field.
	orderDescProductsID := orderFields[3].Descriptor()
	// order.DefaultProductsID holds the default value on creation for the products_id field.
	order.DefaultProductsID = orderDescProductsID.Default.(int64)
	// orderDescUserID is the schema descriptor for user_id field.
	orderDescUserID := orderFields[4].Descriptor()
	// order.DefaultUserID holds the default value on creation for the user_id field.
	order.DefaultUserID = orderDescUserID.Default.(int64)
	// orderDescID is the schema descriptor for id field.
	orderDescID := orderMixinFields0[0].Descriptor()
	// order.DefaultID holds the default value on creation for the id field.
	order.DefaultID = orderDescID.Default.(func() int64)
	productMixin := schema.Product{}.Mixin()
	productMixinFields0 := productMixin[0].Fields()
	_ = productMixinFields0
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescCreatedBy is the schema descriptor for created_by field.
	productDescCreatedBy := productMixinFields0[1].Descriptor()
	// product.DefaultCreatedBy holds the default value on creation for the created_by field.
	product.DefaultCreatedBy = productDescCreatedBy.Default.(int64)
	// productDescUpdatedBy is the schema descriptor for updated_by field.
	productDescUpdatedBy := productMixinFields0[2].Descriptor()
	// product.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	product.DefaultUpdatedBy = productDescUpdatedBy.Default.(int64)
	// productDescCreatedAt is the schema descriptor for created_at field.
	productDescCreatedAt := productMixinFields0[3].Descriptor()
	// product.DefaultCreatedAt holds the default value on creation for the created_at field.
	product.DefaultCreatedAt = productDescCreatedAt.Default.(func() time.Time)
	// productDescUpdatedAt is the schema descriptor for updated_at field.
	productDescUpdatedAt := productMixinFields0[4].Descriptor()
	// product.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	product.DefaultUpdatedAt = productDescUpdatedAt.Default.(func() time.Time)
	// product.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	product.UpdateDefaultUpdatedAt = productDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productDescDeletedAt is the schema descriptor for deleted_at field.
	productDescDeletedAt := productMixinFields0[5].Descriptor()
	// product.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	product.DefaultDeletedAt = productDescDeletedAt.Default.(time.Time)
	// productDescProductName is the schema descriptor for product_name field.
	productDescProductName := productFields[0].Descriptor()
	// product.DefaultProductName holds the default value on creation for the product_name field.
	product.DefaultProductName = productDescProductName.Default.(string)
	// productDescJpgURL is the schema descriptor for jpg_url field.
	productDescJpgURL := productFields[1].Descriptor()
	// product.DefaultJpgURL holds the default value on creation for the jpg_url field.
	product.DefaultJpgURL = productDescJpgURL.Default.(string)
	// productDescComment is the schema descriptor for comment field.
	productDescComment := productFields[2].Descriptor()
	// product.DefaultComment holds the default value on creation for the comment field.
	product.DefaultComment = productDescComment.Default.(string)
	// productDescPrice is the schema descriptor for price field.
	productDescPrice := productFields[3].Descriptor()
	// product.DefaultPrice holds the default value on creation for the price field.
	product.DefaultPrice = productDescPrice.Default.(int64)
	// productDescUnit is the schema descriptor for unit field.
	productDescUnit := productFields[4].Descriptor()
	// product.DefaultUnit holds the default value on creation for the unit field.
	product.DefaultUnit = productDescUnit.Default.(string)
	// productDescBusinessID is the schema descriptor for business_id field.
	productDescBusinessID := productFields[5].Descriptor()
	// product.DefaultBusinessID holds the default value on creation for the business_id field.
	product.DefaultBusinessID = productDescBusinessID.Default.(int64)
	// productDescCount is the schema descriptor for count field.
	productDescCount := productFields[7].Descriptor()
	// product.DefaultCount holds the default value on creation for the count field.
	product.DefaultCount = productDescCount.Default.(int64)
	// productDescID is the schema descriptor for id field.
	productDescID := productMixinFields0[0].Descriptor()
	// product.DefaultID holds the default value on creation for the id field.
	product.DefaultID = productDescID.Default.(func() int64)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedBy is the schema descriptor for created_by field.
	userDescCreatedBy := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedBy holds the default value on creation for the created_by field.
	user.DefaultCreatedBy = userDescCreatedBy.Default.(int64)
	// userDescUpdatedBy is the schema descriptor for updated_by field.
	userDescUpdatedBy := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	user.DefaultUpdatedBy = userDescUpdatedBy.Default.(int64)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescDeletedAt is the schema descriptor for deleted_at field.
	userDescDeletedAt := userMixinFields0[5].Descriptor()
	// user.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	user.DefaultDeletedAt = userDescDeletedAt.Default.(time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescNickName is the schema descriptor for nick_name field.
	userDescNickName := userFields[1].Descriptor()
	// user.DefaultNickName holds the default value on creation for the nick_name field.
	user.DefaultNickName = userDescNickName.Default.(string)
	// userDescJpgURL is the schema descriptor for jpg_url field.
	userDescJpgURL := userFields[2].Descriptor()
	// user.DefaultJpgURL holds the default value on creation for the jpg_url field.
	user.DefaultJpgURL = userDescJpgURL.Default.(string)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[3].Descriptor()
	// user.DefaultPhone holds the default value on creation for the phone field.
	user.DefaultPhone = userDescPhone.Default.(string)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[4].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
	// userDescIsFrozen is the schema descriptor for is_frozen field.
	userDescIsFrozen := userFields[5].Descriptor()
	// user.DefaultIsFrozen holds the default value on creation for the is_frozen field.
	user.DefaultIsFrozen = userDescIsFrozen.Default.(bool)
	// userDescIsRecharge is the schema descriptor for is_recharge field.
	userDescIsRecharge := userFields[6].Descriptor()
	// user.DefaultIsRecharge holds the default value on creation for the is_recharge field.
	user.DefaultIsRecharge = userDescIsRecharge.Default.(bool)
	// userDescPopVersion is the schema descriptor for pop_version field.
	userDescPopVersion := userFields[8].Descriptor()
	// user.DefaultPopVersion holds the default value on creation for the pop_version field.
	user.DefaultPopVersion = userDescPopVersion.Default.(string)
	// userDescAreaCode is the schema descriptor for area_code field.
	userDescAreaCode := userFields[9].Descriptor()
	// user.DefaultAreaCode holds the default value on creation for the area_code field.
	user.DefaultAreaCode = userDescAreaCode.Default.(string)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[10].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
	// userDescCloudSpace is the schema descriptor for cloud_space field.
	userDescCloudSpace := userFields[11].Descriptor()
	// user.DefaultCloudSpace holds the default value on creation for the cloud_space field.
	user.DefaultCloudSpace = userDescCloudSpace.Default.(int64)
	// userDescParentID is the schema descriptor for parent_id field.
	userDescParentID := userFields[12].Descriptor()
	// user.DefaultParentID holds the default value on creation for the parent_id field.
	user.DefaultParentID = userDescParentID.Default.(int64)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() int64)
	vxsocialMixin := schema.VXSocial{}.Mixin()
	vxsocialMixinFields0 := vxsocialMixin[0].Fields()
	_ = vxsocialMixinFields0
	vxsocialFields := schema.VXSocial{}.Fields()
	_ = vxsocialFields
	// vxsocialDescCreatedBy is the schema descriptor for created_by field.
	vxsocialDescCreatedBy := vxsocialMixinFields0[1].Descriptor()
	// vxsocial.DefaultCreatedBy holds the default value on creation for the created_by field.
	vxsocial.DefaultCreatedBy = vxsocialDescCreatedBy.Default.(int64)
	// vxsocialDescUpdatedBy is the schema descriptor for updated_by field.
	vxsocialDescUpdatedBy := vxsocialMixinFields0[2].Descriptor()
	// vxsocial.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	vxsocial.DefaultUpdatedBy = vxsocialDescUpdatedBy.Default.(int64)
	// vxsocialDescCreatedAt is the schema descriptor for created_at field.
	vxsocialDescCreatedAt := vxsocialMixinFields0[3].Descriptor()
	// vxsocial.DefaultCreatedAt holds the default value on creation for the created_at field.
	vxsocial.DefaultCreatedAt = vxsocialDescCreatedAt.Default.(func() time.Time)
	// vxsocialDescUpdatedAt is the schema descriptor for updated_at field.
	vxsocialDescUpdatedAt := vxsocialMixinFields0[4].Descriptor()
	// vxsocial.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	vxsocial.DefaultUpdatedAt = vxsocialDescUpdatedAt.Default.(func() time.Time)
	// vxsocial.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	vxsocial.UpdateDefaultUpdatedAt = vxsocialDescUpdatedAt.UpdateDefault.(func() time.Time)
	// vxsocialDescDeletedAt is the schema descriptor for deleted_at field.
	vxsocialDescDeletedAt := vxsocialMixinFields0[5].Descriptor()
	// vxsocial.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	vxsocial.DefaultDeletedAt = vxsocialDescDeletedAt.Default.(time.Time)
	// vxsocialDescAppID is the schema descriptor for app_id field.
	vxsocialDescAppID := vxsocialFields[0].Descriptor()
	// vxsocial.DefaultAppID holds the default value on creation for the app_id field.
	vxsocial.DefaultAppID = vxsocialDescAppID.Default.(string)
	// vxsocialDescOpenID is the schema descriptor for open_id field.
	vxsocialDescOpenID := vxsocialFields[1].Descriptor()
	// vxsocial.DefaultOpenID holds the default value on creation for the open_id field.
	vxsocial.DefaultOpenID = vxsocialDescOpenID.Default.(string)
	// vxsocialDescUnionID is the schema descriptor for union_id field.
	vxsocialDescUnionID := vxsocialFields[2].Descriptor()
	// vxsocial.DefaultUnionID holds the default value on creation for the union_id field.
	vxsocial.DefaultUnionID = vxsocialDescUnionID.Default.(string)
	// vxsocialDescSessionKey is the schema descriptor for session_key field.
	vxsocialDescSessionKey := vxsocialFields[4].Descriptor()
	// vxsocial.DefaultSessionKey holds the default value on creation for the session_key field.
	vxsocial.DefaultSessionKey = vxsocialDescSessionKey.Default.(string)
	// vxsocialDescAccessToken is the schema descriptor for access_token field.
	vxsocialDescAccessToken := vxsocialFields[5].Descriptor()
	// vxsocial.DefaultAccessToken holds the default value on creation for the access_token field.
	vxsocial.DefaultAccessToken = vxsocialDescAccessToken.Default.(string)
	// vxsocialDescRefreshToken is the schema descriptor for refresh_token field.
	vxsocialDescRefreshToken := vxsocialFields[6].Descriptor()
	// vxsocial.DefaultRefreshToken holds the default value on creation for the refresh_token field.
	vxsocial.DefaultRefreshToken = vxsocialDescRefreshToken.Default.(string)
	// vxsocialDescUserID is the schema descriptor for user_id field.
	vxsocialDescUserID := vxsocialFields[7].Descriptor()
	// vxsocial.DefaultUserID holds the default value on creation for the user_id field.
	vxsocial.DefaultUserID = vxsocialDescUserID.Default.(int64)
	// vxsocialDescID is the schema descriptor for id field.
	vxsocialDescID := vxsocialMixinFields0[0].Descriptor()
	// vxsocial.DefaultID holds the default value on creation for the id field.
	vxsocial.DefaultID = vxsocialDescID.Default.(func() int64)
}
