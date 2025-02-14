package collective

import (
	"encoding/json"
	"fmt"
	"github.com/behavioral-ai/core/jsonx"
)

type Lookup struct {
	Low    int `json:"low"`
	Medium int `json:"medium"`
	High   int `json:"high"`
}

func ExampleResolveFailure() {
	_, err := fileResolver("invalid urn", 1)
	if err != nil {
		fmt.Printf("test: Resolver() -> [err:%v]\n", err)
	}

	//Output:
	//test: Resolver() -> [err:open error: scheme is invalid []: The system cannot find the file specified.]

}

func ExampleResolveSuccess() {
	buf, err := fileResolver(ResiliencyGradient, 1)
	if err != nil {
		fmt.Printf("test: Resolver() -> [err:%v]\n", err)
	} else {
		l, err1 := jsonx.New[[]Lookup](buf, nil)
		fmt.Printf("test: Resolver() -> [err:%v] [%v]\n", err1, l)
	}

	//Output:
	//test: Resolver() -> [err:<nil>] [[{10 40 80}]]

}

func _ExampleLookup() {
	l := Lookup{Low: 10, Medium: 40, High: 80}
	buf, err := json.Marshal(l)
	fmt.Printf("test: Lookup() -> [%v] [%v]\n", err, string(buf))

	//Output:
	//fail
}
