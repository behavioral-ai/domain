package test

import (
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/metrics1"
	"github.com/behavioral-ai/domain/testrsc"
)

func loadProfile(r collective.Resolution) *messaging.Status {
	buf, err := iox.ReadFile(testrsc.ResiliencyTrafficProfile1)
	if err != nil {
		return messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	return r.PutContent(metrics1.ProfileName, "author", buf, 1)
}
