package collective

import "fmt"

func ExampleNewAgent() {
	a := newContentAgent(false)

	fmt.Printf("test: newHttpAgent() -> [%v]\n", a)

	//Output:
	//test: newHttpAgent() -> [content.agent]

}
