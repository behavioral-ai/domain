package collective

import (
	"sync"
)

type content struct {
	body []byte
}

type contentT struct {
	m        *sync.Map
	resolver resolutionFunc
}

func newContentCache(r resolutionFunc) *contentT {
	c := new(contentT)
	c.m = new(sync.Map)
	c.resolver = r
	return c
}

func (c *contentT) load(dir string) error {

	return nil
}

func (c *contentT) resolve(name string, version int) ([]byte, error) {
	key := ResolutionKey{Name: name, Version: version}
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
