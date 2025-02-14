package collective

import (
	"sync"
)

type relationT struct {
	m       *sync.Map
	resolve relationResolver
}

var (
	relationCache = newRelationCache(relationGet)
)

func newRelationCache(r relationResolver) *relationT {
	t := new(relationT)
	t.m = new(sync.Map)
	t.resolve = r
	return t
}

func (r *relationT) get(name string) (relation, error) {
	value, ok := r.m.Load(name)
	if !ok {
		rel, status := r.resolve(name)
		if status != nil {
			return relation{}, status
		}
		r.m.Store(name, rel)
		return rel, nil
	}
	if value1, ok1 := value.(relation); ok1 {
		return value1, nil
	}
	return relation{}, nil
}
