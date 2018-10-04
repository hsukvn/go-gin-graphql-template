package heartbeat

import (
	"github.com/graphql-go/graphql"
)

func HeartbeatQueryFieldGet() *graphql.Field {
	return &graphql.Field{
		Type:    heartbeatType,
		Resolve: heartbeatResolveGet,
	}
}
