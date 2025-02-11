package collective

import (
	"fmt"
	"net/url"
)

func _ExampleThreshold() {
	urn := "urn:app:thing/knowledge/retrieval/threshold"
	//req,_ := http.

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

func ExampleResolutionUrn() {
	in := "urn:thing:/test/resource/name"
	urn := ResolutionUrn(Urn(in))

	fmt.Printf("test: ResoulationNSS() -> [%v] [%v]\n", in, urn)

	//Output:
	//fail
}
