package scalar

import (
	"errors"
	"strconv"
)

type Uint64 uint64

func (Uint64) ImplementsGraphQLType(name string) bool {
	return name == "Uint64"
}

func (i *Uint64) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case int, int32, int64:
		if input.(int64) < 0 {
			val := 0
			*i = Uint64(val)
		} else {
			*i = Uint64(input.(int64))
		}
	case float32, float64:
		if input.(float64) < 0 {
			val := 0
			*i = Uint64(val)
		} else {
			*i = Uint64(input.(float64))
		}
	case uint, uint32, uint64:
		*i = Uint64(input.(uint64))
	case string:
		var value uint64
		value, err = strconv.ParseUint(input, 10, 64)
		if err == nil {
			*i = Uint64(value)
		}
	default:
		err = errors.New("wrong type")
	}
	return err
}

func (i Uint64) MarshalJSON() ([]byte, error) {
	b := []byte(strconv.FormatUint(uint64(i), 10))
	return b, nil
}
