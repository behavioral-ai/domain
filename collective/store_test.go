package collective

import (
	"fmt"
	"github.com/behavioral-ai/domain/common"
)

type host struct {
	EntryId int           `json:"entry-id"`
	Created Timestamp     `json:"created-ts"`
	Origin  common.Origin `json:"origin"`
}

func _ExampleStore() {

	//Output:
	//fail
}

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
