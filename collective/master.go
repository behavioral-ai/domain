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
	//comms := agent.master
	//comms.dispatch(agent, messaging.StartupEvent)

	for {
		// message processing
		select {
		case msg := <-agent.master.C:
			//comms.setup(agent, msg.Event())
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				//comms.finalize()
				//comms.dispatch(agent, msg.Event())
				return

			default:
				agent.Notify(messaging.NewStatusError(messaging.StatusInvalidContent, errors.New("invalid message"))) //messaging.EventError(agent.Uri(), msg))
			}
			//comms.dispatch(agent, msg.Event())
		default:
		}
	}
}
