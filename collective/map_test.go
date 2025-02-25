package collective

import (
	"fmt"
)

func ExampleNewMapCache() {
	name := "test:map"
	key1 := "key1"
	m := make(map[string]string)
	m[key1] = "value for key1"

	c := newMapCache()
	err := c.put(name, m)
	fmt.Printf("test: newMapCache.put() -> [err:%v]\n", err)

	m2, err2 := c.get(name)
	fmt.Printf("test: newMapCache.get() -> [err:%v] [%v]\n", err2, m2)

	//Output:
	//test: newMapCache.put() -> [err:<nil>]
	//test: newMapCache.get() -> [err:<nil>] [map[key1:value for key1]]

}
