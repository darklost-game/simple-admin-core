package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type LogOperation struct {
	ent.Schema
}

func (LogOperation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).
			Comment(" User's UUID | 用户的UUID").
			Annotations(entsql.WithComments(true)),
		field.String("method").
			Comment("HTTP request method|HTTP请求方法").
			Annotations(entsql.WithComments(true)),
		field.String("path").
			Comment("HTTP request path|HTTP请求路径").
			Annotations(entsql.WithComments(true)),
		field.String("headers").
			MaxLen(2048).
			Comment("HTTP request headers|HTTP请求头部").
			Annotations(entsql.WithComments(true)),
		field.Text("body").Optional().
			Comment("HTTP request body|HTTP请求体").
			Annotations(entsql.WithComments(true)),
		field.Int("status_code").
			Comment("HTTP response status code|HTTP响应状态码").
			Annotations(entsql.WithComments(true)),
		field.Time("req_time").Default(time.Now).
			Comment("Time when the request was made|请求发起时间").
			Annotations(entsql.WithComments(true)),
		field.Time("res_time").Default(time.Now).
			Comment("Time when the response was received|响应接收时间").
			Annotations(entsql.WithComments(true)),
		field.Uint64("cost_time").
			Comment("Cost time of the request in milliseconds|请求耗时（毫秒）").
			Annotations(entsql.WithComments(true)),
	}
}

func (LogOperation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

func (LogOperation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("log_operations").
			Field("uuid").
			Required().
			Unique(),
	}
}

func (LogOperation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("req_time"),
		index.Fields("uuid"),
		index.Fields("uuid", "req_time"),
	}
}

func (LogOperation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_log_operations"},
		entsql.WithComments(true),
		schema.Comment("System Operation Log Table| 系统操作日志表"),
	}
}
