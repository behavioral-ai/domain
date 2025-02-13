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

func storeAppend(name Urn, body []byte, version int) error {
	if storeExists(name, version) {
		return errors.New("error: bad request")
	}
	sm.Lock()
	defer sm.Unlock()
	store = append(store, entry{Name: name, Content: body, Version: version, CreatedTS: "2024-02-12"})
	return nil
}
