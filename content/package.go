package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	NsNameKey  = "name"
	VersionKey = "ver"
)

// Resolution - in the real world
type Resolution interface {
	GetValue(nsName string, version int) ([]byte, *messaging.Status)
	AddValue(nsName, author string, content any, version int) *messaging.Status
	AddActivity(agent messaging.Agent, event, source string, content any)
	Notify(e messaging.Event)
}

// Resolver - content resolution in the real world
var (
	Resolver = newHttpResolver()
	Agent    messaging.Agent
)

func init() {
	if r, ok := any(Resolver).(*resolution); ok {
		// Testing
		r.notifier = messaging.Notify
		r.agent.notifier = r.notifier
		r.activity = func(hostName string, agent messaging.Agent, event, source string, content any) {
			fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", messaging.FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), event, source, content)
		}
		r.agent.Run()
		Agent = r.agent
	}
}

// Resolve - generic typed resolution
func Resolve[T any](nsName string, version int, resolver Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", nsName)), AgentNamespaceName)
	}
	body, status := resolver.GetValue(nsName, version)
	if !status.OK() {
		return t, status
	}
	if len(body) == 0 {
		return t, messaging.NewStatusMessage(http.StatusNoContent, fmt.Sprintf("content not found for name: %v", nsName), AgentNamespaceName)
	}
	switch ptr := any(&t).(type) {
	case *string:
		t1, status1 := Resolve[text](nsName, version, resolver)
		if !status1.OK() {
			return t, status1
		}
		*ptr = t1.Value
	case *[]byte:
		*ptr = body
	default:
		err := json.Unmarshal(body, ptr)
		if err != nil {
			return t, messaging.NewStatusError(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, nsName)), AgentNamespaceName)
		}
	}
	return t, messaging.StatusOK()
}

// NewEphemeralResolver - in memory resolver
func NewEphemeralResolver() Resolution {
	return initializedEphemeralResolver(true, true)
}

// NewConfigEphemeralResolver - in memory resolver
func NewConfigEphemeralResolver(activity, notify bool) Resolution {
	return initializedEphemeralResolver(activity, notify)
}
