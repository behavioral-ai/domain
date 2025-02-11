package collective

import "sync"

type resolution struct {
	Id        thing `json:"id"` // Urn created automatically by replacing thing NSS with "resolution"
	Thing     Urn   `json:"thing"`
	Reference Uri   `json:"reference"`
	Version   int   `json:"version"` // System generated auto incrementing version
}

var (
	rsm         sync.Mutex
	version     = 1
	resolutions = []resolution{
		{Id: thing{Name: "urn:agent:resolution/test-thing", Cn: "cn"}, Thing: "urn:agent:thing/test-thing", Reference: "test-ref", Version: 1},
	}
)

func resolutionAppend(thing1, author Urn, ref Uri) bool {
	name := ResolutionUrn(thing1)
	if resolutionExists(name) {
		return false
	}
	rsm.Lock()
	defer rsm.Unlock()
	version++
	resolutions = append(resolutions, resolution{Id: thing{Name: name}, Thing: thing1, Reference: ref, Version: version})
	return true
}

func resolutionExists(name Urn) bool {
	rsm.Lock()
	defer rsm.Unlock()
	for _, item := range resolutions {
		if item.Id.Name == name {
			return true
		}
	}
	return false
}
