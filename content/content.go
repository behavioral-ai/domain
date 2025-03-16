package content

import (
	"errors"
	"fmt"
	"sync"
)

// resolutionKey -
type resolutionKey struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

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
	key := resolutionKey{Name: name, Version: version}
	value, ok := c.m.Load(key)
	if !ok {
		return nil, errors.New(fmt.Sprintf("content [%v] [%v] not found", name, version))
	}
	if value1, ok1 := value.(content); ok1 {
		return value1.body, nil
	}
	return nil, nil
}

func (c *contentT) put(name string, body []byte, version int) {
	c.m.Store(resolutionKey{Name: name, Version: version}, content{body: body})
}
