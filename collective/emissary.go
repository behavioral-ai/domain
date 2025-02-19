package collective

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

// emissary attention
func emissaryAttend(agent *agentT) {
	paused := false
	if paused {
	}
	ticker := messaging.NewPrimaryTicker(agent.duration)

	ticker.Start(-1)
	for {
		select {
		case <-ticker.C():

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
				return
			//case messaging.DataChangeEvent:
			default:
				agent.Notify(messaging.NewStatusError(messaging.StatusInvalidContent, errors.New(fmt.Sprintf("%v %v", agent.Uri(), msg.Event())))) //messaging.EventError(agent.Uri(), msg))
			}
		default:
		}
	}
}
