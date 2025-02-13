package guidance

import (
	"github.com/behavioral-ai/core/messaging"
)

type ProcessingCalendar struct {
	week [7][24]string
}

func NewProcessingCalendar() *ProcessingCalendar {
	c := new(ProcessingCalendar)
	return c
}

func GetCalendar(h messaging.Notifier, agentId string, msg *messaging.Message) *ProcessingCalendar {
	if msg.ContentType() != ContentTypeCalendar {
		return nil
	}
	if p, ok := msg.Body.(*ProcessingCalendar); ok {
		return p
	}
	h.Notify(CalendarTypeErrorStatus(agentId, msg.Body))
	return nil
}
