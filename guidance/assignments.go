package guidance

import (
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
)

// Assignments - assignments functions struct
type Assignments struct {
	All func(h messaging.Notifier, origin common.Origin) ([]HostEntry, *aspect.Status)
	New func(h messaging.Notifier, origin common.Origin) ([]HostEntry, *aspect.Status)
}

var Assign = func() *Assignments {
	return &Assignments{
		All: func(h messaging.Notifier, origin common.Origin) ([]HostEntry, *aspect.Status) {
			entry, status := GetRegion(origin)
			if !status.OK() {
				h.Notify(status)
			}
			return entry, status
		},
		New: func(h messaging.Notifier, origin common.Origin) ([]HostEntry, *aspect.Status) {
			return nil, aspect.StatusNotFound()
		},
	}
}()
