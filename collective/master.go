package collective

import (
	"github.com/behavioral-ai/core/messaging"
)

// master attention
func masterAttend(agent *agentT) {
	paused := false
	if paused {
	}
	agent.dispatch(agent.emissary, messaging.StartupEvent)

	for {
		select {
		case msg := <-agent.master.C:
			agent.dispatch(agent.emissary, msg.Event())
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.dispatch(agent.emissary, msg.Event())
				return
			default:
			}
		default:
		}
	}
}
