package scalar

import (
	"errors"
	"strconv"
)

type Int64 int64

func (Int64) ImplementsGraphQLType(name string) bool {
	return name == "Int64"
}

/* 
Since `internal/common/literals.go` in `graph-gophers/graphql-go` parses inut value as int32 first,
the input value is not able to pass Int64 properly.

Workaround is to use string on the client side.
*/
func (i *Int64) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case int64:
		*i = Int64(input)
	case int32:
		*i = Int64(int64(input))
	case int:
		*i = Int64(int64(input))
	case float64:
		*i = Int64(int64(input))
	case float32:
		*i = Int64(int64(input))
	case string:
		var value int64
		value, err = strconv.ParseInt(input, 10, 64)
		*i = Int64(value)
	default:
		err = errors.New("wrong type")
	}
	return err
}

func (i Int64) MarshalJSON() ([]byte, error) {
	b := []byte(strconv.FormatInt(int64(i), 10))
	return b, nil
}
