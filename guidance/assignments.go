package guidance

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
)

// Assignments - assignments functions struct
type Assignments struct {
	All func(origin common.Origin) ([]HostEntry, *messaging.Status)
	New func(origin common.Origin) ([]HostEntry, *messaging.Status)
}

var Assign = func() *Assignments {
	return &Assignments{
		All: func(origin common.Origin) ([]HostEntry, *messaging.Status) {
			entry, status := GetRegion(origin)
			return entry, status
		},
		New: func(origin common.Origin) ([]HostEntry, *messaging.Status) {
			return nil, messaging.StatusNotFound()
		},
	}
}()
