package collective

import "fmt"

func ExampleThingAppend() {
	ok := thingAppend("urn:agent:thing/three", "", "cn2", "")
	fmt.Printf("test: ThingAppend() -> [ok:%v] [%v]\n", ok, things)

	ok = thingAppend("urn:agent:thing/three", "", "", "")
	fmt.Printf("test: ThingAppend() -> [ok:%v] [%v]\n", ok, things)

	//Output:
	//fail
}

func ExampleThingAppendWithUri() {
	ok := thingAppend("urn:agent:thing/four", "", "cn2", "http://google.com/search")
	fmt.Printf("test: ThingAppend() -> [ok:%v] [%v]\n", ok, things)
	fmt.Printf("test: ResolutionAppend() -> [ok:%v] [%v]\n", ok, resolutions)

	//Output:
	//fail
}
