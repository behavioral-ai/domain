package content

import "fmt"

func ExampleNewAgent() {
	a := newContentAgent(false, nil)

	fmt.Printf("test: newHttpAgent() -> [%v]\n", a)

	//Output:
	//test: newHttpAgent() -> [resiliency:agent/domain/content/content]

}
