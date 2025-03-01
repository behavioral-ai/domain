package timeseries1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
)

const (
	PkgPath = "github/behavioral-ai/operative/timeseries1"
)

// Entry - timeseries data
type Entry struct {
	Origin   common.Origin `json:"origin"`
	Latency  int           `json:"latency"` // Milliseconds for the 95th percentile
	Gradient int           `json:"gradient"`
}

// Observation - observation interface
type Observation struct {
	Timeseries func(origin common.Origin) (Entry, *messaging.Status)
}

var Observe = func() *Observation {
	return &Observation{
		Timeseries: func(origin common.Origin) (Entry, *messaging.Status) {
			return get(origin)
		},
	}
}()

func NewObservation(e Entry, status *messaging.Status) *Observation {
	return &Observation{
		Timeseries: func(origin common.Origin) (Entry, *messaging.Status) {
			return e, status
		},
	}
}
