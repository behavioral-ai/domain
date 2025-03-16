package content

import "fmt"

func Example_Get() {

	_, status := get("name", 1)

	fmt.Printf("test: get() -> [status:%v]\n", status)

	//Output:
	//test: get() -> [status:Not Found]

}

func Example_Put() {

	_, status := put("name", "author", nil, 1)
	fmt.Printf("test: put() -> [status:%v]\n", status)

	//Output:
	//test: put() -> [status:OK]

}
