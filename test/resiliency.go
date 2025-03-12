package test

import (
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/testrsc"
)

const (
	ResiliencyThreshold = "resiliency:type/operative/agent/threshold"
	ResiliencyInterpret = "resiliency:type/operative/agent/interpret"
)

func NewResiliencyResolver() (collective.Resolution, *messaging.Status) {
	r := collective.NewEphemeralResolver()
	status := loadResiliencyContent(r)
	if status.OK() {
		status = loadProfile(r)
	}
	return r, status
}

func loadResiliencyContent(r collective.Resolution) *messaging.Status {
	buf, err := iox.ReadFile(testrsc.ResiliencyInterpret1)
	if err != nil {
		return messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	status := r.PutContent(ResiliencyInterpret, "author", buf, 1)
	if !status.OK() {
		return status
	}
	buf, err = iox.ReadFile(testrsc.ResiliencyThreshold1)
	if err != nil {
		return messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	return r.PutContent(ResiliencyThreshold, "author", buf, 1)
}
