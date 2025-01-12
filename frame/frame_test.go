package frame

import "fmt"

func ExampleNewUri() {
	uri := NewUri(FileScheme, "ingress", "operative1", "frame1")
	fmt.Printf("test: NewUri() -> %v\n", uri)

	uri = NewUri(UrnScheme, "ingress", "operative1", "frame1")
	fmt.Printf("test: NewUri() -> %v\n", uri)

	//Output:
	//test: NewUri() -> file:ingress:operative1.frame1.b55376c4-d038-11ef-a256-00a55441ed8b
	//test: NewUri() -> urn:ingress:operative1.frame1.b5550614-d038-11ef-a256-00a55441ed8b

}

func ExampleNewUriTemplate() {
	uri := NewUriTemplate(FileScheme, "ingress", "operative1", "frame1")
	fmt.Printf("test: NewUriTemplate() -> %v\n", uri)

	uri = NewUriTemplate(UrnScheme, "ingress", "operative1", "frame1")
	fmt.Printf("test: NewUriTemplate() -> %v\n", uri)

	uri = NewUriTemplate(UrnScheme, "", "operative1", "frame1")
	fmt.Printf("test: NewUriTemplate() -> %v\n", uri)

	uri = NewUriTemplate(UrnScheme, "ingress", "", "frame1")
	fmt.Printf("test: NewUriTemplate() -> %v\n", uri)

	uri = NewUriTemplate(UrnScheme, "ingress", "operative1", "")
	fmt.Printf("test: NewUriTemplate() -> %v\n", uri)

	//Output:
	//test: NewUriTemplate() -> file:ingress:operative1.frame1
	//test: NewUriTemplate() -> urn:ingress:operative1.frame1
	//test: NewUriTemplate() -> urn:*:operative1.frame1
	//test: NewUriTemplate() -> urn:ingress:*.frame1
	//test: NewUriTemplate() -> urn:ingress:operative1.*

}
