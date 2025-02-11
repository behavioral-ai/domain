package collective

import "sync"

type Timestamp string // Comparison of timestamps must support a temporal ordering

// Thing - something named
type thing struct {
	Name    Urn       `json:"name"` // Uuid
	Cn      string    `json:"cn"`
	Created Timestamp `json:"created"`
}

var (
	tm     sync.Mutex
	things []thing
)

func thingAppend(name, author Urn, cn string, ref Uri) bool {
	if thingExists(name) {
		return false
	}
	tm.Lock()
	defer tm.Unlock()
	things = append(things, thing{Name: name, Cn: cn, Created: "2024-02-11"})
	if ref != "" {
		return resolutionAppend(name, author, ref)
	}
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
