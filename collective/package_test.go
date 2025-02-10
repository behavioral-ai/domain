package collective

import (
	"fmt"
	"net/url"
)

func ExampleThreshold() {
	urn := "urn:app:thing/knowledge/retrieval/threshold"

	urn = "https://google.com/search?exist"
	u, err := url.Parse(urn)
	q := u.Query()
	fmt.Printf("test: Threshold() -> [err:%v] [scheme:%v] [path:%v] [values:%v]\n", err, u.Scheme, u.Path, q)

	urn = "https://google.com/search?!exist"
	u, err = url.Parse(urn)
	q = u.Query()
	fmt.Printf("test: Threshold() -> [err:%v] [scheme:%v] [path:%v] [values:%v]\n", err, u.Scheme, u.Path, q)

	urn = "https://google.com/search?like&top=5"
	u, err = url.Parse(urn)
	q = u.Query()
	fmt.Printf("test: Threshold() -> [err:%v] [scheme:%v] [path:%v] [values:%v]\n", err, u.Scheme, u.Path, q)

	urn = "https://google.com/search?!like&top=5"
	u, err = url.Parse(urn)
	q = u.Query()
	fmt.Printf("test: Threshold() -> [err:%v] [scheme:%v] [path:%v] [values:%v]\n", err, u.Scheme, u.Path, q)

	//Output:
	//fail

}
