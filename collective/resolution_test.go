package collective

import "fmt"

func ExampleResolutionAppend() {
	ok := resolutionAppend("urn:agent:thing/resource", "", Name)
	fmt.Printf("test: ResolutionAppend() -> [ok:%v] [%v]\n", ok, resolutions)

	/*
		ok = relationAppend("test1", "test3", "")
		fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)

		ok = relationAppend("test1", "test3", "")
		fmt.Printf("test: RelationAppend() -> [ok:%v] [%v]\n", ok, relations)


	*/

	//Output:
	//fail
}
