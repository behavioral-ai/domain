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
	c := new(content)
	c.m = new(sync.Map)
	return c
}

func (c *content) get(name Urn, version int) ([]byte, *aspect.Status) {
	return nil, nil
}
