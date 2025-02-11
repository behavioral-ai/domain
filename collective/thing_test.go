package collective

import "fmt"

func ExampleThingAppend() {
	ok := thingAppend("test", "", "", "")
	fmt.Printf("test: ThingAppend() -> [ok:%v] [%v]\n", ok, things)

	ok = thingAppend("test", "", "", "")
	fmt.Printf("test: ThingAppend() -> [ok:%v] [%v]\n", ok, things)

	//Output:
	//fail
}
