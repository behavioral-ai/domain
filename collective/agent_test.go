package collective

import "fmt"

func ExampleNewAgent() {
	a := newContentAgent(nil, storeGet)

	fmt.Printf("test: newContentAgent() -> [%v]\n", a)

	//Output:
	//fail
}
