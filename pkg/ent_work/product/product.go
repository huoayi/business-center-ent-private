// Code generated by ent, DO NOT EDIT.

package product

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/huoayi/business-center-ent-private/pkg/enum"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
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
	// FieldProductName holds the string denoting the product_name field in the database.
	FieldProductName = "product_name"
	// FieldJpgURL holds the string denoting the jpg_url field in the database.
	FieldJpgURL = "jpg_url"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldUnit holds the string denoting the unit field in the database.
	FieldUnit = "unit"
	// FieldBusinessID holds the string denoting the business_id field in the database.
	FieldBusinessID = "business_id"
	// FieldProduceType holds the string denoting the produce_type field in the database.
	FieldProduceType = "produce_type"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// EdgeMerchant holds the string denoting the merchant edge name in mutations.
	EdgeMerchant = "merchant"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// Table holds the table name of the product in the database.
	Table = "products"
	// MerchantTable is the table that holds the merchant relation/edge.
	MerchantTable = "products"
	// MerchantInverseTable is the table name for the Merchant entity.
	// It exists in this package in order to avoid circular dependency with the "merchant" package.
	MerchantInverseTable = "merchants"
	// MerchantColumn is the table column denoting the merchant relation/edge.
	MerchantColumn = "business_id"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "products_id"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldProductName,
	FieldJpgURL,
	FieldComment,
	FieldPrice,
	FieldUnit,
	FieldBusinessID,
	FieldProduceType,
	FieldCount,
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
	// DefaultProductName holds the default value on creation for the "product_name" field.
	DefaultProductName string
	// DefaultJpgURL holds the default value on creation for the "jpg_url" field.
	DefaultJpgURL string
	// DefaultComment holds the default value on creation for the "comment" field.
	DefaultComment string
	// DefaultPrice holds the default value on creation for the "price" field.
	DefaultPrice int64
	// DefaultUnit holds the default value on creation for the "unit" field.
	DefaultUnit string
	// DefaultBusinessID holds the default value on creation for the "business_id" field.
	DefaultBusinessID int64
	// DefaultCount holds the default value on creation for the "count" field.
	DefaultCount int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultProduceType enum.ProduceType = "tea"

// ProduceTypeValidator is a validator for the "produce_type" field enum values. It is called by the builders before save.
func ProduceTypeValidator(pt enum.ProduceType) error {
	switch pt {
	case "tea", "edible-fungi", "vegetable", "seedlings", "medicinal-materials", "grain-and-oil-crops", "fisheries", "animal-husbandry":
		return nil
	default:
		return fmt.Errorf("product: invalid enum value for produce_type field: %q", pt)
	}
}

// OrderOption defines the ordering options for the Product queries.
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

// ByProductName orders the results by the product_name field.
func ByProductName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductName, opts...).ToFunc()
}

// ByJpgURL orders the results by the jpg_url field.
func ByJpgURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJpgURL, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByUnit orders the results by the unit field.
func ByUnit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnit, opts...).ToFunc()
}

// ByBusinessID orders the results by the business_id field.
func ByBusinessID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessID, opts...).ToFunc()
}

// ByProduceType orders the results by the produce_type field.
func ByProduceType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProduceType, opts...).ToFunc()
}

// ByCount orders the results by the count field.
func ByCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCount, opts...).ToFunc()
}

// ByMerchantField orders the results by merchant field.
func ByMerchantField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMerchantStep(), sql.OrderByField(field, opts...))
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
func newMerchantStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MerchantInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, MerchantTable, MerchantColumn),
	)
}
func newOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
	)
}
