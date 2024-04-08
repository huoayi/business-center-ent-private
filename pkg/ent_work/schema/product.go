package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("product_name").Default("未知产品").StructTag(`json:"product_name"`).Comment("产品名称"),
		field.String("jpg_url").Default("").StructTag(`json:"jpg_url"`).Comment("产品照片"),
		field.String("comment").Default("").StructTag(`json:"comment"`).Comment("产品介绍"),
		field.Int64("price").Default(0).StructTag(`json:"price"`).Comment("单价"),
		field.String("unit").Default("").Comment("单价使用单位"),
		field.Int64("business_id").StructTag(`json:"merchant_id"`).Default(0).Comment("外键商户用户 id"),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("merchant", Merchant.Type).Ref("products").Field("business_id").Unique().Required(),
	}
}

// Mixin of User
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
