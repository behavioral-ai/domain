package collective

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

// InitializedEphemeralResolver - in memory resolver, initialized with state
//func InitializedEphemeralResolver(dir string) Resolution {}

// initializedEphemeralResolver - in memory resolver, initialized with state
func initializedEphemeralResolver(dir string) Resolution {
	r := new(resolution)
	r.notifier = messaging.Notify
	r.activity = func(hostName string, agent messaging.Agent, event, source string, content any) {
		fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", messaging.FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), event, source, content)
	}
	r.agent = newContentAgent(true, nil)
	r.agent.notifier = r.notifier
	if dir != "" {
		r.agent.load(dir)
	}
	r.agent.Run()
	return r
}
