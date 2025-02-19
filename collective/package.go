package collective

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

// Applications can create as many domains/NISD as needed

const (
	ResourceUri = "urn:collective"
	AgentNSS    = "agent"    // urn:{NID}:thing:{module-package}:{type}
	ThingNSS    = "thing"    // urn:{NID}:thing:{module-package}:{type}
	AspectNSS   = "aspect"   // urn:{NID}:aspect:{path}
	FrameNSS    = "frame"    // urn:{NID}:frame:{path}
	LikenessNSS = "likeness" // urn:{NID}:likeness:{path}
	GuidanceNSS = "guidance" // urn:{NID}:guidance:{path}
	EventNSS    = "event"    // urn:{NID}:event:{path}

)

// Append -
var (
	Append = newAppender()
)

// Relation -
type Relation struct {
	Thing1 string `json:"thing1"`
	Thing2 string `json:"thing2"`
}

// Appender - collective append
type Appender interface {
	Thing(name, author, tags string) error
	Relation(name1, name2, author, tags string) error
	Frame(name, author, tags string, aspects []Relation, version int) error
	Likeness(name, author, tags string, terms map[string]string) error
	Guidance(name, author, tags string, text string) error
	Activity(agent messaging.Agent, event, location string, terms map[string]string) error
}

type appender struct{}

// newAppender -
func newAppender() Appender {
	a := new(appender)
	return a
}

// Thing - append a thing
func (a *appender) Thing(name, author, tags string) error {
	return errors.New("error: not implemented")
}

// Relation - append a relation
func (a *appender) Relation(name1, name2, author, tags string) error {
	return errors.New("error: not implemented")
}

// Frame - append a frame
func (a *appender) Frame(name, author, tags string, aspects []Relation, version int) error {
	return errors.New("error: not implemented")
}

// Likeness - append a likeness
func (a *appender) Likeness(name, author, tags string, terms map[string]string) error {
	return errors.New("error: not implemented")
}

// Guidance - append guidance
func (a *appender) Guidance(name, author, tags, text string) error {
	return errors.New("error: not implemented")
}

// Activity - append activity
func (a *appender) Activity(agent messaging.Agent, event, location string, terms map[string]string) error {
	return errors.New("error: not implemented")
}

// Resolver -
var (
	Resolver = newHttpResolver()
)

// ResolutionKey -
type ResolutionKey struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

// HttpExchange - exchange type
type HttpExchange func(r *http.Request) (*http.Response, error)

// StartupResolver - run the content agent
func StartupResolver(uri []string, do HttpExchange) {
	if r, ok := any(Resolver).(*resolution); ok {
		if do != nil {
			r.do = do
		}
		r.agent.uri = uri
		r.agent.Run()
	}
}

// Resolution - of things in the real world
type Resolution interface {
	Get(name string, version int) ([]byte, error)
	Put(name, author string, content any, version int) error
}

type resolution struct {
	do    HttpExchange
	hosts []string
	agent *agentT
}

func newHttpResolver() Resolution {
	r := new(resolution)
	r.agent = newContentAgent(false)
	return r
}

// NewEphemeralResolver - in memory resolver
func NewEphemeralResolver(dir string) (Resolution, error) {
	r := new(resolution)
	r.agent = newContentAgent(true)
	err := r.agent.load(dir)
	r.agent.Run()
	return r, err
}

// Get - resolution get
func (r *resolution) Get(name string, version int) ([]byte, error) {
	if name == "" || version <= 0 {
		return nil, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version))
	}
	return r.agent.resolverGet(name, version)
}

// Put - resolution put
func (r *resolution) Put(name, author string, content any, version int) error {
	if name == "" || content == nil || version <= 0 {
		return errors.New(fmt.Sprintf("error: invalid argument name %v content %v version %v", name, content, version))
	}
	return resolverPut(r, name, author, content, version)
}

// Resolve - generic typed resolution
func Resolve[T any](name string, version int, resolver Resolution) (T, error) {
	var t T

	if resolver == nil {
		return t, errors.New("error: BadRequest - resolver is nil")
	}
	body, status := resolver.Get(name, version)
	if status != nil {
		return t, status
	}
	switch ptr := any(&t).(type) {
	case *string:
		t1, err := Resolve[text](name, version, resolver)
		if err != nil {
			return t, err
		}
		*ptr = t1.Value
	case *[]byte:
		*ptr = body
	default:
		err := json.Unmarshal(body, ptr)
		if err != nil {
			return t, errors.New(fmt.Sprintf("error: JsonEncode - %v", err))
		}
	}
	return t, nil
}

/*
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
*/
