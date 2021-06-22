package iwut

import (
	"time"
)

type UserLog struct {
	Id         uint64    `json:"id" gorm:"primaryKey"`
	Uid        string    `json:"uid"`
	ObjectName string    `json:"object_name"`
	EventName  string    `json:"event_name"`
	ActTime    time.Time `json:"act_time" gorm:"index"`
	AppendInfo string    `json:"append_info"`
}
