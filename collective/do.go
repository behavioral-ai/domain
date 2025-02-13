package collective

import "github.com/behavioral-ai/core/aspect"

func httpResolver(name Urn, version int) ([]byte, *aspect.Status) {
	return nil, aspect.StatusOK()
}
