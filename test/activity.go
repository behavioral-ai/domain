package test

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

func AddActivity(agent messaging.Agent, event, source string, content any) {
	fmt.Printf("activity -> [%v] [event:%v] [src:%v]\n", agent.Uri(), event, source)
}
