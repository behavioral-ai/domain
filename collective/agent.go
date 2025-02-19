package collective

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	agentUri        = "root:agent/domain/collective"
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

	handler  messaging.OpsAgent
	emissary *messaging.Channel
	master   *messaging.Channel
}

func newContentAgent(ephemeral bool) *agentT {
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

func (a *agentT) load(dir string) error {
	if dir == "" {
		return nil
	}
	return loadContent(a.cache, dir)
}

func (a *agentT) resolverGet(name string, version int) ([]byte, error) {
	buf, err := a.cache.get(name, version)
	if err == nil {
		return buf, err
	}
	buf, err = a.resolver(http.MethodGet, name, "", nil, version)
	if err != nil {
		return nil, err
	}
	err = a.cache.put(name, buf, version)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (a *agentT) resolverPut(name, author string, buf []byte, version int) error {
	_, err := a.resolver(http.MethodPut, name, author, buf, version)
	if err != nil {
		return err
	}
	return a.cache.put(name, buf, version)
}
