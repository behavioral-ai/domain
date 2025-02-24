package collective

import "fmt"

func ExampleNewAgent() {
	a := newContentAgent(false, nil)

	fmt.Printf("test: newHttpAgent() -> [%v]\n", a)

	//Output:
	//test: newHttpAgent() -> [resiliency:agent/domain/collective/content]

}
