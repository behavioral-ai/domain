package guidance

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"reflect"
)

const (
	ContentTypeCalendar = "application/calendar"
)

func CalendarTypeErrorStatus(agentId string, t any) *aspect.Status {
	err := errors.New(fmt.Sprintf("error: calendar data change type:%v is invalid for agent:%v", reflect.TypeOf(t), agentId))
	return aspect.NewStatusError(aspect.StatusInvalidArgument, err)
}
