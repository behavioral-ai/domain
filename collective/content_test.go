package collective

import (
	"encoding/json"
	"fmt"
)

func ExampleNewContentCache() {
	name := "test:thing:text"

	c := newContentCache()
	buf, err := json.Marshal(text{Value: "Hello World!"})
	if err != nil {
		fmt.Printf("test: json.Marshall() -> [err:%v]\n", err)
	} else {
		err = c.put(name, buf, 1)
		fmt.Printf("test: newContentCache.put(1) -> [err:%v]\n", err)

		buf, err = c.get(name, 2)
		fmt.Printf("test: newContentCache.get(2) -> [err:%v]\n", err)

		buf, err = c.get(name, 1)
		fmt.Printf("test: newContentCache.get(1) -> [err:%v]\n", err)

		var v text
		err = json.Unmarshal(buf, &v)
		fmt.Printf("test: json.Unmarshal() -> [err:%v] [%v]\n", err, v)
	}

	//Output:
	//test: newContentCache.put(1) -> [err:<nil>]
	//test: newContentCache.get(2) -> [err:error: NotFound - name "test:thing:text" version "2"]
	//test: newContentCache.get(1) -> [err:<nil>]
	//test: json.Unmarshal() -> [err:<nil>] [{Hello World!}]

}
