package timeseries1

import (
	"context"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	timeseriesDuration = time.Second * 2
)

// Observation - observation functions struct
type Observation struct {
	Timeseries func(h messaging.Notifier, origin core.Origin) (Entry, *core.Status)
}

var Observe = func() *Observation {
	return &Observation{
		Timeseries: func(h messaging.Notifier, origin core.Origin) (Entry, *core.Status) {
			ctx, cancel := context.WithTimeout(context.Background(), timeseriesDuration)
			defer cancel()
			if ctx != nil {
			}
			//e, status := timeseries.Query(ctx, origin)
			//if !status.OK() && !status.NotFound() {
			//	h.Notify(status)
			//}
			return Entry{Gradient: 100, Latency: 55}, core.StatusOK()
		},
	}
}()
