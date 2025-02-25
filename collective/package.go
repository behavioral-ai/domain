package collective

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

// Applications can create as many domains/NISD as needed

const (
	ResourceUri = "urn:collective"
	AgentNSS    = "agent"    // urn:{NID}:thing:{module-package}:{type}
	ThingNSS    = "thing"    // urn:{NID}:thing:{module-package}:{type}
	AspectNSS   = "aspect"   // urn:{NID}:aspect:{path}
	FrameNSS    = "frame"    // urn:{NID}:frame:{path}
	LikenessNSS = "likeness" // urn:{NID}:likeness:{path}
	GuidanceNSS = "guidance" // urn:{NID}:guidance2:{path}
	EventNSS    = "event"    // urn:{NID}:event:{path}

	ResolutionKeyName = "resolution-key"
)

// HttpExchange - exchange type
type HttpExchange func(r *http.Request) (*http.Response, error)

// Startup - run the agents
func Startup(uri []string, do HttpExchange) {
	if r, ok := any(Resolver).(*resolution); ok {
		if do != nil {
			r.do = do
		}
		r.agent.notifier = r.Notify
		r.agent.uri = uri
		r.agent.Run()
	}
}

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
}

// Resolver - collective resolution in the real world
var (
	Resolver = newHttpResolver()
)

// ResolutionKey -
type ResolutionKey struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

// Resolution - in the real world
type Resolution interface {
	GetContent(name string, version int) ([]byte, *messaging.Status)
	PutContent(name, author string, content any, version int) *messaging.Status
	GetMap(name string) (map[string]string, *messaging.Status)
	PutMap(name, author string, m map[string]string) *messaging.Status
	AddActivity(agent messaging.Agent, event, source string, content any)
	Notify(e messaging.Event)
}

// NewEphemeralResolver - in memory resolver
func NewEphemeralResolver(dir string, notifier messaging.NotifyFunc, enableActivity bool) Resolution {
	r := new(resolution)
	if notifier == nil {
		notifier = func(e messaging.Event) {
			messaging.Notify(e)
		}
	}
	r.notifier = notifier
	r.activity = func(agent messaging.Agent, event, source string, content any) {
		if enableActivity {
			fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", messaging.FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), event, source, content)
		}
	}
	r.agent = newContentAgent(true, nil)
	r.agent.notifier = notifier
	r.agent.load(dir)
	r.agent.Run()
	return r
}

// Resolve - generic typed resolution
func Resolve[T any](name string, version int, resolver Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatusError(http.StatusBadRequest, errors.New("error: BadRequest - resolver is nil"), "", "<nil>")
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
			return t, messaging.NewStatusError(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonEncode - %v", err)), "", toAgent(resolver).Uri())
		}
	}
	return t, messaging.StatusOK()
}
