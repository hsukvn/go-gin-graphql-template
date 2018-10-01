package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/hsukvn/go-graphql-template/types"
)

func GetHeartbeatQuery() *graphql.Field {
	return &graphql.Field{
		Type: types.HeartbeatType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			h := types.Heartbeat{
				Status: "alive",
			}
			return h, nil
		},
	}
}
