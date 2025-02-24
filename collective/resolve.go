package collective

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

// ResolutionFunc - data store function
type resolutionFunc func(method, name, author string, body []byte, version int) ([]byte, *messaging.Status)

type resolution struct {
	do       HttpExchange
	notifier messaging.NotifyFunc
	activity AddActivityFunc
	hosts    []string
	agent    *agentT
}

func newHttpResolver() Resolution {
	r := new(resolution)
	r.agent = newContentAgent(false, nil)
	return r
}

// GetContent - resolution get
func (r *resolution) GetContent(name string, version int) ([]byte, *messaging.Status) {
	if name == "" || version <= 0 {
		status := messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version)), "", nil)
		r.agent.notify(status)
		return nil, status
	}
	return r.agent.resolverGetContent(name, version)
}

// PutContent - resolution put
func (r *resolution) PutContent(name, author string, content any, version int) *messaging.Status {
	if name == "" || content == nil || version <= 0 {
		status := messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v content %v version %v", name, content, version)), "", nil)
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
			status := messaging.NewStatusError(messaging.StatusJsonEncodeError, err, "", nil)
			r.agent.notify(status)
			return status
		}
	case []byte:
		buf = ptr
	default:
		var err error

		buf, err = json.Marshal(ptr)
		if err != nil {
			status := messaging.NewStatusError(messaging.StatusJsonEncodeError, err, "", nil)
			r.agent.notify(status)
			return status
		}
	}
	return r.agent.resolverPutContent(name, author, buf, version)
}

// GetMap - resolution get
func (r *resolution) GetMap(name string) (map[string]string, *messaging.Status) {
	if name == "" {
		status := messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name is empty")), "", nil)
		r.agent.notify(status)
		return nil, status
	}
	return nil, messaging.StatusNotFound()
}

// PutMap - resolution put
func (r *resolution) PutMap(name, author string, m map[string]string) *messaging.Status {
	if name == "" || m == nil {
		status := messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument, name or map is empty")), "", nil)
		r.agent.notify(status)
		return status
	}
	return messaging.StatusBadRequest()
}

// AddActivity - resolution activity
func (r *resolution) AddActivity(agent messaging.Agent, event, source string, content any) {
	if r.activity != nil {
		r.activity(agent, event, source, content)
	} else {
		// TODO: add call to append activity
	}
}

// Notify - resolution notify
func (r *resolution) Notify(e messaging.Event) {
	if r.notifier != nil {
		r.notifier(e)
	} else {
		// TODO: add call to notify
	}
}
