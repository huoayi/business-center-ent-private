// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNickName holds the string denoting the nick_name field in the database.
	FieldNickName = "nick_name"
	// FieldJpgURL holds the string denoting the jpg_url field in the database.
	FieldJpgURL = "jpg_url"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldIsFrozen holds the string denoting the is_frozen field in the database.
	FieldIsFrozen = "is_frozen"
	// FieldIsRecharge holds the string denoting the is_recharge field in the database.
	FieldIsRecharge = "is_recharge"
	// FieldUserType holds the string denoting the user_type field in the database.
	FieldUserType = "user_type"
	// FieldPopVersion holds the string denoting the pop_version field in the database.
	FieldPopVersion = "pop_version"
	// FieldAreaCode holds the string denoting the area_code field in the database.
	FieldAreaCode = "area_code"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldCloudSpace holds the string denoting the cloud_space field in the database.
	FieldCloudSpace = "cloud_space"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// EdgeLoginRecords holds the string denoting the login_records edge name in mutations.
	EdgeLoginRecords = "login_records"
	// EdgeVxSocials holds the string denoting the vx_socials edge name in mutations.
	EdgeVxSocials = "vx_socials"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeMerchants holds the string denoting the merchants edge name in mutations.
	EdgeMerchants = "merchants"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// Table holds the table name of the user in the database.
	Table = "users"
	// LoginRecordsTable is the table that holds the login_records relation/edge.
	LoginRecordsTable = "login_records"
	// LoginRecordsInverseTable is the table name for the LoginRecord entity.
	// It exists in this package in order to avoid circular dependency with the "loginrecord" package.
	LoginRecordsInverseTable = "login_records"
	// LoginRecordsColumn is the table column denoting the login_records relation/edge.
	LoginRecordsColumn = "user_id"
	// VxSocialsTable is the table that holds the vx_socials relation/edge.
	VxSocialsTable = "vx_socials"
	// VxSocialsInverseTable is the table name for the VXSocial entity.
	// It exists in this package in order to avoid circular dependency with the "vxsocial" package.
	VxSocialsInverseTable = "vx_socials"
	// VxSocialsColumn is the table column denoting the vx_socials relation/edge.
	VxSocialsColumn = "user_id"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "users"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "users"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// MerchantsTable is the table that holds the merchants relation/edge.
	MerchantsTable = "merchants"
	// MerchantsInverseTable is the table name for the Merchant entity.
	// It exists in this package in order to avoid circular dependency with the "merchant" package.
	MerchantsInverseTable = "merchants"
	// MerchantsColumn is the table column denoting the merchants relation/edge.
	MerchantsColumn = "user_id"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldNickName,
	FieldJpgURL,
	FieldPhone,
	FieldPassword,
	FieldIsFrozen,
	FieldIsRecharge,
	FieldUserType,
	FieldPopVersion,
	FieldAreaCode,
	FieldEmail,
	FieldCloudSpace,
	FieldParentID,
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
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy int64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt time.Time
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultNickName holds the default value on creation for the "nick_name" field.
	DefaultNickName string
	// DefaultJpgURL holds the default value on creation for the "jpg_url" field.
	DefaultJpgURL string
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// DefaultIsFrozen holds the default value on creation for the "is_frozen" field.
	DefaultIsFrozen bool
	// DefaultIsRecharge holds the default value on creation for the "is_recharge" field.
	DefaultIsRecharge bool
	// DefaultPopVersion holds the default value on creation for the "pop_version" field.
	DefaultPopVersion string
	// DefaultAreaCode holds the default value on creation for the "area_code" field.
	DefaultAreaCode string
	// DefaultEmail holds the default value on creation for the "email" field.
	DefaultEmail string
	// DefaultCloudSpace holds the default value on creation for the "cloud_space" field.
	DefaultCloudSpace int64
	// DefaultParentID holds the default value on creation for the "parent_id" field.
	DefaultParentID int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// UserType defines the type for the "user_type" enum field.
type UserType string

// UserTypePersonal is the default value of the UserType enum.
const DefaultUserType = UserTypePersonal

// UserType values.
const (
	UserTypePersonal UserType = "personal"
	UserTypeMerchant UserType = "merchant"
)

func (ut UserType) String() string {
	return string(ut)
}

// UserTypeValidator is a validator for the "user_type" field enum values. It is called by the builders before save.
func UserTypeValidator(ut UserType) error {
	switch ut {
	case UserTypePersonal, UserTypeMerchant:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for user_type field: %q", ut)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByNickName orders the results by the nick_name field.
func ByNickName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickName, opts...).ToFunc()
}

// ByJpgURL orders the results by the jpg_url field.
func ByJpgURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJpgURL, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByIsFrozen orders the results by the is_frozen field.
func ByIsFrozen(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFrozen, opts...).ToFunc()
}

// ByIsRecharge orders the results by the is_recharge field.
func ByIsRecharge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsRecharge, opts...).ToFunc()
}

// ByUserType orders the results by the user_type field.
func ByUserType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserType, opts...).ToFunc()
}

// ByPopVersion orders the results by the pop_version field.
func ByPopVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPopVersion, opts...).ToFunc()
}

// ByAreaCode orders the results by the area_code field.
func ByAreaCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAreaCode, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByCloudSpace orders the results by the cloud_space field.
func ByCloudSpace(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCloudSpace, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByLoginRecordsCount orders the results by login_records count.
func ByLoginRecordsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLoginRecordsStep(), opts...)
	}
}

// ByLoginRecords orders the results by login_records terms.
func ByLoginRecords(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLoginRecordsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVxSocialsCount orders the results by vx_socials count.
func ByVxSocialsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVxSocialsStep(), opts...)
	}
}

// ByVxSocials orders the results by vx_socials terms.
func ByVxSocials(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVxSocialsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMerchantsCount orders the results by merchants count.
func ByMerchantsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMerchantsStep(), opts...)
	}
}

// ByMerchants orders the results by merchants terms.
func ByMerchants(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMerchantsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrdersCount orders the results by orders count.
func ByOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrdersStep(), opts...)
	}
}

// ByOrders orders the results by orders terms.
func ByOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newLoginRecordsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LoginRecordsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LoginRecordsTable, LoginRecordsColumn),
	)
}
func newVxSocialsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VxSocialsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VxSocialsTable, VxSocialsColumn),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
func newMerchantsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MerchantsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MerchantsTable, MerchantsColumn),
	)
}
func newOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
	)
}
