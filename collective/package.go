package collective

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	ResourceUri       = "urn:collective"
	ResolutionKeyName = "resolution-key"
)

// HttpExchange - exchange type
type HttpExchange func(r *http.Request) (*http.Response, error)

// Resolver - collective resolution in the real world
var (
	appHostName string
	Resolver    = newHttpResolver()
)

func init() {
	if r, ok := any(Resolver).(*resolution); ok {
		// Testing
		r.notifier = messaging.Notify
		r.agent.notifier = r.notifier
		r.activity = func(hostName string, agent messaging.Agent, event, source string, content any) {
			fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", messaging.FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), event, source, content)
		}
	}
}

// Startup - run the agents
func Startup(uri []string, do HttpExchange, hostName string) {
	appHostName = hostName
	if r, ok := any(Resolver).(*resolution); ok {
		if do != nil {
			r.do = do
		}
		r.agent.uri = uri
		r.agent.Run()
	}
}

func Shutdown() {

}

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
func NewEphemeralResolver() Resolution {
	return initializedEphemeralResolver("")
}

// Resolve - generic typed resolution
func Resolve[T any](name string, version int, resolver Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", name)), Name)
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
			//uri := "<nil>"
			//agent := toAgent(resolver)
			//if agent != nil {
			//	uri = agent.Uri()
			//}
			return t, messaging.NewStatusError(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, name)), Name)
		}
	}
	return t, messaging.StatusOK()
}
