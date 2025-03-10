package metrics1

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"reflect"
)

const (
	ContentTypeCalendar = "application/calendar"
	ProfileName         = "resiliency:type/domain/metrics/profile"
)

func CalendarTypeErrorStatus(agentId string, t any) *aspect.Status {
	err := errors.New(fmt.Sprintf("error: calendar data change type:%v is invalid for agent:%v", reflect.TypeOf(t), agentId))
	return aspect.NewStatusError(aspect.StatusInvalidArgument, err)
}
