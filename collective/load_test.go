package collective

import "fmt"

func ExampleLoadResolver() {
	r := NewEphemeralResolver()
	status := loadResolver(r)
	fmt.Printf("test: loadResolver() -> [status:%v]\n", status)

	buf, status1 := r.GetContent(ResiliencyThreshold, 1)
	fmt.Printf("test: GetContent(\"%v\") -> [status:%v] [buf:%v]\n", ResiliencyThreshold, status1, buf != nil)

	buf, status1 = r.GetContent(ResiliencyInterpret, 1)
	fmt.Printf("test: GetContent(\"%v\") -> [status:%v] [buf:%v]\n", ResiliencyInterpret, status1, buf != nil)

	//Output:
	//test: loadResolver() -> [status:OK]
	//test: GetContent("resiliency:type/operative/agent/threshold") -> [status:OK] [buf:true]
	//test: GetContent("resiliency:type/operative/agent/interpret") -> [status:OK] [buf:true]

}
