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
	Append = newHttpAppender()
)

// Appender - collective append
type Appender interface {
	Thing(name, author string, related []string) *messaging.Status
	Relation(name1, name2, author string) *messaging.Status
	Frame(name, author string, contains []string, version int) *messaging.Status
	Guidance(name, author, text string, related []string) *messaging.Status
	Activity(agent messaging.Agent, event, source string, content any)
}

type appender struct{}

// newAppender -
func newHttpAppender() Appender {
	a := new(appender)
	return a
}

// Thing - append a thing
func (a *appender) Thing(name, author string, related []string) *messaging.Status {
	return messaging.NewStatusError(http.StatusBadRequest, errors.New("error: not implemented"))
}

// Relation - append a relation
func (a *appender) Relation(name1, name2, author string) *messaging.Status {
	return messaging.NewStatusError(http.StatusBadRequest, errors.New("error: not implemented"))
}

// Frame - append a frame
func (a *appender) Frame(name, author string, contains []string, version int) *messaging.Status {
	return messaging.NewStatusError(http.StatusBadRequest, errors.New("error: not implemented"))
}

// Guidance - append guidance
func (a *appender) Guidance(name, author, text string, related []string) *messaging.Status {
	return messaging.NewStatusError(http.StatusBadRequest, errors.New("error: not implemented"))
}

// Activity - append activity
func (a *appender) Activity(agent messaging.Agent, event, location string, content any) {
	// if there is an error, use the agents notify.
	fmt.Printf("agent: %v event: %v location: %v", agent, event, location) //, terms)
	//return messaging.StatusOK() //errors.New("error: not implemented")
}

// Resolver - collective resolution in the real world
var (
	Resolver = newHttpResolver()
)

const (
	ContentMapped     = "mapped"
	ContentBinary     = "binary"
	ContentNotDefined = "undefined"
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

// Resolution - in the real world
type Resolution interface {
	GetContent(name string, version int) ([]byte, *messaging.Status)
	PutContent(name, author string, content any, version int) *messaging.Status
	GetMap(name string) (map[string]string, *messaging.Status)
	PutMap(name, author string, m map[string]string) *messaging.Status
}

type resolution struct {
	do    HttpExchange
	hosts []string
	agent *agentT
}

func newHttpResolver() Resolution {
	r := new(resolution)
	r.agent = newContentAgent(false, nil, nil)
	return r
}

// NewEphemeralResolver - in memory resolver
func NewEphemeralResolver(dir string, notify messaging.NotifyFunc) (Resolution, *messaging.Status) {
	r := new(resolution)
	r.agent = newContentAgent(true, notify, nil)
	err := r.agent.load(dir)
	r.agent.Run()
	return r, err
}

// GetContent - resolution get
func (r *resolution) GetContent(name string, version int) ([]byte, *messaging.Status) {
	if name == "" || version <= 0 {
		status := messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version)))
		r.agent.notify(status)
		return nil, status
	}
	return r.agent.resolverGetContent(name, version)
}

// PutContent - resolution put
func (r *resolution) PutContent(name, author string, content any, version int) *messaging.Status {
	if name == "" || content == nil || version <= 0 {
		status := messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v content %v version %v", name, content, version)))
		r.agent.notify(status)
		return status
	}
	var buf []byte

	switch ptr := content.(type) {
	case string:
		var err error

		v := text{ptr}
		buf, err = json.Marshal(v)
		if err != nil {
			status := messaging.NewStatusError(messaging.StatusJsonEncodeError, err)
			r.agent.notify(status)
			return status
		}
	case []byte:
		buf = ptr
	default:
		var err error

		buf, err = json.Marshal(ptr)
		if err != nil {
			status := messaging.NewStatusError(messaging.StatusJsonEncodeError, err)
			r.agent.notify(status)
			return status
		}
	}
	return r.agent.resolverPutContent(name, author, buf, version)
}

// GetMap - resolution get
func (r *resolution) GetMap(name string) (map[string]string, *messaging.Status) {
	if name == "" {
		status := messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name is empty")))
		r.agent.notify(status)
		return nil, status
	}
	return nil, messaging.StatusNotFound()
}

// PutMap - resolution put
func (r *resolution) PutMap(name, author string, m map[string]string) *messaging.Status {
	if name == "" || m == nil {
		status := messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument, name or map is empty")))
		r.agent.notify(status)
		return status
	}
	return messaging.StatusBadRequest()
}

// Resolve - generic typed resolution
func Resolve[T any](name string, version int, resolver Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatusError(http.StatusBadRequest, errors.New("error: BadRequest - resolver is nil"))
	}
	body, status := resolver.GetContent(name, version)
	if !status.OK() {
		return t, status
	}
	switch ptr := any(&t).(type) {
	case *string:
		t1, status1 := Resolve[text](name, version, resolver)
		if !status1.OK() {
			return t, status1
		}
		*ptr = t1.Value
	case *[]byte:
		*ptr = body
	default:
		err := json.Unmarshal(body, ptr)
		if err != nil {
			return t, messaging.NewStatusError(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonEncode - %v", err)))
		}
	}
	return t, messaging.StatusOK()
}

type text struct {
	Value string
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
