package unit

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/hsukvn/go-gin-graphql-template/lib/util"
)

type Property map[string]string

func NewProperty(name string) (Property, error) {
	out, err := exec.Command("/bin/systemctl", "show", name).Output()
	if err != nil {
		return nil, err
	}

	kvs := strings.Split(strings.TrimSpace(string(out)), "\n")
	property, err := util.ParseKeyValuePairs(kvs)
	if err != nil {
		return nil, err
	}

	return property, nil
}

func (p *Property) GetName() (string, error) {
	if ns, ok := (*p)["Names"]; ok {
		names := strings.Fields(ns)

		if len(names) < 1 {
			return "", fmt.Errorf("property Names is empty")
		}

		return names[0], nil
	}

	return "", fmt.Errorf("property Names does not exist")
}
