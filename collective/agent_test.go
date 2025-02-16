package collective

import "fmt"

func ExampleNewAgent() {
	a := newHttpAgent(nil)

	fmt.Printf("test: newHttpAgent() -> [%v]\n", a)

	//Output:
	//test: newHttpAgent() -> [content.agent]
	
}
