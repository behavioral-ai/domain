package timeseries1

import (
	"fmt"
	"github.com/behavioral-ai/domain/common"
)

func ExampleObserver() {
	o := common.Origin{Region: common.WestRegion, Zone: common.WestZoneA}
	e, status := Observations.Timeseries(o)

	fmt.Printf("test: Observer() -> [status:%v] [entry:%v]\n", status, e)

	//Output:
	//fail
}
