package timeseries1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
)

const (
	PkgPath = "github/behavioral-ai/operative/timeseries1"
)

// Observation - timeseries data
type Observation struct {
	Origin   common.Origin `json:"origin"`
	Latency  int           `json:"latency"` // Milliseconds for the 95th percentile
	Gradient int           `json:"gradient"`
}

// Observer - observation interface
type Observer struct {
	Timeseries func(origin common.Origin) (Observation, *messaging.Status)
}

var Observations = func() *Observer {
	return &Observer{
		Timeseries: func(origin common.Origin) (Observation, *messaging.Status) {
			return getObservation(origin)
		},
	}
}()

func NewObservation(e Observation, status *messaging.Status) *Observer {
	return &Observer{
		Timeseries: func(origin common.Origin) (Observation, *messaging.Status) {
			return e, status
		},
	}
}

// Assignment - host
type Assignment struct {
	Origin common.Origin `json:"origin"`
}

type SelectAssignments func(origin common.Origin) ([]Assignment, *messaging.Status)

// Assigner - assignments functions struct
type Assigner struct {
	All func(origin common.Origin) ([]Assignment, *messaging.Status)
	New func(origin common.Origin) ([]Assignment, *messaging.Status)
}

var Assignments = func() *Assigner {
	return &Assigner{
		All: func(origin common.Origin) ([]Assignment, *messaging.Status) {
			return getAssignment(origin)
		},
		New: func(origin common.Origin) ([]Assignment, *messaging.Status) {
			return nil, messaging.StatusNotFound()
		},
	}
}()
