package timeseries1

import (
	"fmt"
	"github.com/behavioral-ai/domain/common"
)

func ExampleGet() {
	o := common.Origin{Region: common.EastRegion, Zone: common.WestZoneA}
	e, status := get(o)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [%v]\n", o, status, e)

	o = common.Origin{Region: common.WestRegion, Zone: common.WestZoneA}
	e, status = get(o)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [%v]\n", o, status, e)

	o = common.Origin{Region: common.WestRegion, Zone: common.WestZoneB}
	e, status = get(o)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [%v]\n", o, status, e)

	o = common.Origin{Region: common.CentralRegion, Zone: common.CentralZoneA}
	e, status = get(o)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [%v]\n", o, status, e)

	o = common.Origin{Region: common.CentralRegion, Zone: common.CentralZoneB}
	e, status = get(o)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [%v]\n", o, status, e)

	//Output:
	//test: Get("us-east1.w-zone-a") -> [status:Not Found] [{ 0 0}]
	//test: Get("us-west1.w-zone-a") -> [status:OK] [{us-west1.w-zone-a.host1.com 1500 2}]
	//test: Get("us-west1.w-zone-b") -> [status:OK] [{us-west1.w-zone-b.host2.com 900 33}]
	//test: Get("us-central1.c-zone-a") -> [status:OK] [{us-central1.c-zone-a.host3.com 2000 55}]
	//test: Get("us-central1.c-zone-b") -> [status:OK] [{us-central1.c-zone-b.host4.com 850 25}]

}

func ExampleGetReset() {
	o := common.Origin{Region: common.WestRegion, Zone: common.WestZoneA}
	for i := 0; i < 5; i++ {
		e, status := get(o)
		fmt.Printf("test: Get(\"%v\") -> [status:%v] [%v]\n", o, status, e)

	}

	//Output:
	//test: Get("us-west1.w-zone-a") -> [status:OK] [{us-west1.w-zone-a.host1.com 1200 25}]
	//test: Get("us-west1.w-zone-a") -> [status:OK] [{us-west1.w-zone-a.host1.com 1500 2}]
	//test: Get("us-west1.w-zone-a") -> [status:OK] [{us-west1.w-zone-a.host1.com 1200 25}]
	//test: Get("us-west1.w-zone-a") -> [status:OK] [{us-west1.w-zone-a.host1.com 1500 2}]
	//test: Get("us-west1.w-zone-a") -> [status:OK] [{us-west1.w-zone-a.host1.com 1200 25}]

}
