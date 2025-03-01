package metrics1

import "time"

const (
	Low    = "low"
	Medium = "med"
	High   = "high"
)

type TrafficProfile struct {
	Week [7][24]string
}

func NewTrafficProfile() *TrafficProfile {
	c := new(TrafficProfile)
	return c
}

func (t *TrafficProfile) Now() string {
	ts := time.Now().UTC()
	return t.Week[ts.Day()][ts.Hour()]
}

//func dayHour(t *TrafficPofile)

/*
func GetCalendar(h messaging.Notifier, agentId string, msg *messaging.Message) *ProcessingCalendar {
	if !msg.IsContentType(ContentTypeCalendar) {
		return nil
	}
	if p, ok := msg.Body.(*ProcessingCalendar); ok {
		return p
	}
	h.Notify(CalendarTypeErrorStatus(agentId, msg.Body))
	return nil
}


*/
