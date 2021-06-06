package models

import "time"

type PushJson struct {
	Uid         string       `json:"uid"`
	DeviceId    string       `json:"device_id"`
	UserObjects []UserObject `json:"user_objects"`
}

type UserObject struct {
	ObjectName string      `json:"object_name"`
	UserEvents []UserEvent `json:"user_events"`
}

type UserEvent struct {
	EventName  string      `json:"event_name"`
	EventInfos []EventInfo `json:"event_infos"`
}

type EventInfo struct {
	ActTime time.Time `json:"act_time"`
	Append  string    `json:"append"`
}
