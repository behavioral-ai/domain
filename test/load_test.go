package test

import (
	"fmt"
	"github.com/behavioral-ai/domain/content"
)

func ExampleLoadResolver() {
	r := content.NewEphemeralResolver()
	status := loadResolver(r)
	fmt.Printf("test: loadResolver() -> [status:%v]\n", status)

	buf, status1 := r.GetValue(ResiliencyThreshold, 1)
	fmt.Printf("test: GetValue(\"%v\") -> [status:%v] [buf:%v]\n", ResiliencyThreshold, status1, string(buf))

	buf, status1 = r.GetValue(ResiliencyInterpret, 1)
	fmt.Printf("test: GetValue(\"%v\") -> [status:%v] [buf:%v]\n", ResiliencyInterpret, status1, string(buf))

	//Output:
	//test: loadResolver() -> [status:OK]
	//test: GetValue("resiliency:type/operative/agent/threshold") -> [status:OK] [buf:true]
	//test: GetValue("resiliency:type/operative/agent/interpret") -> [status:OK] [buf:true]

}
