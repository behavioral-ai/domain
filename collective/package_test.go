package collective

import (
	"encoding/json"
	"fmt"
	"github.com/behavioral-ai/core/test"
)

func ExampleResolveString() {
	name := "test:thing:text@2"

	c := newContentCache()
	buf, err := json.Marshal(text{Value: "Generic typed get"})
	if err != nil {
		fmt.Printf("test: json.Marshall() -> [err:%v]\n", err)
	} else {
		status := c.put(name, buf, 1)
		fmt.Printf("test: newContentCache.put(1) -> [status:%v]\n", status)

		v, status1 := Resolve[text](name, 1, nil)
		fmt.Printf("test: Resolve[text]() -> [status:%v] [%v]\n", status1, v)
	}

	//Output:
	//test: newContentCache.put(1) -> [status:OK]
	//test: Resolve[text]() -> [status:Bad Request [err:error: BadRequest - resolver is nil] [src:]] [{}]

}

func ExampleEphemeralResolver() {
	name := "test:thing/string"
	s := "test Ephemeral resolver"

	r := NewEphemeralResolver("", test.Notify)
	//fmt.Printf("test: NewEphemeralResolver() -> [status:%v]\n", status)

	status := r.PutContent(name, "", s, 1)
	fmt.Printf("test: Resolver.Put() -> [status:%v]\n", status)

	v, status1 := Resolve[string](name, 1, r)
	fmt.Printf("test: Resolve[string] -> [status:%v] [%v]\n", status1, v)

	v, status1 = Resolve[string](name, 2, r)
	fmt.Printf("test: Resolve[string] -> [status:%v] [%v]\n", status1, v)

	//Output:
	//test: Resolver.Put() -> [status:OK]
	//test: Resolve[string] -> [status:OK] [test Ephemeral resolver]
	//notify-> [event:messaging:status] [msg:Not Found] [src:] [agent:]
	//test: Resolve[string] -> [status:Not Found] []

}
