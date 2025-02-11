package collective

import (
	"context"
	"github.com/behavioral-ai/core/aspect"
	"io"
	"net/http"
	"net/url"
)

const (
	Name          = "urn:collective"
	AnonymousName = "urn:author:anonymous"

	AgentNID     = "agent" // Restricted NID/Domain
	EventNID     = "event"
	EventChanged = "urn:event:changed"

	ThingNSS    = "thing"    // urn:{NID}:thing:{module-package}:{type}
	AuthorNSS   = "author"   // urn:{NID}:author:testing-aspect
	RuleNSS     = "rule"     // urn:{NID}:rule:testing-rule
	GuidanceNSS = "guidance" // urn:{NID}:guidance:testing-rule

	ResolutionNSS = "resolution" // urn:{NID}:resolution:testing-frame

)

type Urn string
type Uri string

// Urn
// Applications can create as many domains/NISD as needed
// "agent" is the reserved domain for the agent collective supporting agent development

// ResolutionUrn -
func ResolutionUrn(name Urn) Urn {
	return name
}

// IAppend - append
type IAppend struct {
	Thing      func(ctx context.Context, name, author Urn, cn string, ref Uri) *aspect.Status
	Relation   func(ctx context.Context, thing1, thing2, author Urn) *aspect.Status
	Resolution func(ctx context.Context, thing, author Urn, ref Uri) *aspect.Status
}

var Append = func() *IAppend {
	return &IAppend{
		Thing: func(ctx context.Context, name, author Urn, cn string, ref Uri) *aspect.Status {
			return aspect.StatusOK()
		},
		Relation: func(ctx context.Context, thing1, thing2, author Urn) *aspect.Status {
			return aspect.StatusOK()
		},
		Resolution: func(ctx context.Context, thing, author Urn, ref Uri) *aspect.Status {
			return aspect.StatusOK()
		},
	}
}()

// IResolver - resolution, versioning supported via values
type IResolver struct {
	Get        func(ctx context.Context, name Urn, values url.Values, fragment string) (body []byte, status *aspect.Status)
	Put        func(ctx context.Context, name Urn, body []byte, values url.Values, fragment string) (status *aspect.Status)
	GetRelated func(ctx context.Context, thing1, thing2 Urn, values url.Values, fragment string) (body []byte, status *aspect.Status)
	Request    func(ctx context.Context, name Urn, method string, headers http.Header, body io.Reader, values url.Values, fragment string) (resp *http.Response, status *aspect.Status)
}

var Resolver = func() *IResolver {
	return &IResolver{
		Get: func(ctx context.Context, name Urn, values url.Values, fragment string) (body []byte, status *aspect.Status) {
			return nil, nil
		},
		GetRelated: func(ctx context.Context, thing1, thing2 Urn, values url.Values, fragment string) (body []byte, status *aspect.Status) {
			return nil, nil
		},
		Put: func(ctx context.Context, name Urn, body []byte, values url.Values, fragment string) (status *aspect.Status) {
			return nil
		},
		Request: func(ctx context.Context, name Urn, method string, headers http.Header, body io.Reader, values url.Values, fragment string) (resp *http.Response, status *aspect.Status) {
			return nil, nil
		},
	}
}()

type Notify func(thing, event Urn)

func AddNotification(thing Urn, fn Notify) *aspect.Status {
	return addNotification(thing, fn)
}

func RemoveNotification(thing Urn) *aspect.Status {
	return removeNotification(thing)
}
