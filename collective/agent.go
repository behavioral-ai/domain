package collective

import (
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

func (a *agentT) getContent(name string, version int) ([]byte, *messaging.Status) {
	buf, status := a.cache.get(name, version)
	if status.OK() {
		return buf, status
	}
	if status.Code == http.StatusBadRequest {
		//a.notify(status)
		return nil, status
	}
	buf, status = a.resolver(http.MethodGet, name, "", nil, version)
	if !status.OK() {
		//a.notify(status)
		return nil, status
	}
	status = a.cache.put(name, buf, version)
	if !status.OK() {
		//a.notify(status)
		return nil, status
	}
	return buf, messaging.StatusOK()
}

func (a *agentT) putContent(name, author string, buf []byte, version int) *messaging.Status {
	_, status := a.resolver(http.MethodPut, name, author, buf, version)
	if !status.OK() {
		//a.notify(status)
		return status
	}
	status = a.cache.put(name, buf, version)
	if !status.OK() {
		//a.notify(status)
	}
	return status
}

func (a *agentT) getMap(name string) (map[string]string, *messaging.Status) {

	return nil, messaging.StatusNotFound()
}

func (a *agentT) putMap(name, author string, m map[string]string) *messaging.Status {

	return messaging.StatusBadRequest()
}
