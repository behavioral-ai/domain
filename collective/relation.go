package collective

import (
	"errors"
	"fmt"
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

func relationGet(name Urn) (relation, error) {
	tm.Lock()
	defer tm.Unlock()
	for _, item := range relations {
		if item.Thing1 == name {
			return item, nil
		}
	}
	return relation{}, errors.New(fmt.Sprintf("error: name %v not found", name))
}

type relationResolver func(name Urn) (relation, error)

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

func (r *relationT) get(name Urn) (relation, error) {
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
