package test

import (
	"fmt"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/metrics1"
)

func ExampleStartup() {
	r := collective.Resolver
	Startup()

	name := ResiliencyThreshold
	buf, status1 := r.GetContent(name, 1)
	fmt.Printf("test: GetContent(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, buf != nil)

	name = ResiliencyInterpret
	buf, status1 = r.GetContent(name, 1)
	fmt.Printf("test: GetContent(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, buf != nil)

	name = metrics1.ProfileName
	buf, status1 = r.GetContent(name, 1)
	fmt.Printf("test: GetContent(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, buf != nil)

	//Output:
	//test: GetContent("resiliency:type/operative/agent/threshold") -> [status:OK] [buf:true]
	//test: GetContent("resiliency:type/operative/agent/interpret") -> [status:OK] [buf:true]
	//test: GetContent("resiliency:type/domain/metrics/profile") -> [status:OK] [buf:true]

}
