package test

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/content"
	"github.com/behavioral-ai/domain/testrsc"
	url2 "net/url"
)

func loadResolver(resolver content.Resolution) *messaging.Status {
	url, _ := url2.Parse(testrsc.ResiliencyInterpret1)
	status := resolver.AddValue(ResiliencyInterpret, "author", url, 1)
	if !status.OK() {
		return status
	}
	url, _ = url2.Parse(testrsc.ResiliencyThreshold1)
	return resolver.AddValue(ResiliencyThreshold, "author", url, 1)
}
