package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/huoayi/business-center-ent-private/pkg/enum"
)

type Merchant struct {
	ent.Schema
}

func (Merchant) Fields() []ent.Field {
	return []ent.Field{
		field.String("merchant_name").Default("默认商户").StructTag(`json:"merchant_name"`).Comment("商户姓名"),
		field.String("jpg_url").Default("").StructTag(`json:"jpg_url"`).Comment("商户头像"),
		field.String("comment").Default("").StructTag(`json:"comment"`).Comment("商户介绍"),
		field.Int("amount").Default(0).Comment("上架商品总数"),
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Default(0).Comment("外键用户 id"),
		field.Enum("provence").GoType(enum.HeNan).StructTag(`json:"provence"`).Default(string(enum.HeNan)).Comment("省份"),
		field.String("pay_url").Default("").StructTag(`json:"pay_url"`).Comment("支付路径"),
	}
}

func (Merchant) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("merchants").Field("user_id").Unique().Required(),
		edge.To("products", Product.Type),
	}
}

// Mixin of User
func (Merchant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
