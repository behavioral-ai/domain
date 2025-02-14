package collective

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

const (
	ResourceUri = "urn:collective"

	AgentNID = "agent" // Restricted NID/Domain

	ThingNSS  = "thing"  // urn:{NID}:thing:{module-package}:{type}
	AspectNSS = "aspect" // urn:{NID}:aspect:{path}

)

// type Urn string
// type Uri string
type HttpExchange func(r *http.Request) (*http.Response, error)

var (
	do = func(r *http.Request) (*http.Response, error) {
		return nil, errors.New("error: Collective HttpExchange function has not be initialized")
	}
)

// Initialize - collective initialize, hosts are service hosts for cloud collective
// TODO: configure content resolver
func Initialize(ex HttpExchange, handler messaging.OpsAgent, hosts []string) error {
	if ex == nil || handler == nil {
		return errors.New("error: bad request, exchange or handler is nil")
	}
	initialize(ex, handler, httpResolver, hosts)
	return nil
}

func initialize(ex HttpExchange, handler messaging.OpsAgent, r contentResolver, hosts []string) {
	do = ex
	newContentAgent(handler, r)
}

// Applications can create as many domains/NISD as needed
// "agent" is the reserved domain for the agent collective supporting agent development

// IAppend - append
type IAppend struct {
	Thing    func(name, cn string) error
	Relation func(thing1, thing2 string) error
}

// Append -
var Append = func() *IAppend {
	return &IAppend{
		Thing: func(name, cn string) error {
			ok := thingAppend(name, cn)
			if !ok {
				return errors.New("error: bad request")
			}
			return nil
		},
		Relation: func(thing1, thing2 string) error {
			ok := relationAppend(thing1, thing2)
			if !ok {
				return errors.New("error: bad request")
			}
			return nil
		},
	}
}()

// IResolver - resolution
type IResolver struct {
	Get        func(name string, version int) ([]byte, error)
	GetRelated func(name string, version int) ([]byte, error)
	Append     func(name string, body []byte, version int) error
}

// Resolver -
var Resolver = func() *IResolver {
	return &IResolver{
		Get: func(name string, version int) ([]byte, error) {
			return contentAgent.get(name, version)
		},
		GetRelated: func(name string, version int) ([]byte, error) {
			rel, status := relationCache.get(name)
			if status != nil {
				return nil, status
			}
			return contentAgent.get(rel.Thing2, version)
		},
		Append: func(name string, body []byte, version int) error {
			return storeAppend(name, body, version)
		},
	}
}()

// Get - generic typed get
func Get[T any](name string, version int) (T, error) {
	var t T
	body, status := Resolver.Get(name, version)
	if status != nil {
		return t, status
	}
	err := json.Unmarshal(body, &t)
	if err != nil {
		return t, errors.New(fmt.Sprintf("error: JsonEncode %v", err))
	}
	return t, nil
}

// GetRelated - generic typed get
func GetRelated[T any](name string, version int) (T, error) {
	var t T

	rel, status := relationCache.get(name)
	if status != nil {
		return t, status
	}
	return Get[T](rel.Thing2, version)
}
