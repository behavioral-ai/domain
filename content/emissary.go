package content

import (
	"github.com/behavioral-ai/core/messaging"
)

// emissary attention
func emissaryAttend(agent *agentT) {
	agent.dispatch(agent.emissary, messaging.StartupEvent)
	var paused = false
	if paused {
	}
	agent.ticker.Start(-1)
	for {
		select {
		case <-agent.ticker.C():
			agent.dispatch(agent.ticker, messaging.TickEvent)
		default:
		}
		select {
		case msg := <-agent.emissary.C:
			agent.dispatch(agent.emissary, msg.Event())
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.emissaryFinalize()
				return
			default:
			}
		default:
		}
	}
}
