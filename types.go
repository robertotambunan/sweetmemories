package sweetmemories

import "time"

type Memory struct {
	Data         interface{}
	TimeToForget time.Duration
}
