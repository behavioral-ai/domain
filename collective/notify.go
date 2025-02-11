package collective

import (
	"github.com/behavioral-ai/core/aspect"
	"sync"
)

var (
	notifications = new(sync.Map)
)

func addNotification(thing Urn, fn Notify) *aspect.Status {
	if thing == "" || fn == nil {
		return aspect.StatusBadRequest()
	}
	value, ok := notifications.Load(thing)
	if !ok {
		notifications.Store(thing, []Notify{fn})
	} else {
		if list, ok1 := value.([]Notify); ok1 {
			list = append(list, fn)
			notifications.Store(thing, list)
		}
	}
	return aspect.StatusOK()
}

func getNotification(thing Urn) []Notify {
	if value, ok := notifications.Load(thing); ok {
		if list, ok1 := value.([]Notify); ok1 {
			return list
		}
	}
	return nil
}

/*
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
*/
func sendNotification(thing, event Urn) *aspect.Status {
	if thing == "" || event == "" {
		return aspect.StatusBadRequest()
	}
	value, ok := notifications.Load(thing)
	if ok {
		if list, ok1 := value.([]Notify); ok1 {
			for _, fn := range list {
				fn(thing, event)
			}
		}
	}
	return aspect.StatusOK()
}
