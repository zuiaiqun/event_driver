package event_driver

import (
	"sync"
)

type NewFunc func([]interface{})
type NewFuncs []NewFunc

type EventHandler2 struct {
	mutex     sync.Mutex
	eventData map[int32]NewFuncs
}

func NewEventHandler2() *EventHandler2 {
	return &EventHandler2{
		eventData: make(map[int32]NewFuncs),
	}
}

func (this *EventHandler2) AddEvent(eventId int32, f NewFunc) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	_, ok := this.eventData[eventId]
	if !ok {
		this.eventData[eventId] = []NewFunc{f}
	} else {
		this.eventData[eventId] = append(this.eventData[eventId], f)
	}
}

func (this *EventHandler2) RemoveEvent(eventId int32) {
	this.mutex.Lock()
	defer this.mutex.Lock()
	if _, ok := this.eventData[eventId]; ok {
		delete(this.eventData, eventId)
	}
}

func (this *EventHandler2) TriggerEvent(eventId int32, params ...interface{}) {
	this.mutex.Lock()
	funcs, ok := this.eventData[eventId]
	this.mutex.Unlock()
	if !ok {
		return
	}
	for _, function := range funcs {
		function(params)
	}
}
