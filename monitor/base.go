package base

import (
	"time"
)

const (
	CPU  = "cpu"
	HDD  = "hdd"
	RAM  = "ram"
	NET  = "net"
	TMPR = "temperature"
)

type MonitoringData struct {
	Name      string      `json:"type"`
	CreatedAt time.Time   `json:"created_at"`
	Data      interface{} `json:"data"`
}
