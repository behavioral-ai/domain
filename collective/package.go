package collective

import (
	"context"
	"net/url"
)

type Urn string
type Tags string
type Resource string
type Timestamp string

// Append
//

// IAppend - append to the collective
// TODO: add tags for context and resemblance query support
//
//	add descriptions for context
type IAppend struct {
	Thing    func(ctx context.Context, name Urn, content []byte, tags Tags) error
	Aspect   func(ctx context.Context, name Urn, tags Tags) error
	Frame    func(ctx context.Context, name Urn, aspects []Urn, tags Tags) error
	Relation func(ctx context.Context, aspect, thing Urn, tags Tags) error
}

var Append = func() *IAppend {
	return &IAppend{
		Thing: func(ctx context.Context, name Urn, content []byte, tags Tags) error {
			return nil
		},
		Aspect: func(ctx context.Context, name Urn, tags Tags) error {
			return nil
		},
		Frame: func(ctx context.Context, name Urn, aspects []Urn, tags Tags) error {
			return nil
		},
		Relation: func(ctx context.Context, aspect, thing Urn, tags Tags) error {
			return nil
		},
	}
}()

// Retrieval

// ThingRequest -
type ThingRequest struct {
	Name    Urn
	Version int
}

// ThingResponse -
type ThingResponse struct {
	Name    Urn
	Content []byte
	Status  error
}

type QueryRequest struct {
	NSID string
	NSS  string
}

// IRetrieval -
// TODO : add a way to query/search list of aspects or frames
type IRetrieval struct {
	Things func(ctx context.Context, urns []ThingRequest) ([]ThingResponse, error)
	Frame  func(ctx context.Context, name Urn, version int) ([]Urn, error)
	Query  func(ctx context.Context, req QueryRequest) ([]Urn, error)
}

var Retrieve = func() *IRetrieval {
	return &IRetrieval{
		Things: func(ctx context.Context, req []ThingRequest) ([]ThingResponse, error) {
			return nil, nil
		},
		Frame: func(ctx context.Context, name Urn, version int) ([]Urn, error) {
			return nil, nil
		},
		Query: func(ctx context.Context, req QueryRequest) ([]Urn, error) {
			return nil, nil
		},
	}
}()

// Relations
//

const (
	RelationNone        = "none"
	RelationDirect      = "direct"
	RelationResemblance = "resemblance"

	VersionName = "ver"
	CompareName = "cmp"
	TopName     = "top"

	ExistValue    = "exist"
	NotExistValue = "!exist"
	LikeValue     = "like"
	NotLikeValue  = "!like"
)

type ThingConstraints interface {
	Urn | Resource | Tags
}

type RelateFilter struct {
	Urns      []Urn
	Resources []Resource
	Tags      Tags
	From      Timestamp
	To        Timestamp
}

func Relate[T ThingConstraints](ctx context.Context, thing1, thing2 T, filter1, filter2 RelateFilter, values url.Values) (string, error) {
	return RelationNone, nil
}
