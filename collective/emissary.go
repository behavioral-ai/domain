package collective

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

// emissary attention
func emissaryAttend(agent *agentT) {
	var paused = false
	if paused {
	}
	ticker := messaging.NewPrimaryTicker(agent.duration)
	agent.dispatch(agent.emissary, messaging.StartupEvent)

	ticker.Start(-1)
	for {
		select {
		case <-ticker.C():
			agent.dispatch(ticker, messaging.TickEvent)
		default:
		}
		select {
		case msg := <-agent.emissary.C:
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				ticker.Stop()
				agent.dispatch(agent.emissary, msg.Event())
				return
			default:
				agent.notify(messaging.NewStatusError(messaging.StatusInvalidContent, errors.New(fmt.Sprintf("%v %v", agent.Uri(), msg.Event())))) //messaging.EventError(agent.Uri(), msg))
			}
			agent.dispatch(agent.emissary, msg.Event())
		default:
		}
	}
}
