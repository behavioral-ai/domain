package collective

import (
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/jsonx"
)

const (
	ResourceUri = "urn:collective"

	AgentNID     = "agent" // Restricted NID/Domain
	EventNID     = "event"
	EventChanged = "event:changed"

	ThingNSS  = "thing"  // urn:{NID}:thing:{module-package}:{type}
	AspectNSS = "aspect" // urn:{NID}:aspect:{path}

)

type Urn string
type Uri string

// Urn
// Applications can create as many domains/NISD as needed
// "agent" is the reserved domain for the agent collective supporting agent development

// IAppend - append
type IAppend struct {
	Thing    func(name Urn, cn string) *aspect.Status
	Relation func(thing1, thing2 Urn) *aspect.Status
}

// Append -
var Append = func() *IAppend {
	return &IAppend{
		Thing: func(name Urn, cn string) *aspect.Status {
			ok := thingAppend(name, cn)
			if !ok {
				return aspect.StatusBadRequest()
			}
			return aspect.StatusOK()
		},
		Relation: func(thing1, thing2 Urn) *aspect.Status {
			ok := relationAppend(thing1, thing2)
			if !ok {
				return aspect.StatusBadRequest()
			}
			return aspect.StatusOK()
		},
	}
}()

// IResolver - resolution
type IResolver struct {
	Get        func(name Urn, version int) ([]byte, *aspect.Status)
	GetRelated func(name Urn, version int) ([]byte, *aspect.Status)
	Append     func(name Urn, body []byte) *aspect.Status
}

// Resolver -
var Resolver = func() *IResolver {
	return &IResolver{
		Get: func(name Urn, version int) ([]byte, *aspect.Status) {
			return nil, nil
		},
		GetRelated: func(name Urn, version int) ([]byte, *aspect.Status) {
			return nil, nil
		},
		Append: func(name Urn, body []byte) *aspect.Status {
			return nil
		},
	}
}()

func Get[T any](name Urn, version int) (T, *aspect.Status) {
	var t T
	body, status := Resolver.Get(name, version)
	if !status.OK() {
		return t, status
	}
	return jsonx.New[T](body, nil)
}

// Notify - notification function
type Notify func(thing, event Urn)

// AddNotification - create a notification
func AddNotification(thing Urn, fn Notify) *aspect.Status {
	return addNotification(thing, fn)
}

// ResolutionUrn - create resolution Urn
/*
//ResolutionNSS = "resolution" // urn:{NID}:resolution:testing-frame
func ResolutionUrn(name Urn) Urn {
	return Urn(strings.Replace(string(name), ThingNSS, ResolutionNSS, 1))

}

type Where struct {
	Partition, Version string
}

*/
