package collective

import (
	"encoding/json"
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/domain/testrsc"
)

func ExampleFileGet() {

	fmt.Printf("test: File() -> [%v]\n", "test")

	//Output:
	//test: File() -> [test]

}

type lookup struct {
	Low    int `json:"low"`
	Medium int `json:"medium"`
	High   int `json:"high"`
}

type lookupKey struct {
	Key    ResolutionKey
	Low    int `json:"low"`
	Medium int `json:"medium"`
	High   int `json:"high"`
}

func ExampleLookup() {
	var l []lookup

	buf, err := iox.ReadFile(testrsc.ResiliencyGradient)
	fmt.Printf("test: iox.ReadFile() -> [err:%v]\n", err)

	err = json.Unmarshal(buf, &l)
	fmt.Printf("test: json.Unmarshal() -> [err:%v] [%v]\n", err, l)

	buf, err = iox.ReadFile(testrsc.ResiliencyGradientKey)
	fmt.Printf("test: iox.ReadFile() -> [err:%v]\n", err)

	err = json.Unmarshal(buf, &l)
	fmt.Printf("test: json.Unmarshal() -> [err:%v] [%v]\n", err, l)

	//Output:
	//test: iox.ReadFile() -> [err:<nil>]
	//test: json.Unmarshal() -> [err:<nil>] [[{10 40 80}]]
	//test: iox.ReadFile() -> [err:<nil>]
	//test: json.Unmarshal() -> [err:<nil>] [[{10 40 80}]]

}

func ExampleLookupKey() {
	var l []lookupKey

	buf, err := iox.ReadFile(testrsc.ResiliencyGradient)
	fmt.Printf("test: iox.ReadFile() -> [err:%v]\n", err)

	err = json.Unmarshal(buf, &l)
	fmt.Printf("test: json.Unmarshal() -> [err:%v] [%v]\n", err, l)

	buf, err = iox.ReadFile(testrsc.ResiliencyGradientKey)
	fmt.Printf("test: iox.ReadFile() -> [err:%v]\n", err)

	err = json.Unmarshal(buf, &l)
	fmt.Printf("test: json.Unmarshal() -> [err:%v] [%v]\n", err, l)

	//Output:
	//test: iox.ReadFile() -> [err:<nil>]
	//test: json.Unmarshal() -> [err:<nil>] [[{{ 0} 10 40 80}]]
	//test: iox.ReadFile() -> [err:<nil>]
	//test: json.Unmarshal() -> [err:<nil>] [[{{test:name 2} 10 40 80}]]

}
