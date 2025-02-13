package collective

import (
	"github.com/behavioral-ai/core/aspect"
	"sync"
)

type relation struct {
	Created Timestamp `json:"created"`
	Thing1  Urn       `json:"thing1"`
	Thing2  Urn       `json:"thing2"`
}

var (
	rm        sync.Mutex
	relations []relation
)

func relationAppend(thing1, thing2 Urn) bool {
	if relationExists(thing1, thing2) {
		return false
	}
	rm.Lock()
	defer rm.Unlock()
	relations = append(relations, relation{Created: "2024-02-02", Thing1: thing1, Thing2: thing2})
	return true
}

func relationExists(thing1, thing2 Urn) bool {
	tm.Lock()
	defer tm.Unlock()
	for _, item := range relations {
		if item.Thing1 == thing1 && item.Thing2 == thing2 {
			return true
		}
	}
	return false
}

func relationGet(name Urn) (relation, *aspect.Status) {
	tm.Lock()
	defer tm.Unlock()
	for _, item := range relations {
		if item.Thing1 == name {
			return item, aspect.StatusOK()
		}
	}
	return relation{}, aspect.StatusNotFound()
}

type relationResolver func(name Urn) (relation, *aspect.Status)

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

func (r *relationT) get(name Urn) (relation, *aspect.Status) {
	value, ok := r.m.Load(name)
	if !ok {
		rel, status := r.resolve(name)
		if !status.OK() {
			return relation{}, status
		}
		r.m.Store(name, rel)
		return rel, aspect.StatusOK()
	}
	if value1, ok1 := value.(relation); ok1 {
		return value1, aspect.StatusOK()
	}
	return relation{}, nil
}
