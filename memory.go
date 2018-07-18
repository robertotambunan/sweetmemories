package sweetmemories

import (
	"time"
)

func (m *Memory) Forget() {
	m.Data = nil
	return
}

func (m *Memory) Save(data interface{}, ttl time.Duration) {
	m.Data = data
	m.TimeToForget = ttl
	go m.forgetByTime()
	return
}

func (m *Memory) forgetByTime() {
	time.Sleep(m.TimeToForget)
	m.Forget()
	return
}

func (m *Memory) GetMemory() interface{} {
	return m.Data
}

func (m *Memory) IsForget() bool {
	if m.Data == nil {
		return true
	}
	return false
}

func New() *Memory {
	return &Memory{}
}
