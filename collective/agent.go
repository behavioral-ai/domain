package collective

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	Name            = "resiliency:agent/domain/collective/content"
	agentUri        = Name
	defaultDuration = time.Second * 10
)

type agentT struct {
	running   bool
	ephemeral bool
	agentId   string
	uri       []string
	duration  time.Duration
	cache     *contentT
	mapCache  *mapT
	resolver  resolutionFunc

	emissary   *messaging.Channel
	master     *messaging.Channel
	notifier   messaging.NotifyFunc
	dispatcher messaging.Dispatcher
}

func newContentAgent(ephemeral bool, dispatcher messaging.Dispatcher) *agentT {
	a := new(agentT)
	a.ephemeral = ephemeral
	a.agentId = agentUri
	a.duration = defaultDuration
	a.cache = newContentCache()
	a.mapCache = newMapCache()
	if ephemeral {
		a.resolver = fileResolution
	} else {
		a.resolver = httpResolution
	}

	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	a.dispatcher = dispatcher
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return a.agentId }

// Name - agent name
func (a *agentT) Name() string { return Name }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Channel() {
	case messaging.EmissaryChannel:
		a.emissary.Send(m)
	case messaging.MasterChannel:
		a.master.Send(m)
	default:
		a.emissary.Send(m)
	}
}

// Run - run the agent
func (a *agentT) Run() {
	if a.running {
		return
	}
	go masterAttend(a)
	go emissaryAttend(a)
	a.running = true
}

// Shutdown - shutdown the agent
func (a *agentT) Shutdown() {
	//if !a.running {
	//	return
	//}
	//a.running = false
	//a.emissary.Enable()
	if !a.emissary.IsClosed() {
		a.emissary.Send(messaging.Shutdown)
	}
	//a.master.Enable()
	if !a.master.IsClosed() {
		a.master.Send(messaging.Shutdown)
	}
}

func (a *agentT) notify(e messaging.Event) {
	a.notifier(e)
}

func (a *agentT) dispatch(channel any, event string) {
	messaging.Dispatch(a, a.dispatcher, channel, event)
}

func (a *agentT) load(dir string) *messaging.Status {
	if dir == "" {
		return messaging.StatusOK()
	}
	err := loadContent(a.cache, dir)
	if err != nil {
		status := messaging.NewStatusError(messaging.StatusIOError, err, "", a.Uri())
		a.notify(status)
		return status
	}
	return messaging.StatusOK()
}

func (a *agentT) getContent(name string, version int) (buf []byte, status *messaging.Status) {
	if name == "" || version <= 0 {
		return nil, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version)), "", a.Uri())
	}
	var err error
	buf, err = a.cache.get(name, version)
	if err == nil {
		return buf, messaging.StatusOK()
	}
	// Cache miss
	buf, status = a.resolver(http.MethodGet, name, "", nil, version)
	if !status.OK() {
		status.SetAgent(a.Uri())
		status.SetMessage(fmt.Sprintf("name %v and version %v", name, version))
		return nil, status
	}
	a.cache.put(name, buf, version)
	return buf, messaging.StatusOK()
}

func (a *agentT) putContent(name, author string, buf []byte, version int) *messaging.Status {
	if name == "" || author == "" || buf == nil || version <= 0 {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version)), "", a.Uri())
	}
	_, status := a.resolver(http.MethodPut, name, author, buf, version)
	if !status.OK() {
		return status.SetAgent(a.Uri())
	}
	a.cache.put(name, buf, version)
	return status
}

func (a *agentT) getMap(name string) (map[string]string, *messaging.Status) {
	if name == "" {
		return nil, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("map name [%v] is empty", name)), "", a.Uri())
	}
	m, err := a.mapCache.get(name)
	if err == nil {
		return m, messaging.StatusOK()
	}
	// Cache miss
	buf, status := a.resolver(http.MethodGet, name, "", nil, 1)
	if !status.OK() {
		status.SetAgent(a.Uri())
		status.SetMessage(fmt.Sprintf("map name [%v] not found", name))
		return nil, status
	}
	// TODO : parse buf into map
	if len(buf) > 0 {
	}
	return nil, messaging.StatusNotFound().SetAgent(a.Uri())
}

func (a *agentT) putMap(name, author string, m map[string]string) *messaging.Status {
	if name == "" || author == "" || m == nil {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid argument name [%v],author [%v] or map", name, author)), "", a.Uri())
	}
	/*
		_, status := a.resolver(http.MethodPut, name, author, nil, 1)
		if !status.OK() {
			return status.SetAgent(a.Uri())
		}
	*/
	err := a.mapCache.put(name, m)
	if err == nil {
		return messaging.StatusOK()
	}
	return messaging.StatusBadRequest().SetAgent(a.Uri())
}
