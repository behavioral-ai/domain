package guidance

import (
	"errors"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
)

// Assignments - assignments functions struct
type Assignments struct {
	All func(h messaging.Notifier, origin common.Origin) ([]HostEntry, error)
	New func(h messaging.Notifier, origin common.Origin) ([]HostEntry, error)
}

var Assign = func() *Assignments {
	return &Assignments{
		All: func(h messaging.Notifier, origin common.Origin) ([]HostEntry, error) {
			entry, status := GetRegion(origin)
			if status != nil {
				h.Notify(messaging.NewStatusError(messaging.StatusIOError, status))
			}
			return entry, status
		},
		New: func(h messaging.Notifier, origin common.Origin) ([]HostEntry, error) {
			return nil, errors.New("error: not found")
		},
	}
}()
