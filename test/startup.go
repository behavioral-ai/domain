package test

import (
	"fmt"
	"github.com/behavioral-ai/domain/collective"
)

const (
	ResiliencyThreshold = "resiliency:type/operative/agent/threshold"
	ResiliencyInterpret = "resiliency:type/operative/agent/interpret"
)

// Testing
//r.activity = func(hostName string, agent messaging.Agent, event, source string, content any) {
//	fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", messaging.FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), event, source, content)
//}

func Startup() {
	status := loadResolver(collective.Resolver)
	if !status.OK() {
		fmt.Printf("error on loading Resolver: %v\n", status)
	}
	status = LoadProfile(collective.Resolver)
	if !status.OK() {
		fmt.Printf("error on loading Resolver: %v\n", status)
	}
}
