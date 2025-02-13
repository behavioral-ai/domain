package collective

import (
	"github.com/behavioral-ai/core/aspect"
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

func storeGet(name Urn, version int) ([]byte, *aspect.Status) {
	sm.Lock()
	defer sm.Unlock()
	for _, item := range store {
		if item.Name == name && item.Version == version {
			return item.Content, aspect.StatusOK()
		}
	}
	return nil, aspect.StatusOK()
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

func storeAppend(name Urn, body []byte, version int) *aspect.Status {
	if storeExists(name, version) {
		return aspect.StatusBadRequest()
	}
	sm.Lock()
	defer sm.Unlock()
	store = append(store, entry{Name: name, Content: body, Version: version, CreatedTS: "2024-02-12"})
	return aspect.StatusOK()
}
