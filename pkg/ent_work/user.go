// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/user"
)

// 用户表
type User struct {
	config `json:"-"`
	// ID of the ent.
	// 19 位雪花 ID
	ID int64 `json:"id,string"`
	// 创建者 ID
	CreatedBy int64 `json:"created_by"`
	// 更新者 ID
	UpdatedBy int64 `json:"updated_by"`
	// 创建时刻，带时区
	CreatedAt time.Time `json:"created_at"`
	// 更新时刻，带时区
	UpdatedAt time.Time `json:"updated_at"`
	// 软删除时刻，带时区
	DeletedAt time.Time `json:"deleted_at"`
	// 用户名
	Name string `json:"name"`
	// 用户昵称
	NickName string `json:"nick_name"`
	// 头像
	JpgURL string `json:"jpg_url"`
	// 用户的手机号
	Phone string `json:"phone"`
	// 密码
	Password string `json:"-"`
	// 是否冻结
	IsFrozen bool `json:"is_frozen"`
	// 是否充值过
	IsRecharge bool `json:"is_recharge"`
	// 用户类型
	UserType user.UserType `json:"user_type"`
	// 用户最新弹窗版本
	PopVersion string `json:"pop_version"`
	// 国家区号
	AreaCode string `json:"area_code"`
	// 邮箱
	Email string `json:"email"'`
	// 云盘空间
	CloudSpace int64 `json:"cloud_space"`
	// 邀请人用户 id
	ParentID int64 `json:"parent_id,string"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// LoginRecords holds the value of the login_records edge.
	LoginRecords []*LoginRecord `json:"login_records,omitempty"`
	// VxSocials holds the value of the vx_socials edge.
	VxSocials []*VXSocial `json:"vx_socials,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *User `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*User `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// LoginRecordsOrErr returns the LoginRecords value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LoginRecordsOrErr() ([]*LoginRecord, error) {
	if e.loadedTypes[0] {
		return e.LoginRecords, nil
	}
	return nil, &NotLoadedError{edge: "login_records"}
}

// VxSocialsOrErr returns the VxSocials value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) VxSocialsOrErr() ([]*VXSocial, error) {
	if e.loadedTypes[1] {
		return e.VxSocials, nil
	}
	return nil, &NotLoadedError{edge: "vx_socials"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ParentOrErr() (*User, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ChildrenOrErr() ([]*User, error) {
	if e.loadedTypes[3] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsFrozen, user.FieldIsRecharge:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldCreatedBy, user.FieldUpdatedBy, user.FieldCloudSpace, user.FieldParentID:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldNickName, user.FieldJpgURL, user.FieldPhone, user.FieldPassword, user.FieldUserType, user.FieldPopVersion, user.FieldAreaCode, user.FieldEmail:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int64(value.Int64)
		case user.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				u.CreatedBy = value.Int64
			}
		case user.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				u.UpdatedBy = value.Int64
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = value.Time
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldNickName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nick_name", values[i])
			} else if value.Valid {
				u.NickName = value.String
			}
		case user.FieldJpgURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field jpg_url", values[i])
			} else if value.Valid {
				u.JpgURL = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldIsFrozen:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_frozen", values[i])
			} else if value.Valid {
				u.IsFrozen = value.Bool
			}
		case user.FieldIsRecharge:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_recharge", values[i])
			} else if value.Valid {
				u.IsRecharge = value.Bool
			}
		case user.FieldUserType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_type", values[i])
			} else if value.Valid {
				u.UserType = user.UserType(value.String)
			}
		case user.FieldPopVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pop_version", values[i])
			} else if value.Valid {
				u.PopVersion = value.String
			}
		case user.FieldAreaCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field area_code", values[i])
			} else if value.Valid {
				u.AreaCode = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldCloudSpace:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cloud_space", values[i])
			} else if value.Valid {
				u.CloudSpace = value.Int64
			}
		case user.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				u.ParentID = value.Int64
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryLoginRecords queries the "login_records" edge of the User entity.
func (u *User) QueryLoginRecords() *LoginRecordQuery {
	return NewUserClient(u.config).QueryLoginRecords(u)
}

// QueryVxSocials queries the "vx_socials" edge of the User entity.
func (u *User) QueryVxSocials() *VXSocialQuery {
	return NewUserClient(u.config).QueryVxSocials(u)
}

// QueryParent queries the "parent" edge of the User entity.
func (u *User) QueryParent() *UserQuery {
	return NewUserClient(u.config).QueryParent(u)
}

// QueryChildren queries the "children" edge of the User entity.
func (u *User) QueryChildren() *UserQuery {
	return NewUserClient(u.config).QueryChildren(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent_work: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(u.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("nick_name=")
	builder.WriteString(u.NickName)
	builder.WriteString(", ")
	builder.WriteString("jpg_url=")
	builder.WriteString(u.JpgURL)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("is_frozen=")
	builder.WriteString(fmt.Sprintf("%v", u.IsFrozen))
	builder.WriteString(", ")
	builder.WriteString("is_recharge=")
	builder.WriteString(fmt.Sprintf("%v", u.IsRecharge))
	builder.WriteString(", ")
	builder.WriteString("user_type=")
	builder.WriteString(fmt.Sprintf("%v", u.UserType))
	builder.WriteString(", ")
	builder.WriteString("pop_version=")
	builder.WriteString(u.PopVersion)
	builder.WriteString(", ")
	builder.WriteString("area_code=")
	builder.WriteString(u.AreaCode)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("cloud_space=")
	builder.WriteString(fmt.Sprintf("%v", u.CloudSpace))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", u.ParentID))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
