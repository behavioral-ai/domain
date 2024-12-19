package guidance

import (
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/messaging"
)

// Assignments - assignments functions struct
type Assignments struct {
	All func(h messaging.Notifier, origin core.Origin) ([]HostEntry, *core.Status)
	New func(h messaging.Notifier, origin core.Origin) ([]HostEntry, *core.Status)
}

var Assign = func() *Assignments {
	return &Assignments{
		All: func(h messaging.Notifier, origin core.Origin) ([]HostEntry, *core.Status) {
			entry, status := GetRegion(origin)
			if !status.OK() {
				h.Notify(status)
			}
			return entry, status
		},
		New: func(h messaging.Notifier, origin core.Origin) ([]HostEntry, *core.Status) {
			return nil, core.StatusNotFound()
		},
	}
}()
