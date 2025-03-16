package test

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/content"
	"github.com/behavioral-ai/domain/testrsc"
	url2 "net/url"
)

const (
	ProfileName = "resiliency:type/domain/metrics/profile"
)

func LoadProfile(r content.Resolution) *messaging.Status {
	url, _ := url2.Parse(testrsc.ResiliencyTrafficProfile1)
	return r.AddValue(ProfileName, "author", url, 1)
}
