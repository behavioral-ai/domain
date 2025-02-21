package collective

import "fmt"

func ExampleNewAgent() {
	a := newContentAgent(false, nil, nil)

	fmt.Printf("test: newHttpAgent() -> [%v]\n", a)

	//Output:
	//test: newHttpAgent() -> [root:agent/domain/collective/content]

}
