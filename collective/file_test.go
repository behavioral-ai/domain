package collective

import (
	"encoding/json"
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/domain/testrsc"
	"io/fs"
	"log"
	"os"
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

func _ExampleParseResolutionKey() {
	buf, _ := iox.ReadFile(testrsc.ResiliencyGradientKey)
	s := string(buf)

	k, err := parseResolutionKey(s)
	fmt.Printf("test: ParseResolutionKey() -> [err:%v] [%v]\n", err, k)

	//Output:
	//test: ParseResolutionKey() -> [err:<nil>] [{test:name 2}]

}

func ExampleReadDir() {
	//dir, err := os.Getwd()
	var err error
	dir := "c:\\Users\\markb\\GitHub\\domain\\testrsc\\files\\resiliency"
	if err != nil {
		fmt.Printf("test: os.Getwd() -> [err:%v]\n", err)
	}
	fmt.Printf("test: wd() -> [%v]\n", dir)
	fileSystem := os.DirFS(dir)
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})

	//Output:
	//fail

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
