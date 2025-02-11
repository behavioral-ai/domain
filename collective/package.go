package collective

import (
	"context"
	"github.com/behavioral-ai/core/aspect"
	"io"
	"net/http"
	"net/url"
)

const (
	Name = "urn:collective"
)

type Urn string
type Uri string

// IAppend - append to the collective
type IAppend struct {
	Thing    func(ctx context.Context, name Urn, cn string, ref Uri) *aspect.Status
	Relation func(ctx context.Context, thing1, thing2 Urn, ref Uri) *aspect.Status
}

var Append = func() *IAppend {
	return &IAppend{
		Thing: func(ctx context.Context, name Urn, cn string, ref Uri) *aspect.Status {
			return aspect.StatusOK()
		},

		Relation: func(ctx context.Context, thing1, thing2 Urn, ref Uri) *aspect.Status {
			return aspect.StatusOK()
		},
	}
}()

// IResolver - resolve
type IResolver struct {
	Get     func(ctx context.Context, name Urn, values url.Values, fragment string) (body []byte, status *aspect.Status)
	Put     func(ctx context.Context, name Urn, body []byte, values url.Values, fragment string) (status *aspect.Status)
	Relate  func(ctx context.Context, thing1, thing2 Urn, values url.Values, fragment string) (body []byte, status *aspect.Status)
	Request func(ctx context.Context, name Urn, method string, headers http.Header, body io.Reader, values url.Values, fragment string) (resp *http.Response, status *aspect.Status)
}

var Resolver = func() *IResolver {
	return &IResolver{
		Get: func(ctx context.Context, name Urn, values url.Values, fragment string) (body []byte, status *aspect.Status) {
			return nil, nil
		},
		Put: func(ctx context.Context, name Urn, body []byte, values url.Values, fragment string) (status *aspect.Status) {
			return nil
		},
		Relate: func(ctx context.Context, thing1, thing2 Urn, values url.Values, fragment string) (body []byte, status *aspect.Status) {
			return nil, nil
		},
		Request: func(ctx context.Context, name Urn, method string, headers http.Header, body io.Reader, values url.Values, fragment string) (resp *http.Response, status *aspect.Status) {
			return nil, nil
		},
	}
}()
