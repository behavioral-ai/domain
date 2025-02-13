package collective

import (
	"errors"
	"fmt"
	"sync"
)

// TODO: need to determine a partitioning scheme. Entries work like transactions with no updates or
//
//	deletions, only appending
type entry struct {
	Name      Urn    `json:"name"`    // Uuid
	Version   int    `json:"version"` // Semantic versioning MAJOR component only
	Content   []byte `json:"content"`
	CreatedTS string `json:"created-ts"`
}

var (
	sm    sync.Mutex
	store []entry
)

func storeAppend(name Urn, body []byte, version int) error {
	if storeExists(name, version) {
		return errors.New("error: bad request")
	}
	sm.Lock()
	defer sm.Unlock()
	store = append(store, entry{Name: name, Content: body, Version: version, CreatedTS: "2024-02-12"})
	return nil
}

func storeExists(name Urn, version int) bool {
	sm.Lock()
	defer sm.Unlock()
	for _, item := range store {
		if item.Name == name && item.Version == version {
			return true
		}
	}
	return false
}

func storeGet(name Urn, version int) ([]byte, error) {
	sm.Lock()
	defer sm.Unlock()
	for _, item := range store {
		if item.Name == name && item.Version == version {
			return item.Content, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("error: name %v not found", name))
}

// Thing
//

type Timestamp string // Comparison of timestamps must support a temporal ordering

type thing struct {
	Name    Urn       `json:"name"` // Uuid
	Cn      string    `json:"cn"`
	Created Timestamp `json:"created"`
}

var (
	tm     sync.Mutex
	things []thing
)

func thingAppend(name Urn, cn string) bool {
	if thingExists(name) {
		return false
	}
	tm.Lock()
	defer tm.Unlock()
	things = append(things, thing{Name: name, Cn: cn, Created: "2024-02-11"})
	return true
}

func thingExists(name Urn) bool {
	tm.Lock()
	defer tm.Unlock()
	for _, item := range things {
		if item.Name == name {
			return true
		}
	}
	return false
}

// Relation
//

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
	rm.Lock()
	defer rm.Unlock()
	for _, item := range relations {
		if item.Thing1 == thing1 && item.Thing2 == thing2 {
			return true
		}
	}
	return false
}

func relationGet(name Urn) (relation, error) {
	rm.Lock()
	defer rm.Unlock()
	for _, item := range relations {
		if item.Thing1 == name {
			return item, nil
		}
	}
	return relation{}, errors.New(fmt.Sprintf("error: name %v not found", name))
}
