package frame

import "fmt"

func ExampleNewUri() {
	uri := NewUri("ingress", "operative1", "frame1")

	fmt.Printf("test: NewUri() -> %v\n", uri)

	//Output:
	//test: NewUri() -> ingress:operative1.frame1.1699ede7-d02d-11ef-82a3-00a55441ed8b

}
