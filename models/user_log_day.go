package models

import "time"

type UserLogByDay struct {
	UID    string    `json:"uid"`
	ActDay time.Time `json:"act_day"`
}
