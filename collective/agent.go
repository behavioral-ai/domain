package collective

import (
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	Class           = "content.agent"
	defaultDuration = time.Second * 10
)

var (
	contentAgent *agentT
)

type agentT struct {
	running  bool
	agentId  string
	duration time.Duration
	cache    *contentT

	handler  messaging.OpsAgent
	emissary *messaging.Channel
	master   *messaging.Channel
}

func contentAgentUri() string {
	return Class
}

// newContentAgent - create a new agent1 agent
func newContentAgent(handler messaging.OpsAgent, r contentResolver) messaging.Agent {
	return newAgent(handler, r)
}

func newAgent(handler messaging.OpsAgent, r contentResolver) *agentT {
	a := new(agentT)
	a.agentId = contentAgentUri()
	a.duration = defaultDuration
	a.cache = newContentCache(r)
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

func (s *agentT) get(name Urn, version int) ([]byte, *aspect.Status) {
	body, status := s.cache.get(name, version)
	if !status.OK() {
		s.handler.Notify(status)
	}
	return body, status
}
