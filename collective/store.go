package collective

import (
	"errors"
	"fmt"
	"sync"
)

//Urn -> "resiliency:thing/operative/agent/lookup/gradient"
//"resiliency:thing/operative/agent/lookup/gradient"
//"resiliency:thing/operative/agent/lookup/saturation"

// TODO: need to determine a partitioning scheme, append only
type entry struct {
	Name      string `json:"name"`    // Uuid
	Version   int    `json:"version"` // Semantic versioning MAJOR component only
	Content   []byte `json:"content"`
	CreatedTS string `json:"created-ts"`
}

var (
	sm    sync.Mutex
	store []entry
)

func storeAppend(name string, body []byte, version int) error {
	if storeExists(name, version) {
		return errors.New("error: bad request")
	}
	sm.Lock()
	defer sm.Unlock()
	store = append(store, entry{Name: name, Content: body, Version: version, CreatedTS: "2024-02-12"})
	return nil
}

func storeExists(name string, version int) bool {
	sm.Lock()
	defer sm.Unlock()
	for _, item := range store {
		if item.Name == name && item.Version == version {
			return true
		}
	}
	return false
}

func storeGet(name string, version int) ([]byte, error) {
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

//type Timestamp string // Comparison of timestamps must support a temporal ordering

type thing struct {
	Name      string `json:"name"` // Uuid
	Cn        string `json:"cn"`
	CreatedTS string `json:"created-ts"`
}

var (
	tm     sync.Mutex
	things []thing
)

func thingAppend(name string, cn string) bool {
	if thingExists(name) {
		return false
	}
	tm.Lock()
	defer tm.Unlock()
	things = append(things, thing{Name: name, Cn: cn, CreatedTS: "2024-02-11"})
	return true
}

func thingExists(name string) bool {
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
	Created string `json:"created-ts"`
	Thing1  string `json:"thing1"`
	Thing2  string `json:"thing2"`
}

var (
	rm        sync.Mutex
	relations []relation
)

func relationAppend(thing1, thing2 string) bool {
	if relationExists(thing1, thing2) {
		return false
	}
	rm.Lock()
	defer rm.Unlock()
	relations = append(relations, relation{Created: "2024-02-02", Thing1: thing1, Thing2: thing2})
	return true
}

func relationExists(thing1, thing2 string) bool {
	rm.Lock()
	defer rm.Unlock()
	for _, item := range relations {
		if item.Thing1 == thing1 && item.Thing2 == thing2 {
			return true
		}
	}
	return false
}

func relationGet(name string) (relation, error) {
	rm.Lock()
	defer rm.Unlock()
	for _, item := range relations {
		if item.Thing1 == name {
			return item, nil
		}
	}
	return relation{}, errors.New(fmt.Sprintf("error: name %v not found", name))
}
