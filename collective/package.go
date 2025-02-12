package collective

import (
	"context"
	"github.com/behavioral-ai/core/aspect"
	"io"
	"net/http"
	"net/url"
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
	Thing      func(ctx context.Context, name Urn, cn string) *aspect.Status
	Relation   func(ctx context.Context, thing1, thing2 Urn) *aspect.Status
	Resolution func(ctx context.Context, thing Urn, ref Uri) *aspect.Status
}

// Append -
var Append = func() *IAppend {
	return &IAppend{
		Thing: func(ctx context.Context, name Urn, cn string) *aspect.Status {
			ok := thingAppend(name, cn)
			if !ok {
				return aspect.StatusBadRequest()
			}
			return aspect.StatusOK()
		},
		Relation: func(ctx context.Context, thing1, thing2 Urn) *aspect.Status {
			ok := relationAppend(thing1, thing2)
			if !ok {
				return aspect.StatusBadRequest()
			}
			return aspect.StatusOK()
		},
		Resolution: func(ctx context.Context, thing Urn, ref Uri) *aspect.Status {
			ok := resolutionAppend(thing, ref)
			if !ok {
				return aspect.StatusBadRequest()
			}
			return aspect.StatusOK()
		},
	}
}()

// IResolver - resolution
type IResolver struct {
	Get        func(ctx context.Context, name Urn, version int) (body []byte, status *aspect.Status)
	GetRelated func(ctx context.Context, thing1, thing2 Urn, version int) (body []byte, status *aspect.Status)
	Put        func(ctx context.Context, name Urn, body []byte) (status *aspect.Status)
	Request    func(ctx context.Context, name Urn, method string, headers http.Header, body io.Reader, values url.Values, fragment string) (resp *http.Response, status *aspect.Status)
}

// Resolver -
var Resolver = func() *IResolver {
	return &IResolver{
		Get: func(ctx context.Context, name Urn, version int) (body []byte, status *aspect.Status) {
			return nil, nil
		},
		GetRelated: func(ctx context.Context, thing1, thing2 Urn, version int) (body []byte, status *aspect.Status) {
			return nil, nil
		},
		Put: func(ctx context.Context, name Urn, body []byte) *aspect.Status {
			return nil
		},
		Request: func(ctx context.Context, name Urn, method string, headers http.Header, body io.Reader, values url.Values, fragment string) (resp *http.Response, status *aspect.Status) {
			return nil, nil
		},
	}
}()

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
