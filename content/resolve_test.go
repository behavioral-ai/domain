package content

import (
	"fmt"
	"github.com/behavioral-ai/domain/testrsc"
	url2 "net/url"
)

func ExampleResolution_PutValue() {
	urn := "type/test"
	r := NewEphemeralResolver()

	status := r.AddValue("", "author", nil, 1)
	fmt.Printf("test: PutValue_Name() -> [status:%v]\n", status)

	status = r.AddValue(urn, "author", nil, 1)
	fmt.Printf("test: PutValue_Name() -> [status:%v]\n", status)

	var buff []byte
	status = r.AddValue(urn, "author", buff, 1)
	fmt.Printf("test: PutValue_Name() -> [status:%v]\n", status)

	url, _ := url2.Parse(testrsc.ResiliencyTrafficProfile1)
	status = r.AddValue(urn, "author", url, 1)
	fmt.Printf("test: PutValue() -> [status:%v]\n", status)

	buf, status1 := r.GetValue(urn, 1)
	fmt.Printf("test: GetValue() -> [status:%v] [%v]\n", status1, len(buf) > 0)

	//Output:
	//test: PutValue_Name() -> [status:Bad Request [err:nsName is empty on call to PutValue()] [agent:resiliency:agent/domain/content/content]]
	//test: PutValue_Name() -> [status:No Content [err:content is nil on call to PutValue() for nsName : type/test] [agent:resiliency:agent/domain/content/content]]
	//test: PutValue_Name() -> [status:No Content [err:content is empty on call to PutValue() for nsName : type/test] [agent:resiliency:agent/domain/content/content]]
	//test: PutValue() -> [status:OK]
	//test: GetValue() -> [status:OK] [true]

}
