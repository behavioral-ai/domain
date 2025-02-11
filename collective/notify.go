package collective

import "github.com/behavioral-ai/core/aspect"

var (
	notifications map[Urn][]Notify
)

func addNotification(thing Urn, fn Notify) *aspect.Status {
	if thing == "" || fn == nil {
		return aspect.StatusBadRequest()
	}
	list, ok := notifications[thing]
	if !ok {
		notifications[thing] = []Notify{fn}
	} else {
		list = append(list, fn)
		notifications[thing] = list
	}
	return aspect.StatusOK()
}

func removeNotification(thing Urn) *aspect.Status {
	if thing == "" {
		return aspect.StatusBadRequest()
	}
	_, ok := notifications[thing]
	if !ok {
		return aspect.StatusOK()
	}
	notifications[thing] = nil
	return aspect.StatusOK()
}

func sendNotification(thing, event Urn) *aspect.Status {
	if thing == "" || event == "" {
		return aspect.StatusBadRequest()
	}
	list, ok := notifications[thing]
	if !ok {
		return aspect.StatusBadRequest()
	}
	for _, fn := range list {
		fn(thing, event)
	}
	return aspect.StatusOK()
}
