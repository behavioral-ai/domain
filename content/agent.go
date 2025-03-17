package content

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	AgentNamespaceName = "resiliency:agent/behavioral-ai/domain/content"
	agentUri           = AgentNamespaceName
	defaultDuration    = time.Second * 10
)

type agentT struct {
	running   bool
	ephemeral bool
	agentId   string
	hostName  string
	uri       []string
	duration  time.Duration
	cache     *contentT
	resolver  resolutionFunc

	ticker     *messaging.Ticker
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
		a.resolver = ephemeralResolution
	} else {
		a.resolver = httpResolution
	}
	a.ticker = messaging.NewTicker(messaging.Emissary, a.duration)
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
func (a *agentT) Name() string { return AgentNamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Channel() {
	case messaging.Emissary:
		a.emissary.Send(m)
	case messaging.Master:
		a.master.Send(m)
	case messaging.Control:
		a.emissary.Send(m)
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
	if !a.emissary.IsClosed() {
		a.emissary.Send(messaging.Shutdown)
	}
	if !a.master.IsClosed() {
		a.master.Send(messaging.Shutdown)
	}
}

func (a *agentT) notify(e messaging.NotifyItem) {
	a.notifier(e)
}

func (a *agentT) dispatch(channel any, event string) {
	messaging.Dispatch(a, a.dispatcher, channel, event)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterFinalize() {
	a.master.Close()
}

func (a *agentT) getValue(name string, version int) (buf []byte, status *messaging.Status) {
	if name == "" || version <= 0 {
		return nil, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version)), a.Uri())
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

func (a *agentT) addValue(name, author string, buf []byte, version int) *messaging.Status {
	if name == "" || author == "" || buf == nil || version <= 0 {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version)), a.Uri())
	}
	_, status := a.resolver(http.MethodPut, name, author, buf, version)
	if !status.OK() {
		status.SetAgent(a.Uri())
		status.SetMessage(fmt.Sprintf("name %v and version %v", name, version))
		return status
	}
	a.cache.put(name, buf, version)
	return status
}
