package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/huoayi/business-center-ent-private/pkg/enum"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("count").Default(0).StructTag(`json:"count"`).Comment("数量，按照商品计量单位计算"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("总价"),
		field.String("address").Default("").StructTag(`json:"address"`).Comment("收货地址"),
		field.Int64("products_id").StructTag(`json:"merchant_id"`).Default(0).Comment("产品 id"),
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("用户 id"),
		field.Enum("status").GoType(enum.Canceled).StructTag(`json:"status"`).Default(string(enum.Canceled)).Comment("订单"),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("products", Product.Type).Ref("orders").Field("products_id").Unique().Required(),
		edge.From("user", User.Type).Ref("orders").Field("user_id").Unique().Required(),
	}
}

// Mixin of User
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
