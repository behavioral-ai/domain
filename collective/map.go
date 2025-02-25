package collective

import (
	"errors"
	"fmt"
	"sync"
)

type mapT struct {
	m *sync.Map
}

func newMapCache() *mapT {
	c := new(mapT)
	c.m = new(sync.Map)
	return c
}

func (c *mapT) get(name string) (map[string]string, error) {
	value, ok := c.m.Load(name)
	if !ok {
		return nil, errors.New(fmt.Sprintf("map %v not found", name))
	}
	if value1, ok1 := value.(map[string]string); ok1 {
		return value1, nil
	}
	return nil, errors.New(fmt.Sprintf("map %v type is invalid", name))
}

func (c *mapT) put(name string, m map[string]string) error {
	if m == nil {
		return errors.New(fmt.Sprintf(" map %v is nil", name))
	}
	var ok bool
	if name == "" {
		name, ok = m[ResolutionKeyName]
		if !ok {
			return errors.New(fmt.Sprintf("map key %v is not found in map", ResolutionKeyName))
		}
	}
	c.m.Store(name, m)
	return nil
}
