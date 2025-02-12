package collective

import "fmt"

func ExampleThingAppend() {
	ok := thingAppend("agent:thing/three", "cn2")
	fmt.Printf("test: ThingAppend() -> [ok:%v] [%v]\n", ok, things)

	ok = thingAppend("agent:thing/three", "")
	fmt.Printf("test: ThingAppend() -> [ok:%v] [%v]\n", ok, things)

	//Output:
	//fail
}
