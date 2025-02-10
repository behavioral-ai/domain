package guidance

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/core"
	"reflect"
)

const (
	ContentTypeCalendar = "application/calendar"
)

func CalendarTypeErrorStatus(agentId string, t any) *core.Status {
	err := errors.New(fmt.Sprintf("error: calendar data change type:%v is invalid for agent:%v", reflect.TypeOf(t), agentId))
	return core.NewStatusError(core.StatusInvalidArgument, err)
}
