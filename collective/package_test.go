package collective

import (
	"encoding/json"
	"fmt"
	"github.com/behavioral-ai/core/test"
)

func _ExampleResolveString() {
	name := "test:thing:text@2"

	c := newContentCache()
	buf, err := json.Marshal(text{Value: "Generic typed get"})
	if err != nil {
		fmt.Printf("test: json.Marshall() -> [err:%v]\n", err)
	} else {
		err = c.put(name, buf, 1)
		fmt.Printf("test: newContentCache.put(1) -> [err:%v]\n", err)

		v, err1 := Resolve[text](name, 1)
		fmt.Printf("test: Resolve[text]() -> [err:%v] [%v]\n", err1, v)
	}

	//Output:
	//fail
}

func ExampleInitializeEphemeral() {
	a := "test:agent/operative"
	name := "test:thing/string"
	s := "test Ephemeral resolver"

	err := InitializeEphemeral(test.NewAgent(a), "")
	fmt.Printf("test: InitializeEphemeral() -> [err:%v]\n", err)

	err = Resolver.Put(name, "", s, 1)
	fmt.Printf("test: Resolver.Put() -> [err:%v]\n", err)

	v, err1 := Resolve[string](name, 1)
	fmt.Printf("test: Resolve[string] -> [err:%v] [%v]\n", err1, v)

	//Output:
	//test: InitializeEphemeral() -> [err:<nil>]
	//test: Resolver.Put() -> [err:<nil>]
	//test: Resolve[string] -> [err:<nil>] [test Ephemeral resolver]

}
