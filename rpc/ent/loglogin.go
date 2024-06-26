// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/loglogin"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

// System Login Log Table|系统登录日志表
type LogLogin struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	//  User's UUID | 用户的UUID
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Login type | 登录类型
	Type string `json:"type,omitempty"`
	// Login Auth ID|登录认证ID
	AuthID string `json:"auth_id,omitempty"`
	// IP address | IP地址
	IP string `json:"ip,omitempty"`
	// Login location | 登录地点
	Location string `json:"location,omitempty"`
	// Login device | 登录设备
	Device string `json:"device,omitempty"`
	// Login browser | 登录浏览器
	Browser string `json:"browser,omitempty"`
	// Login OS | 登录操作系统
	Os string `json:"os,omitempty"`
	// Login result | 登录结果
	Result string `json:"result,omitempty"`
	// Login message | 登录消息
	Message string `json:"message,omitempty"`
	// Login time | 登录时间
	LoginAt time.Time `json:"login_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LogLoginQuery when eager-loading is set.
	Edges        LogLoginEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LogLoginEdges holds the relations/edges for other nodes in the graph.
type LogLoginEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LogLoginEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LogLogin) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case loglogin.FieldID:
			values[i] = new(sql.NullInt64)
		case loglogin.FieldType, loglogin.FieldAuthID, loglogin.FieldIP, loglogin.FieldLocation, loglogin.FieldDevice, loglogin.FieldBrowser, loglogin.FieldOs, loglogin.FieldResult, loglogin.FieldMessage:
			values[i] = new(sql.NullString)
		case loglogin.FieldCreatedAt, loglogin.FieldUpdatedAt, loglogin.FieldLoginAt:
			values[i] = new(sql.NullTime)
		case loglogin.FieldUUID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LogLogin fields.
func (ll *LogLogin) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case loglogin.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ll.ID = uint64(value.Int64)
		case loglogin.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ll.CreatedAt = value.Time
			}
		case loglogin.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ll.UpdatedAt = value.Time
			}
		case loglogin.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				ll.UUID = *value
			}
		case loglogin.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ll.Type = value.String
			}
		case loglogin.FieldAuthID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth_id", values[i])
			} else if value.Valid {
				ll.AuthID = value.String
			}
		case loglogin.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				ll.IP = value.String
			}
		case loglogin.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				ll.Location = value.String
			}
		case loglogin.FieldDevice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device", values[i])
			} else if value.Valid {
				ll.Device = value.String
			}
		case loglogin.FieldBrowser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field browser", values[i])
			} else if value.Valid {
				ll.Browser = value.String
			}
		case loglogin.FieldOs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field os", values[i])
			} else if value.Valid {
				ll.Os = value.String
			}
		case loglogin.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				ll.Result = value.String
			}
		case loglogin.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				ll.Message = value.String
			}
		case loglogin.FieldLoginAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field login_at", values[i])
			} else if value.Valid {
				ll.LoginAt = value.Time
			}
		default:
			ll.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LogLogin.
// This includes values selected through modifiers, order, etc.
func (ll *LogLogin) Value(name string) (ent.Value, error) {
	return ll.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the LogLogin entity.
func (ll *LogLogin) QueryUser() *UserQuery {
	return NewLogLoginClient(ll.config).QueryUser(ll)
}

// Update returns a builder for updating this LogLogin.
// Note that you need to call LogLogin.Unwrap() before calling this method if this LogLogin
// was returned from a transaction, and the transaction was committed or rolled back.
func (ll *LogLogin) Update() *LogLoginUpdateOne {
	return NewLogLoginClient(ll.config).UpdateOne(ll)
}

// Unwrap unwraps the LogLogin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ll *LogLogin) Unwrap() *LogLogin {
	_tx, ok := ll.config.driver.(*txDriver)
	if !ok {
		panic("ent: LogLogin is not a transactional entity")
	}
	ll.config.driver = _tx.drv
	return ll
}

// String implements the fmt.Stringer.
func (ll *LogLogin) String() string {
	var builder strings.Builder
	builder.WriteString("LogLogin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ll.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ll.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ll.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", ll.UUID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(ll.Type)
	builder.WriteString(", ")
	builder.WriteString("auth_id=")
	builder.WriteString(ll.AuthID)
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(ll.IP)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(ll.Location)
	builder.WriteString(", ")
	builder.WriteString("device=")
	builder.WriteString(ll.Device)
	builder.WriteString(", ")
	builder.WriteString("browser=")
	builder.WriteString(ll.Browser)
	builder.WriteString(", ")
	builder.WriteString("os=")
	builder.WriteString(ll.Os)
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(ll.Result)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(ll.Message)
	builder.WriteString(", ")
	builder.WriteString("login_at=")
	builder.WriteString(ll.LoginAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// LogLogins is a parsable slice of LogLogin.
type LogLogins []*LogLogin
