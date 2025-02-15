package collective

import (
	"fmt"
)

func ExampleRelationAppend() {
	ok := relationAppend("agent:aspect/test3", "agent:aspect/test4")
	fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)

	ok = relationAppend("agent:aspect/test1", "agent:aspect/test3")
	fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)

	ok = relationAppend("agent:aspect/test1", "agent:aspect/test3")
	fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)

	//Output:
	//fail
}
