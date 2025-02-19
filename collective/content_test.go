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
		status := c.put(name, buf, 1)
		fmt.Printf("test: newContentCache.put(1) -> [status:%v]\n", status)

		buf, status = c.get(name, 2)
		fmt.Printf("test: newContentCache.get(2) -> [status:%v]\n", status)

		buf, status = c.get(name, 1)
		fmt.Printf("test: newContentCache.get(1) -> [status:%v]\n", status)

		var v text
		err = json.Unmarshal(buf, &v)
		fmt.Printf("test: json.Unmarshal() -> [err:%v] [%v]\n", err, v)
	}

	//Output:
	//test: newContentCache.put(1) -> [status:OK]
	//test: newContentCache.get(2) -> [status:Not Found [error: name "test:thing:text" version "2"]]
	//test: newContentCache.get(1) -> [status:OK]
	//test: json.Unmarshal() -> [err:<nil>] [{Hello World!}]

}
