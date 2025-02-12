package collective

import "sync"

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
