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

type LogLogin struct {
	ent.Schema
}

func (LogLogin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).
			Optional().
			Comment(" User's UUID | 用户的UUID").
			Annotations(entsql.WithComments(true)),
		field.String("type").
			Comment("Login type | 登录类型").
			Annotations(entsql.WithComments(true)),
		field.String("auth_id").
			Comment("Login Auth ID|登录认证ID").
			Annotations(entsql.WithComments(true)),
		field.String("ip").
			Comment("IP address | IP地址").
			Annotations(entsql.WithComments(true)),
		field.String("location").
			Optional().
			Comment("Login location | 登录地点").
			Annotations(entsql.WithComments(true)),
		field.String("device").
			Optional().
			Comment("Login device | 登录设备").
			Annotations(entsql.WithComments(true)),
		field.String("browser").
			Optional().
			Comment("Login browser | 登录浏览器").
			Annotations(entsql.WithComments(true)),
		field.String("os").
			Optional().
			Comment("Login OS | 登录操作系统").
			Annotations(entsql.WithComments(true)),
		field.String("result").
			Optional().
			Comment("Login result | 登录结果").
			Annotations(entsql.WithComments(true)),
		field.String("message").
			Optional().
			Comment("Login message | 登录消息").
			Annotations(entsql.WithComments(true)),
		field.Time("login_at").
			Default(time.Now).
			Comment("Login time | 登录时间").
			Annotations(entsql.WithComments(true)),
	}
}

func (LogLogin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

func (LogLogin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("log_logins").
			Field("uuid").
			Unique(),
	}
}

func (LogLogin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("login_at"),
		index.Fields("uuid"),
		index.Fields("auth_id"),
		index.Fields("uuid", "login_at"),
		index.Fields("auth_id", "login_at"),
	}
}

func (LogLogin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_log_logins"},
		entsql.WithComments(true),
		schema.Comment("System Login Log Table|系统登录日志表"),
	}
}
