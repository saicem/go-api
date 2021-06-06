package api

import (
	"github.com/saicem/api/models/api/code"
)

type Response struct {
	Status  code.ApiCode `json:"status"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}
