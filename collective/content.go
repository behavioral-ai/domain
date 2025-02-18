package collective

import (
	"errors"
	"fmt"
	"sync"
)

type content struct {
	body []byte
}

type contentT struct {
	m *sync.Map
}

func newContentCache() *contentT {
	c := new(contentT)
	c.m = new(sync.Map)
	return c
}

func (c *contentT) get(name string, version int) ([]byte, error) {
	if name == "" || version <= 0 {
		return nil, errors.New(fmt.Sprintf("error: BadRequest - name \"%v\" or version \"%v\" is empty", name, version))
	}
	key := ResolutionKey{Name: name, Version: version}
	value, ok := c.m.Load(key)
	if !ok {
		return nil, errors.New(fmt.Sprintf("error: NotFound - name \"%v\" version \"%v\"", name, version))
	}
	if value1, ok1 := value.(content); ok1 {
		return value1.body, nil
	}
	return nil, nil
}

func (c *contentT) put(name string, body []byte, version int) error {
	if name == "" || body == nil || version <= 0 {
		return errors.New(fmt.Sprintf("error: BadRequest - name \"%v\", body \"%v\", or version \"%v\" is empty", name, body, version))
	}
	c.m.Store(ResolutionKey{Name: name, Version: version}, content{body: body})
	return nil
}
