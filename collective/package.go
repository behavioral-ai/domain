package collective

import (
	"encoding/json"
	"errors"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

const (
	ResourceUri = "urn:collective"

	AgentNID = "agent" // Restricted NID/Domain

	ThingNSS  = "thing"  // urn:{NID}:thing:{module-package}:{type}
	AspectNSS = "aspect" // urn:{NID}:aspect:{path}

)

type Urn string
type Uri string
type HttpExchange func(r *http.Request) (*http.Response, error)

var (
	do = func(r *http.Request) (*http.Response, error) {
		return nil, errors.New("error: Collective HttpExchange function has not be initialized")
	}
)

// Initialize - collective initialize, hosts are service hosts for cloud collective
// TODO: configure content resolver
func Initialize(ex HttpExchange, handler messaging.OpsAgent, hosts []string) *aspect.Status {
	if ex == nil || handler == nil {
		return aspect.StatusBadRequest()
	}
	initialize(ex, handler, httpResolver, hosts)
	return aspect.StatusOK()
}

func initialize(ex HttpExchange, handler messaging.OpsAgent, r contentResolver, hosts []string) {
	do = ex
	newContentAgent(handler, r)
}

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
	Append     func(name Urn, body []byte, version int) *aspect.Status
}

// Resolver -
var Resolver = func() *IResolver {
	return &IResolver{
		Get: func(name Urn, version int) ([]byte, *aspect.Status) {
			return contentAgent.get(name, version)
		},
		GetRelated: func(name Urn, version int) ([]byte, *aspect.Status) {
			rel, status := relationCache.get(name)
			if !status.OK() {
				return nil, status
			}
			return contentAgent.get(rel.Thing2, version)
		},
		Append: func(name Urn, body []byte, version int) *aspect.Status {
			return storeAppend(name, body, version)
		},
	}
}()

// Get - generic typed get
func Get[T any](name Urn, version int) (T, *aspect.Status) {
	var t T
	body, status := Resolver.Get(name, version)
	if !status.OK() {
		return t, status
	}
	err := json.Unmarshal(body, &t)
	if err != nil {
		return t, aspect.NewStatusError(aspect.StatusJsonEncodeError, err)
	}
	return t, aspect.StatusOK()
}

// GetRelated - generic typed get
func GetRelated[T any](name Urn, version int) (T, *aspect.Status) {
	var t T

	rel, status := relationCache.get(name)
	if !status.OK() {
		return t, status
	}
	return Get[T](rel.Thing2, version)
}
