package collective

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
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
	notifyFn  messaging.NotifyFunc

	emissary *messaging.Channel
	master   *messaging.Channel
}

func newContentAgent(ephemeral bool, notify messaging.NotifyFunc) *agentT {
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
	a.notifyFn = notify
	a.emissary = messaging.NewEmissaryChannel(true)
	a.master = messaging.NewMasterChannel(true)
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return a.agentId }

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

// Notify - notifier
func (a *agentT) Notify(status *messaging.Status) {
	if a.notifyFn != nil {
		a.notifyFn(status)
	}
}

// Trace - activity tracing
//func (a *agentT) Trace(agent messaging.Agent, channel, event, activity string) {}

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

func (a *agentT) IsFinalized() bool {
	return true
}

func (a *agentT) load(dir string) *messaging.Status {
	if dir == "" {
		return messaging.StatusOK()
	}
	err := loadContent(a.Notify, a.cache, dir)
	if err != nil {
		return messaging.StatusBadRequest()
	}
	return messaging.StatusOK()
}

func (a *agentT) resolverGet(name string, version int) ([]byte, *messaging.Status) {
	buf, status := a.cache.get(name, version)
	if status.OK() {
		return buf, status
	}
	if status.Code == http.StatusBadRequest {
		a.Notify(status)
		return nil, status
	}
	buf, status = a.resolver(http.MethodGet, name, "", nil, version)
	if !status.OK() {
		a.Notify(status)
		return nil, status
	}
	status = a.cache.put(name, buf, version)
	if !status.OK() {
		a.Notify(status)
		return nil, status
	}
	return buf, messaging.StatusOK()
}

func (a *agentT) resolverPut(name, author string, buf []byte, version int) *messaging.Status {
	_, status := a.resolver(http.MethodPut, name, author, buf, version)
	if !status.OK() {
		a.Notify(status)
		return status
	}
	status = a.cache.put(name, buf, version)
	if !status.OK() {
		a.Notify(status)
	}
	return status
}
