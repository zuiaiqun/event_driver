package event_driver

import (
	"reflect"
	"sync"
)

type Func interface{}
type Funcs []Func

type EventHandler struct {
	mutex     sync.Mutex
	eventData map[int32]Funcs
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		eventData: make(map[int32]Funcs),
	}
}

func (this *EventHandler) AddEvent(eventId int32, f Func) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	_, ok := this.eventData[eventId]
	if !ok {
		this.eventData[eventId] = []Func{f}
	} else {
		this.eventData[eventId] = append(this.eventData[eventId], f)
	}
}

func (this *EventHandler) RemoveEvent(eventId int32) {
	this.mutex.Lock()
	defer this.mutex.Lock()
	if _, ok := this.eventData[eventId]; ok {
		delete(this.eventData, eventId)
	}
}

func (this *EventHandler) TriggerEvent(eventId int32, params ...interface{}) {
	this.mutex.Lock()
	funcs, ok := this.eventData[eventId]
	this.mutex.Unlock()
	if !ok {
		return
	}
	values := make([]reflect.Value, len(params))
	for i, value := range params {
		values[i] = reflect.ValueOf(value)
	}
	for _, function := range funcs {
		reflect.ValueOf(function).Call(values)
	}
}
