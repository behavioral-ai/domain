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

var (
	agent *agentT
	cache = newContentCache()
	do    = func(r *http.Request) (*http.Response, error) {
		return nil, errors.New("error: Collective HttpExchange function has not be initialized")
	}
)

// ResolutionKey -
type ResolutionKey struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

// HttpExchange - exchange type
type HttpExchange func(r *http.Request) (*http.Response, error)

// Initialize - collective initialize, hosts are service hosts for cloud collective
func Initialize(handler messaging.OpsAgent, ex HttpExchange, hosts []string) error {
	if ex == nil || handler == nil || len(hosts) == 0 {
		return errors.New("error: bad request, handler, exchange, or hosts are empty")
	}
	// Where to set hosts??
	do = ex
	agent = newHttpAgent(handler)
	return nil
}

func InitializeEphemeral(handler messaging.OpsAgent, dir string) error {
	if handler == nil || dir == "" {
		return errors.New("error: bad request, handler or dir is empty")
	}
	agent = newEphemeralAgent(handler)
	err := agent.load(dir)
	if err != nil {
		return err
	}

	return nil
}

// Appender - append
type Appender struct {
	Thing    func(name, cn string) error
	Relation func(thing1, thing2 string) error
}

// Append -
var Append = func() *Appender {
	return &Appender{
		Thing: func(name, cn string) error {
			buf, err := json.Marshal(thing{Name: name, Cn: cn})
			if err != nil {
				return err
			}
			return agent.put(name, buf, 1)
		},
		Relation: func(thing1, thing2 string) error {
			buf, err := json.Marshal(relation{Thing1: thing1, Thing2: thing2})
			if err != nil {
				return err
			}
			return agent.put("relation", buf, 1)
		},
	}
}()

// Resolution - resolution
type Resolution struct {
	Get func(name string, version int) ([]byte, error)
	Put func(name string, content any, version int) error
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Get: func(name string, version int) ([]byte, error) {
			return agent.get(name, version)
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
			return agent.put(name, buf, version)
		},
	}
}()

// Resolve - generic typed resolution
func Resolve[T any](name string, version int) (T, error) {
	var t T

	body, status := Resolver.Get(name, version)
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
