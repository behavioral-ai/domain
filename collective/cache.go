package collective

import (
	"github.com/behavioral-ai/core/aspect"
	"sync"
)

type content struct {
	m *sync.Map
}

var (
	cache = newCache()
)

func newCache() *content {
	cache = new(content)
	cache.m = new(sync.Map)
	return cache
}

func (c *content) get(name Urn, version int) ([]byte, *aspect.Status) {
	return nil, nil
}

func (c *content) getRelated(thing1, thing2 Urn, version int) ([]byte, *aspect.Status) {
	return nil, nil
}
