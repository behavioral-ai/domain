package collective

import "fmt"

func ExampleRelationAppend() {
	ok := relationAppend("test1", "test2", "")
	fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)

	ok = relationAppend("test1", "test3", "")
	fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)

	ok = relationAppend("test1", "test3", "")
	fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)

	//Output:
	//fail
}
