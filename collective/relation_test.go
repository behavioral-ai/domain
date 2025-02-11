package collective

import "fmt"

func ExampleRelationAppend() {
	ok := relationAppend("urn:agent:aspect/test3", "urn:agent:aspect/test4", "")
	fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)

	ok = relationAppend("urn:agent:aspect/test1", "urn:agent:aspect/test3", "")
	fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)

	ok = relationAppend("urn:agent:aspect/test1", "urn:agent:aspect/test3", "")
	fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)

	//Output:
	//fail
}
