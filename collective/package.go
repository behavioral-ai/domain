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
	return agent.load(dir)
}

type Relation struct {
	Thing1 string `json:"thing1"`
	Thing2 string `json:"thing2"`
}

// Appender - append
type Appender struct {
	Thing    func(name, author, tags string) error
	Frame    func(name, author, tags string, aspects []Relation, version int) error
	Likeness func(name, author, tags string, terms map[string]string) error
	Guidance func(name, author, tags string, text string) error
	Relation func(name1, name2, author, tags string) error
	Activity func(agent, event, location string, terms map[string]string) error
}

// Append -
var Append = func() *Appender {
	return &Appender{
		Thing: func(name, author, tags string) error {
			return errors.New("error: not implemented")
		},
		Frame: func(name, author, tags string, aspects []Relation, version int) error {
			return errors.New("error: not implemented")
		},
		Likeness: func(name, author, tags string, terms map[string]string) error {
			return errors.New("error: not implemented")
		},
		Guidance: func(name, author, tags, text string) error {
			return errors.New("error: not implemented")
		},
		Relation: func(name1, name2, author, tags string) error {
			return errors.New("error: not implemented")
		},
		Activity: func(agent, event, location string, terms map[string]string) error {
			return errors.New("error: not implemented")
		},
	}
}()

// Resolution - resolution
type Resolution struct {
	Get func(name string, version int) ([]byte, error)
	Put func(name, author string, content any, version int) error
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Get: func(name string, version int) ([]byte, error) {
			return agent.resolverGet(name, version)
		},
		Put: func(name, author string, content any, version int) error {
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
			return agent.resolverPut(name, author, buf, version)
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
