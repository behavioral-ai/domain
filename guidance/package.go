package guidance

import (
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/domain/common"
)

const (
	PkgPath    = "github/behavioral-ai/guidance/guidance"
	WestRegion = "us-west1"
	WestZoneA  = "w-a"
	WestZoneB  = "w-b"

	CentralRegion = "us-central1"
	CentralZoneA  = "c-a"
	CentralZoneB  = "c-b"

	EastRegion = "us-east1"
	EastZoneA  = "e-a"
	EastZoneB  = "e-b"
)

func GetRegion(origin common.Origin) ([]HostEntry, *aspect.Status) {
	if origin.Region == WestRegion {
		return westData, aspect.StatusOK()
	}
	if origin.Region == CentralRegion {
		return centralData, aspect.StatusOK()
	}
	return []HostEntry{}, aspect.StatusNotFound()
}
