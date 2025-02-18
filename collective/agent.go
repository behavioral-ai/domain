package collective

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	Class           = "content.agent"
	defaultDuration = time.Second * 10
)

type agentT struct {
	running   bool
	ephemeral bool
	agentId   string
	duration  time.Duration
	cache     *contentT
	resolver  resolutionFunc

	handler  messaging.OpsAgent
	emissary *messaging.Channel
	master   *messaging.Channel
}

func contentAgentUri() string {
	return Class
}

func newHttpAgent(handler messaging.OpsAgent) *agentT {
	return newContentAgent(handler, httpResolution, false)
}

func newEphemeralAgent(handler messaging.OpsAgent) *agentT {
	return newContentAgent(handler, fileResolution, true)
}

func newContentAgent(handler messaging.OpsAgent, resolver resolutionFunc, ephemeral bool) *agentT {
	a := new(agentT)
	a.ephemeral = ephemeral
	a.agentId = contentAgentUri()
	a.duration = defaultDuration
	a.cache = newContentCache()
	a.resolver = resolver
	a.handler = handler
	a.emissary = messaging.NewEmissaryChannel(true)
	a.master = messaging.NewMasterChannel(true)
	return a
}

// String - identity
func (s *agentT) String() string { return s.Uri() }

// Uri - agent identifier
func (s *agentT) Uri() string { return s.agentId }

// Message - message the agent
func (s *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.To() {
	case messaging.EmissaryChannel:
		s.emissary.Send(m)
	case messaging.MasterChannel:
		s.master.Send(m)
	default:
		s.emissary.Send(m)
	}
}

// Run - run the agent
func (s *agentT) Run() {
	if s.running {
		return
	}
	go masterAttend(s)
	go emissaryAttend(s)
	s.running = true
}

// Shutdown - shutdown the agent
func (s *agentT) Shutdown() {
	if !s.running {
		return
	}
	s.running = false
	msg := messaging.NewControlMessage(s.Uri(), s.Uri(), messaging.ShutdownEvent)
	s.emissary.Enable()
	s.emissary.Send(msg)
	s.master.Enable()
	s.master.Send(msg)
}

func (s *agentT) IsFinalized() bool {
	return true
}

func (s *agentT) load(dir string) error {
	if dir == "" {
		return nil
	}
	return load(dir)
}

func (s *agentT) resolverGet(name string, version int) ([]byte, error) {
	buf, err := cache.get(name, version)
	if err == nil {
		return buf, err
	}
	buf, err = s.resolver(http.MethodGet, name, "", nil, version)
	if err != nil {
		return nil, err
	}
	err = s.cache.put(name, buf, version)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (s *agentT) resolverPut(name, author string, buf []byte, version int) error {
	_, err := s.resolver(http.MethodPut, name, author, buf, version)
	if err != nil {
		return err
	}
	return cache.put(name, buf, version)
}
