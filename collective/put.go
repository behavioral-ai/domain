package collective

import (
	"encoding/json"
)

type text struct {
	Value string
}

func resolverPut(r *resolution, name, author string, content any, version int) error {
	var buf []byte

	switch ptr := content.(type) {
	case string:
		var err error

		v := text{ptr}
		buf, err = json.Marshal(v)
		if err != nil {
			return err
		}
	case []byte:
		buf = ptr
	default:
		var err error

		buf, err = json.Marshal(ptr)
		if err != nil {
			return err
		}
	}
	return r.agent.resolverPut(name, author, buf, version)
}
