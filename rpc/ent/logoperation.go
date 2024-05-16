// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/logoperation"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

// System Operation Log Table| 系统操作日志表
type LogOperation struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	//  User's UUID | 用户的UUID
	UUID uuid.UUID `json:"uuid,omitempty"`
	// HTTP request method|HTTP请求方法
	Method string `json:"method,omitempty"`
	// HTTP request path|HTTP请求路径
	Path string `json:"path,omitempty"`
	// HTTP request headers|HTTP请求头部
	Headers string `json:"headers,omitempty"`
	// HTTP request body|HTTP请求体
	Body string `json:"body,omitempty"`
	// HTTP response status code|HTTP响应状态码
	StatusCode int `json:"status_code,omitempty"`
	// Time when the request was made|请求发起时间
	ReqTime time.Time `json:"req_time,omitempty"`
	// Time when the response was received|响应接收时间
	ResTime time.Time `json:"res_time,omitempty"`
	// Cost time of the request in milliseconds|请求耗时（毫秒）
	CostTime uint64 `json:"cost_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LogOperationQuery when eager-loading is set.
	Edges        LogOperationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LogOperationEdges holds the relations/edges for other nodes in the graph.
type LogOperationEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LogOperationEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LogOperation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case logoperation.FieldID, logoperation.FieldStatusCode, logoperation.FieldCostTime:
			values[i] = new(sql.NullInt64)
		case logoperation.FieldMethod, logoperation.FieldPath, logoperation.FieldHeaders, logoperation.FieldBody:
			values[i] = new(sql.NullString)
		case logoperation.FieldCreatedAt, logoperation.FieldUpdatedAt, logoperation.FieldReqTime, logoperation.FieldResTime:
			values[i] = new(sql.NullTime)
		case logoperation.FieldUUID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LogOperation fields.
func (lo *LogOperation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case logoperation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lo.ID = uint64(value.Int64)
		case logoperation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				lo.CreatedAt = value.Time
			}
		case logoperation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				lo.UpdatedAt = value.Time
			}
		case logoperation.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				lo.UUID = *value
			}
		case logoperation.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				lo.Method = value.String
			}
		case logoperation.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				lo.Path = value.String
			}
		case logoperation.FieldHeaders:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field headers", values[i])
			} else if value.Valid {
				lo.Headers = value.String
			}
		case logoperation.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				lo.Body = value.String
			}
		case logoperation.FieldStatusCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status_code", values[i])
			} else if value.Valid {
				lo.StatusCode = int(value.Int64)
			}
		case logoperation.FieldReqTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field req_time", values[i])
			} else if value.Valid {
				lo.ReqTime = value.Time
			}
		case logoperation.FieldResTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field res_time", values[i])
			} else if value.Valid {
				lo.ResTime = value.Time
			}
		case logoperation.FieldCostTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cost_time", values[i])
			} else if value.Valid {
				lo.CostTime = uint64(value.Int64)
			}
		default:
			lo.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LogOperation.
// This includes values selected through modifiers, order, etc.
func (lo *LogOperation) Value(name string) (ent.Value, error) {
	return lo.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the LogOperation entity.
func (lo *LogOperation) QueryUser() *UserQuery {
	return NewLogOperationClient(lo.config).QueryUser(lo)
}

// Update returns a builder for updating this LogOperation.
// Note that you need to call LogOperation.Unwrap() before calling this method if this LogOperation
// was returned from a transaction, and the transaction was committed or rolled back.
func (lo *LogOperation) Update() *LogOperationUpdateOne {
	return NewLogOperationClient(lo.config).UpdateOne(lo)
}

// Unwrap unwraps the LogOperation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lo *LogOperation) Unwrap() *LogOperation {
	_tx, ok := lo.config.driver.(*txDriver)
	if !ok {
		panic("ent: LogOperation is not a transactional entity")
	}
	lo.config.driver = _tx.drv
	return lo
}

// String implements the fmt.Stringer.
func (lo *LogOperation) String() string {
	var builder strings.Builder
	builder.WriteString("LogOperation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lo.ID))
	builder.WriteString("created_at=")
	builder.WriteString(lo.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(lo.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", lo.UUID))
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(lo.Method)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(lo.Path)
	builder.WriteString(", ")
	builder.WriteString("headers=")
	builder.WriteString(lo.Headers)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(lo.Body)
	builder.WriteString(", ")
	builder.WriteString("status_code=")
	builder.WriteString(fmt.Sprintf("%v", lo.StatusCode))
	builder.WriteString(", ")
	builder.WriteString("req_time=")
	builder.WriteString(lo.ReqTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("res_time=")
	builder.WriteString(lo.ResTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("cost_time=")
	builder.WriteString(fmt.Sprintf("%v", lo.CostTime))
	builder.WriteByte(')')
	return builder.String()
}

// LogOperations is a parsable slice of LogOperation.
type LogOperations []*LogOperation
