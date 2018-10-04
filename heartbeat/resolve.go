package heartbeat

import (
	"github.com/graphql-go/graphql"
)

func heartbeatResolveGet(p graphql.ResolveParams) (interface{}, error) {
	h := heartbeat{
		Status: "alive",
	}
	return h, nil
}
