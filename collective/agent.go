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
	if ephemeral {
		a.resolver = fileResolution
	} else {
		a.resolver = httpResolution
	}

	a.emissary = messaging.NewEmissaryChannel(true)
	a.master = messaging.NewMasterChannel(true)
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
	if !a.running {
		return
	}
	a.running = false
	a.emissary.Enable()
	a.emissary.Send(messaging.Shutdown)
	a.master.Enable()
	a.master.Send(messaging.Shutdown)
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
		status := messaging.NewStatusError(messaging.StatusIOError, err, "", a)
		a.notify(status)
		return status
	}
	return messaging.StatusOK()
}

func (a *agentT) getContent(name string, version int) (buf []byte, status *messaging.Status) {
	if name == "" || version <= 0 {
		return nil, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version)), "", a)
	}
	var err error
	buf, err = a.cache.get(name, version)
	if err == nil {
		return buf, messaging.StatusOK()
	}
	// Cache miss
	buf, status = a.resolver(http.MethodGet, name, "", nil, version)
	if !status.OK() {
		status.AgentUri = a.Uri()
		status.Msg = fmt.Sprintf("name %v and version %v", name, version)
		return nil, status
	}
	a.cache.put(name, buf, version)
	return buf, messaging.StatusOK()
}

func (a *agentT) putContent(name, author string, buf []byte, version int) *messaging.Status {
	if name == "" || author == "" || buf == nil || version <= 0 {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version)), "", a)
	}
	_, status := a.resolver(http.MethodPut, name, author, buf, version)
	if !status.OK() {
		status.AgentUri = a.Uri()
		return status
	}
	a.cache.put(name, buf, version)
	return status
}

func (a *agentT) getMap(name string) (map[string]string, *messaging.Status) {
	if name == "" {
		return nil, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v", name)), "", a)
	}
	return nil, messaging.StatusNotFound()
}

func (a *agentT) putMap(name, author string, m map[string]string) *messaging.Status {
	if name == "" || author == "" || m == nil {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v or map", name)), "", a)
	}
	return messaging.StatusBadRequest()
}
