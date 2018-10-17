package util

import (
	"fmt"
	"regexp"
)

type Pair struct {
	key   string
	value string
}

func ParseKeyValuePairs(inputs []string) (map[string]string, error) {
	m := make(map[string]string)

	for _, input := range inputs {
		pair, err := ParseKeyValuePair(input)

		if err != nil {
			return nil, err
		}

		m[pair.key] = pair.value
	}

	return m, nil
}

func ParseKeyValuePair(input string) (Pair, error) {
	rexp := regexp.MustCompile("([^=]*)=(.*)")
	match := rexp.FindStringSubmatch(input)

	// match[0]: whole match string
	// match[1]: first captured group
	// match[2]: second captured group
	if len(match) != 3 {
		return Pair{}, fmt.Errorf("(%v) is not a key value pair", input)
	}

	pair := Pair{
		key:   match[1],
		value: match[2],
	}

	return pair, nil
}
