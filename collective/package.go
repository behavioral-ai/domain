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

// Startup - run the agents
func Startup(uri []string, do HttpExchange, hostName string) {
	appHostName = hostName
	if r, ok := any(Resolver).(*resolution); ok {
		if do != nil {
			r.do = do
		}
		r.agent.notifier = r.Notify
		r.agent.uri = uri
		r.agent.Run()
		status := loadResolver(r)
		if !status.OK() {
			fmt.Printf("error on loading Resolver: %v\n", status)
		}
	}
}

func Shutdown() {

}

// Resolver - collective resolution in the real world
var (
	appHostName string
	Resolver    = newHttpResolver()
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
func NewEphemeralResolver() Resolution {
	return InitializedEphemeralResolver("")
}

// InitializedEphemeralResolver - in memory resolver, initialized with state
func InitializedEphemeralResolver(dir string) Resolution {
	r := new(resolution)
	r.notifier = messaging.Notify
	r.activity = func(hostName string, agent messaging.Agent, event, source string, content any) {
		fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", messaging.FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), event, source, content)
	}
	r.agent = newContentAgent(true, nil)
	r.agent.notifier = r.notifier
	if dir != "" {
		r.agent.load(dir)
	}
	r.agent.Run()
	return r
}

// Resolve - generic typed resolution
func Resolve[T any](name string, version int, resolver Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatusError(http.StatusBadRequest, errors.New("error: BadRequest - resolver is nil"), "<nil>")
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
			uri := ""
			agent := toAgent(resolver)
			if agent == nil {
				uri = ""
			} else {
				uri = agent.Uri()
			}
			return t, messaging.NewStatusError(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonEncode - %v", err)), uri)
		}
	}
	return t, messaging.StatusOK()
}
