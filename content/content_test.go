package content

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
		var status error //status *messaging.Status
		c.put(name, buf, 1)
		//fmt.Printf("test: newContentCache.put(1) -> [status:%v]\n", status)

		buf, status = c.get(name, 2)
		fmt.Printf("test: newContentCache.get(2) -> [status:%v]\n", status)

		buf, status = c.get(name, 1)
		fmt.Printf("test: newContentCache.get(1) -> [status:%v]\n", status)

		var v text
		err = json.Unmarshal(buf, &v)
		fmt.Printf("test: json.Unmarshal() -> [err:%v] [%v]\n", err, v)
	}

	//Output:
	//test: newContentCache.get(2) -> [status:content [test:thing:text] [2] not found]
	//test: newContentCache.get(1) -> [status:<nil>]
	//test: json.Unmarshal() -> [err:<nil>] [{Hello World!}]

}
