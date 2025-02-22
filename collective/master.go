package collective

import (
	"errors"
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
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.dispatch(agent.emissary, msg.Event())
				return
			default:
				agent.notify(messaging.NewStatusError(messaging.StatusInvalidContent, errors.New("invalid message"))) //messaging.EventError(agent.Uri(), msg))
			}
			agent.dispatch(agent.emissary, msg.Event())
		default:
		}
	}
}
