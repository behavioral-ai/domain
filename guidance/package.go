package guidance

import "github.com/behavioral-ai/core/core"

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

func GetRegion(origin core.Origin) ([]HostEntry, *core.Status) {
	if origin.Region == WestRegion {
		return westData, core.StatusOK()
	}
	if origin.Region == CentralRegion {
		return centralData, core.StatusOK()
	}
	return []HostEntry{}, core.StatusNotFound()
}
