package test

import (
	"fmt"
	"github.com/behavioral-ai/domain/metrics1"
)

func ExampleNewResiliencyResolver() {
	r, status := NewResiliencyResolver()

	fmt.Printf("test: NewResiliencyResolver() -> [status:%v]\n", status)

	name := ResiliencyThreshold
	buf, status1 := r.GetContent(name, 1)
	fmt.Printf("test: GetContent(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, len(buf))

	name = ResiliencyInterpret
	buf, status1 = r.GetContent(name, 1)
	fmt.Printf("test: GetContent(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, len(buf))

	name = metrics1.ProfileName
	buf, status1 = r.GetContent(name, 1)
	fmt.Printf("test: GetContent(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, len(buf))

	//Output:
	//test: NewResiliencyResolver() -> [status:OK]
	//test: GetContent("resiliency:type/operative/agent/threshold") -> [status:OK] [buf:303]
	//test: GetContent("resiliency:type/operative/agent/interpret") -> [status:OK] [buf:316]
	//test: GetContent("resiliency:type/domain/metrics/profile") -> [status:OK] [buf:1218]

}
