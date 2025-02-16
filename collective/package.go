package collective

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

// Applications can create as many domains/NISD as needed
// "agent" is the reserved domain for the agent collective supporting agent development

const (
	ResourceUri = "urn:collective"

	AgentNID = "agent" // Restricted NID/Domain

	ThingNSS  = "thing"  // urn:{NID}:thing:{module-package}:{type}
	AspectNSS = "aspect" // urn:{NID}:aspect:{path}

)

// HttpExchange - exchange type
type HttpExchange func(r *http.Request) (*http.Response, error)

var (
	do = func(r *http.Request) (*http.Response, error) {
		return nil, errors.New("error: Collective HttpExchange function has not be initialized")
	}
)

// Initialize - collective initialize, hosts are service hosts for cloud collective
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

// Append - append
type Append struct {
	Thing    func(name, cn string) error
	Relation func(thing1, thing2 string) error
}

// Appender -
var Appender = func() *Append {
	return &Append{
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

// Resolve - resolution
type Resolve struct {
	Get func(name string, version int) ([]byte, error)
	Put func(name string, content any, version int) error
}

// Resolver -
var Resolver = func() *Resolve {
	return &Resolve{
		Get: func(name string, version int) ([]byte, error) {
			return contentAgent.resolve(name, version)
		},
		Put: func(name string, content any, version int) error {
			var buf []byte
			if name == "" || content == nil || version <= 0 {
				return errors.New(fmt.Sprintf("error: invalid argument name %v content %v version %v", name, content, version))
			}
			switch ptr := content.(type) {
			case string:
				buf = []byte(ptr)
			case []byte:
				buf = ptr
			default:
				var err error

				buf, err = json.Marshal(ptr)
				if err != nil {
					return err
				}
			}
			return storeAppend(name, buf, version)
		},
	}
}()

// Get - generic typed get
func Get[T any](name string, version int, resolver *Resolve) (T, error) {
	var t T

	if resolver == nil {
		resolver = Resolver
	}
	body, status := resolver.Get(name, version)
	if status != nil {
		return t, status
	}
	switch ptr := any(&t).(type) {
	case *string:
		*ptr = string(body)
	case *[]byte:
		*ptr = body
	default:
		err := json.Unmarshal(body, ptr)
		if err != nil {
			return t, errors.New(fmt.Sprintf("error: JsonEncode %v", err))
		}
	}
	return t, nil
}
