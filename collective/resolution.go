package collective

import "sync"

type resolution struct {
	Id thing `json:"id"` // Urn of related thing
	//Thing     Urn   `json:"thing"`
	Reference Uri `json:"reference"`
	Version   int `json:"version"` // System generated auto incrementing version
}

var (
	rsm         sync.Mutex
	version     = 1
	resolutions []resolution
)

func resolutionAppend(thing1 Urn, ref Uri) bool {
	//name := ResolutionUrn(thing1)
	if resolutionExists(thing1) {
		return false
	}
	rsm.Lock()
	defer rsm.Unlock()
	version++
	resolutions = append(resolutions, resolution{Id: thing{Name: thing1}, Reference: ref, Version: version})
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
