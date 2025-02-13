package collective

import (
	"github.com/behavioral-ai/core/aspect"
	"sync"
)

type contentResolver func(name Urn, version int) ([]byte, *aspect.Status)

type content struct {
	body []byte
}

type contentKey struct {
	name    Urn
	version int
}

type contentT struct {
	m       *sync.Map
	resolve contentResolver
}

var (
	contentCache = newContentCache(get)
)

func newContentCache(r contentResolver) *contentT {
	c := new(contentT)
	c.m = new(sync.Map)
	c.resolve = r
	return c
}

func (c *contentT) get(name Urn, version int) ([]byte, *aspect.Status) {
	key := contentKey{name: name, version: version}
	value, ok := c.m.Load(key)
	if !ok {
		body, status := c.resolve(name, version)
		if !status.OK() {
			return nil, status
		}
		c.m.Store(key, content{body: body})
		return body, aspect.StatusOK()
	}
	if value1, ok1 := value.(content); ok1 {
		return value1.body, aspect.StatusOK()
	}
	return nil, aspect.StatusOK()
}
