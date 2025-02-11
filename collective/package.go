package collective

import (
	"context"
	"github.com/behavioral-ai/core/aspect"
	"io"
	"net/http"
	"net/url"
)

const (
	Name            = "urn:collective"
	AnonymousAuthor = "urn:author:anonymous"
)

type Urn string
type Uri string

// IAppend - append
type IAppend struct {
	Thing    func(ctx context.Context, name, author Urn, cn string, ref Uri) *aspect.Status
	Relation func(ctx context.Context, thing1, thing2, author Urn) *aspect.Status
}

var Append = func() *IAppend {
	return &IAppend{
		Thing: func(ctx context.Context, name, author Urn, cn string, ref Uri) *aspect.Status {
			return aspect.StatusOK()
		},
		Relation: func(ctx context.Context, thing1, thing2, author Urn) *aspect.Status {
			return aspect.StatusOK()
		},
	}
}()

// IResolver - resolution
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

func ResolutionNSS(name Urn) Urn {
	return name
}
