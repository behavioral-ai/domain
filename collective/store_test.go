package collective

import "fmt"

func _ExampleStore() {
	ok := resolutionAppend("agent:thing/assignment/host", ResourceUri)
	fmt.Printf("test: ResolutionAppend() -> [ok:%v] [%v]\n", ok, resolutions)

	//Output:
	//fail
}

func ExamplePartition() {
	u := "agent:thing/assignment/host"
	p := partition(Urn(u))
	fmt.Printf("test: partition() -> [%v] [%v]\n", u, p)

	u = "agent:thing/assignment/host#partition-1"
	p = partition(Urn(u))
	fmt.Printf("test: partition() -> [%v] [%v]\n", u, p)

	//Output:
	//fail

}
