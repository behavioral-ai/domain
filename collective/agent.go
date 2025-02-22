package collective

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	Name            = "resiliency:agent/domain/collective"
	agentUri        = "root:agent/domain/collective/content"
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

func newContentAgent(ephemeral bool, notifier messaging.NotifyFunc, dispatcher messaging.Dispatcher) *agentT {
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
	a.notifier = notifier
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
	switch m.To() {
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
	msg := messaging.NewControlMessage(a.Uri(), a.Uri(), messaging.ShutdownEvent)
	a.emissary.Enable()
	a.emissary.Send(msg)
	a.master.Enable()
	a.master.Send(msg)
}

// Notify - notifier
func (a *agentT) notify(status *messaging.Status) *messaging.Status {
	if a.notifier != nil {
		a.notifier(status)
	}
	return status
}

func (a *agentT) dispatch(channel any, event string) {
	if a.dispatcher == nil || channel == nil {
		return
	}
	if ch, ok := channel.(*messaging.Channel); ok {
		a.dispatcher.Dispatch(a, ch.Name(), event)
		return
	}
	if t, ok := channel.(*messaging.Ticker); ok {
		a.dispatcher.Dispatch(a, t.Name(), event)
	}
}

func (a *agentT) load(dir string) *messaging.Status {
	if dir == "" {
		return messaging.StatusOK()
	}
	err := loadContent(a.notify, a.cache, dir)
	if err != nil {
		return messaging.StatusBadRequest()
	}
	return messaging.StatusOK()
}

func (a *agentT) resolverGetContent(name string, version int) ([]byte, *messaging.Status) {
	buf, status := a.cache.get(name, version)
	if status.OK() {
		return buf, status
	}
	if status.Code == http.StatusBadRequest {
		a.notify(status)
		return nil, status
	}
	buf, status = a.resolver(http.MethodGet, name, "", nil, version)
	if !status.OK() {
		a.notify(status)
		return nil, status
	}
	status = a.cache.put(name, buf, version)
	if !status.OK() {
		a.notify(status)
		return nil, status
	}
	return buf, messaging.StatusOK()
}

func (a *agentT) resolverPutContent(name, author string, buf []byte, version int) *messaging.Status {
	_, status := a.resolver(http.MethodPut, name, author, buf, version)
	if !status.OK() {
		a.notify(status)
		return status
	}
	status = a.cache.put(name, buf, version)
	if !status.OK() {
		a.notify(status)
	}
	return status
}
