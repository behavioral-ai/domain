package timeseries1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
)

var (
	westZoneA = common.Origin{Region: common.WestRegion, Zone: common.WestZoneA, Host: "host1.com"}
	westZoneB = common.Origin{Region: common.WestRegion, Zone: common.WestZoneB, Host: "host2.com"}

	centralZoneA = common.Origin{Region: common.CentralRegion, Zone: common.CentralZoneA, Host: "host3.com"}
	centralZoneB = common.Origin{Region: common.CentralRegion, Zone: common.CentralZoneB, Host: "host4.com"}

	westZoneAIndex  = 0
	westZoneASeries = []Entry{
		{Origin: westZoneA, Latency: 1500, Gradient: 2},
		{Origin: westZoneA, Latency: 1200, Gradient: 25},
	}

	westZoneBIndex  = 0
	westZoneBSeries = []Entry{
		{Origin: westZoneB, Latency: 900, Gradient: 33},
		{Origin: westZoneB, Latency: 500, Gradient: 2},
	}

	centralZoneAIndex  = 0
	centralZoneASeries = []Entry{
		{Origin: centralZoneA, Latency: 2000, Gradient: 55},
		{Origin: centralZoneA, Latency: 300, Gradient: 6},
	}

	centralZoneBIndex  = 0
	centralZoneBSeries = []Entry{
		{Origin: centralZoneB, Latency: 850, Gradient: 25},
		{Origin: centralZoneB, Latency: 1256, Gradient: 76},
	}
)

func get(o common.Origin) (Entry, *messaging.Status) {
	switch o.Region {
	case common.WestRegion:
		switch o.Zone {
		case common.WestZoneA:
			return getEntry(&westZoneAIndex, westZoneASeries), messaging.StatusOK()
		case common.WestZoneB:
			return getEntry(&westZoneBIndex, westZoneBSeries), messaging.StatusOK()
		}
	case common.CentralRegion:
		switch o.Zone {
		case common.CentralZoneA:
			return getEntry(&centralZoneAIndex, centralZoneASeries), messaging.StatusOK()
		case common.CentralZoneB:
			return getEntry(&centralZoneBIndex, centralZoneBSeries), messaging.StatusOK()
		}
	}
	return Entry{}, messaging.StatusNotFound()
}

func getEntry(index *int, series []Entry) Entry {
	if *index >= len(series) {
		*index = 0
	}
	e := series[*index]
	*index++
	return e

}
