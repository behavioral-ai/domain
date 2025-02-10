package timeseries1

import (
	"context"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"time"
)

const (
	timeseriesDuration = time.Second * 2
)

// Observation - observation functions struct
type Observation struct {
	Timeseries func(h messaging.Notifier, origin common.Origin) (Entry, *aspect.Status)
}

var Observe = func() *Observation {
	return &Observation{
		Timeseries: func(h messaging.Notifier, origin common.Origin) (Entry, *aspect.Status) {
			ctx, cancel := context.WithTimeout(context.Background(), timeseriesDuration)
			defer cancel()
			if ctx != nil {
			}
			//e, status := timeseries.Query(ctx, origin)
			//if !status.OK() && !status.NotFound() {
			//	h.Notify(status)
			//}
			return Entry{Gradient: 100, Latency: 55}, aspect.StatusOK()
		},
	}
}()
