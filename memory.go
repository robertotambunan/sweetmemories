package sweetmemories

import (
	"encoding/json"
	"time"
)

func (m *Memory) Forget() {
	m.Data = nil
	return
}

func (m *Memory) Remember(sweet interface{}, timeToForget time.Duration) {
	m.Data = sweet
	m.TimeToForget = timeToForget
	go m.forgetByTime(timeToForget)
	return
}

func (m *Memory) forgetByTime(timeToForget time.Duration) {
	time.Sleep(timeToForget)
	m.Forget()
	return
}

func (m *Memory) GetMemory() (result []byte, err error) {
	result, err = json.Marshal(m.Data)
	return
}

func (m *Memory) IsForget() bool {
	if m.Data == nil {
		return true
	}
	return false
}
