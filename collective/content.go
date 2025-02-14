package collective

import (
	"sync"
)

type content struct {
	body []byte
}

type contentKey struct {
	name    string
	version int
}

type contentT struct {
	m       *sync.Map
	resolve contentResolver
}

func newContentCache(r contentResolver) *contentT {
	c := new(contentT)
	c.m = new(sync.Map)
	c.resolve = r
	return c
}

func (c *contentT) get(name string, version int) ([]byte, error) {
	key := contentKey{name: name, version: version}
	value, ok := c.m.Load(key)
	if !ok {
		body, status := c.resolve(name, version)
		if status != nil {
			return nil, status
		}
		c.m.Store(key, content{body: body})
		return body, nil
	}
	if value1, ok1 := value.(content); ok1 {
		return value1.body, nil
	}
	return nil, nil
}
