package models

import (
	"time"
)

type TxContext struct {
	TransactionID string
	ClientIp      string
	ClientChannel string
	ClientApiKey  string
	UserNickname  string
	Provider      string
	Event         Event
	LastStageData Event
	Environment   string
	Origin        string
}

type Event map[string]interface{}

func (e *Event) AddTimeEndpoint(endpoint string, timeStart time.Time) {
	event, ok := (*e)["datos"].(map[string]interface{})
	endpoints, valid := event["duracionEndpoints"].(map[string]interface{})
	if !ok {
		event = make(map[string]interface{})
	}
	if !valid {
		endpoints = make(map[string]interface{})
	}
	endpoints[endpoint] = time.Since(timeStart).Milliseconds()
	event["duracionEndpoints"] = endpoints
	(*e)["datos"] = event
}

func (e *Event) AddTimeTotal(timeStart time.Time) {
	(*e)["duracionTotal"] = time.Since(timeStart).Milliseconds()
}

func (e *Event) Pop() map[string]interface{} {
	aux := make(map[string]interface{})
	for k, v := range *e {
		aux[k] = v
	}
	*e = Event{}
	return aux
}

func (e *Event) Append(data map[string]interface{}) {
	eventData, ok := (*e)["datos"].(map[string]interface{})
	if !ok {
		eventData = make(map[string]interface{})
	}
	for k, v := range data {
		eventData[k] = v
	}
	(*e)["datos"] = eventData
}

func (e *Event) Restart(data map[string]interface{}) {
	eventData := make(map[string]interface{})
	for k, v := range data {
		eventData[k] = v
	}
	(*e)["datos"] = eventData
}
