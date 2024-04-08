// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LoginRecordsColumns holds the columns for the "login_records" table.
	LoginRecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建者 ID", Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新者 ID", Default: 0},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "ua", Type: field.TypeString, Comment: "用户登录时的 user-agent", Default: ""},
		{Name: "ip", Type: field.TypeString, Comment: "用户登录时的 ip 地址", Default: ""},
		{Name: "way", Type: field.TypeString, Comment: "用户登录时的登录方式，比如手机号验证码等", Default: ""},
		{Name: "user_id", Type: field.TypeInt64, Comment: "用户 id，外键关联用户", Default: 0},
	}
	// LoginRecordsTable holds the schema information for the "login_records" table.
	LoginRecordsTable = &schema.Table{
		Name:       "login_records",
		Comment:    "登录记录，记录用户登录的一些信息",
		Columns:    LoginRecordsColumns,
		PrimaryKey: []*schema.Column{LoginRecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "login_records_users_login_records",
				Columns:    []*schema.Column{LoginRecordsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "loginrecord_user_id",
				Unique:  false,
				Columns: []*schema.Column{LoginRecordsColumns[9]},
			},
		},
	}
	// MerchantsColumns holds the columns for the "merchants" table.
	MerchantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建者 ID", Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新者 ID", Default: 0},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "merchant_name", Type: field.TypeString, Comment: "商户姓名", Default: "默认商户"},
		{Name: "jpg_url", Type: field.TypeString, Comment: "商户头像", Default: ""},
		{Name: "comment", Type: field.TypeString, Comment: "商户介绍", Default: ""},
		{Name: "amount", Type: field.TypeInt, Comment: "上架商品总数", Default: 0},
		{Name: "provence", Type: field.TypeEnum, Comment: "省份", Enums: []string{"安徽", "北京", "重庆", "福建", "甘肃", "广东", "广西", "贵州", "海南", "黑龙江", "河北", "河南", "湖北", "湖南", "江苏", "江西", "吉林", "辽宁", "内蒙古", "宁夏", "青海", "山东", "山西", "陕西", "上海", "四川", "天津", "新疆", "云南", "浙江"}, Default: "河南"},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键用户 id", Default: 0},
	}
	// MerchantsTable holds the schema information for the "merchants" table.
	MerchantsTable = &schema.Table{
		Name:       "merchants",
		Columns:    MerchantsColumns,
		PrimaryKey: []*schema.Column{MerchantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "merchants_users_merchants",
				Columns:    []*schema.Column{MerchantsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建者 ID", Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新者 ID", Default: 0},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "product_name", Type: field.TypeString, Comment: "产品名称", Default: "未知产品"},
		{Name: "jpg_url", Type: field.TypeString, Comment: "产品照片", Default: ""},
		{Name: "comment", Type: field.TypeString, Comment: "产品介绍", Default: ""},
		{Name: "price", Type: field.TypeInt64, Comment: "单价", Default: 0},
		{Name: "unit", Type: field.TypeString, Comment: "单价使用单位", Default: ""},
		{Name: "produce_type", Type: field.TypeEnum, Comment: "产品类型", Enums: []string{"tea", "edible-fungi", "vegetable", "seedlings", "medicinal-materials", "grain-and-oil-crops", "fisheries", "animal-husbandry"}, Default: "tea"},
		{Name: "business_id", Type: field.TypeInt64, Comment: "外键商户用户 id", Default: 0},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_merchants_products",
				Columns:    []*schema.Column{ProductsColumns[12]},
				RefColumns: []*schema.Column{MerchantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建者 ID", Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新者 ID", Default: 0},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "name", Type: field.TypeString, Comment: "用户名", Default: "请设置昵称"},
		{Name: "nick_name", Type: field.TypeString, Comment: "用户昵称", Default: "请设置昵称"},
		{Name: "jpg_url", Type: field.TypeString, Comment: "头像", Default: ""},
		{Name: "phone", Type: field.TypeString, Comment: "用户的手机号", Default: ""},
		{Name: "password", Type: field.TypeString, Comment: "密码", Default: ""},
		{Name: "is_frozen", Type: field.TypeBool, Comment: "是否冻结", Default: false},
		{Name: "is_recharge", Type: field.TypeBool, Comment: "是否充值过", Default: false},
		{Name: "user_type", Type: field.TypeEnum, Comment: "用户类型", Enums: []string{"personal", "enterprise"}, Default: "personal"},
		{Name: "pop_version", Type: field.TypeString, Comment: "用户最新弹窗版本", Default: ""},
		{Name: "area_code", Type: field.TypeString, Comment: "国家区号", Default: "+86"},
		{Name: "email", Type: field.TypeString, Comment: "邮箱", Default: ""},
		{Name: "cloud_space", Type: field.TypeInt64, Comment: "云盘空间", Default: 0},
		{Name: "parent_id", Type: field.TypeInt64, Nullable: true, Comment: "邀请人用户 id", Default: 0},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Comment:    "用户表",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_children",
				Columns:    []*schema.Column{UsersColumns[18]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VxSocialsColumns holds the columns for the "vx_socials" table.
	VxSocialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建者 ID", Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新者 ID", Default: 0},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "app_id", Type: field.TypeString, Comment: "微信应用 id", Default: ""},
		{Name: "open_id", Type: field.TypeString, Comment: "微信身份源的 open_id", Default: ""},
		{Name: "union_id", Type: field.TypeString, Comment: "微信身份源的 union_id", Default: ""},
		{Name: "scope", Type: field.TypeEnum, Comment: "账户的权限级别，一般为 base", Enums: []string{"base"}, Default: "base"},
		{Name: "session_key", Type: field.TypeString, Comment: "小程序专用的会话密钥，不可以返回给前端", Default: ""},
		{Name: "access_token", Type: field.TypeString, Comment: "微信能力访问凭证", Default: ""},
		{Name: "refresh_token", Type: field.TypeString, Comment: "刷新微信凭证的刷新凭证", Default: ""},
		{Name: "user_id", Type: field.TypeInt64, Comment: "外键用户 id", Default: 0},
	}
	// VxSocialsTable holds the schema information for the "vx_socials" table.
	VxSocialsTable = &schema.Table{
		Name:       "vx_socials",
		Comment:    "微信社会源信息，记录用户在微信方的身份信息",
		Columns:    VxSocialsColumns,
		PrimaryKey: []*schema.Column{VxSocialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vx_socials_users_vx_socials",
				Columns:    []*schema.Column{VxSocialsColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "vxsocial_user_id",
				Unique:  false,
				Columns: []*schema.Column{VxSocialsColumns[13]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LoginRecordsTable,
		MerchantsTable,
		ProductsTable,
		UsersTable,
		VxSocialsTable,
	}
)

func init() {
	LoginRecordsTable.ForeignKeys[0].RefTable = UsersTable
	LoginRecordsTable.Annotation = &entsql.Annotation{}
	MerchantsTable.ForeignKeys[0].RefTable = UsersTable
	MerchantsTable.Annotation = &entsql.Annotation{}
	ProductsTable.ForeignKeys[0].RefTable = MerchantsTable
	ProductsTable.Annotation = &entsql.Annotation{}
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.Annotation = &entsql.Annotation{}
	VxSocialsTable.ForeignKeys[0].RefTable = UsersTable
	VxSocialsTable.Annotation = &entsql.Annotation{}
}
