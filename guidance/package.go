package guidance

import (
	"github.com/behavioral-ai/domain/common"
)

const (
	PkgPath    = "github/behavioral-ai/domain/guidance"
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

func GetRegion(origin common.Origin) ([]HostEntry, error) {
	if origin.Region == WestRegion {
		return westData, nil
	}
	if origin.Region == CentralRegion {
		return centralData, nil
	}
	return []HostEntry{}, nil
}
