package guidance

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	ContentTypeCalendar = "application/calendar"
)

func CalendarTypeErrorStatus(agentId string, t any) error {
	err := errors.New(fmt.Sprintf("error: calendar data change type:%v is invalid for agent:%v", reflect.TypeOf(t), agentId))
	return err //aspect.NewStatusError(aspect.StatusInvalidArgument, err)
}
