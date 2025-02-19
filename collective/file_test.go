package collective

import (
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/testrsc"
	"strings"
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
	Key    ResolutionKey `json:"resolution-key"`
	Low    int           `json:"low"`
	Medium int           `json:"medium"`
	High   int           `json:"high"`
}

/*
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
	//test: json.Unmarshal() -> [err:<nil>] [[{{resiliency:thing/operative/agent/gradient 1} 10 40 80}]]

}


*/

func ExampleParseResolutionKey() {
	buf, _ := iox.ReadFile(testrsc.ResiliencyGradient)
	s := string(buf)

	k, err := parseResolutionKey(s)
	fmt.Printf("test: ParseResolutionKey() -> [err:%v] [%v]\n", err, k)

	//Output:
	//test: ParseResolutionKey() -> [err:<nil>] [{resiliency:thing/operative/agent/gradient 1}]

}

func _ExampleFileLoad() {
	dir := "file:///c:/Users/markb/GitHub/domain/testrsc/files/resiliency"
	name1 := "resiliency:thing/operative/agent/gradient"
	name2 := "resiliency:thing/operative/agent/threshold"

	c := newContentCache()
	err := loadContent(notifyFunc, c, dir)
	fmt.Printf("test: loadContent() -> [err:%v]\n", err)

	buf, status := c.get(name1, 1)
	fmt.Printf("test: c.get() -> [status:%v] [%v]\n", status, string(buf))

	buf, status = c.get(name2, 2)
	fmt.Printf("test: c.get() -> [status:%v] [%v]\n", status, string(buf))

	//Output:
	//fail

}

func notifyFunc(status *messaging.Status) {
	fmt.Printf("status: %v", status)
}

func ExampleEphemeralLoad() {
	name1 := "resiliency:thing/operative/agent/gradient"
	name2 := "resiliency:thing/operative/agent/threshold"
	dir := "file:///c:/Users/markb/GitHub/domain/testrsc/files/resiliency"

	r, status := NewEphemeralResolver(dir, notifyFunc)
	fmt.Printf("test: NewEphemeralResolver() -> [status:%v]\n", status)

	v, status1 := Resolve[[]lookup](name1, 1, r)
	fmt.Printf("test: Resolve[[]lookup] -> [status:%v] [%v]\n", status1, v)

	v, status1 = Resolve[[]lookup](name2, 2, r)
	fmt.Printf("test: Resolve[[]lookup] -> [status:%v] [%v]\n", status1, v)

	//Output:
	//test: NewEphemeralResolver() -> [status:OK]
	//test: Resolve[[]lookup] -> [status:OK] [[{10 40 80}]]
	//test: Resolve[[]lookup] -> [status:OK] [[{15 42 85}]]

}

func extractString(s string) string {
	tokens := strings.Split(s, "\"")
	//fmt.Printf("extractString -> %v\n", tokens[1])
	return tokens[1]
}

func extractValue(s string) string {
	tokens := strings.Split(s, "\r")
	//fmt.Printf("extractValue -> %v\n", tokens[0])
	return tokens[0]
}
