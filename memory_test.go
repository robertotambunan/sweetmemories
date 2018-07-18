package sweetmemories

import (
	"reflect"
	"testing"
	"time"
)

func TestMemory_GetMemory(t *testing.T) {
	type fields struct {
		Data         interface{}
		TimeToForget time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"integer", fields{
			Data:         1,
			TimeToForget: 10,
		}, 1},
		{"string", fields{
			Data:         "for test",
			TimeToForget: 10,
		}, "for test"},
		{"struct", fields{
			Data: internalTest{
				a: 100,
				b: "an hundred",
			},
			TimeToForget: 10,
		}, internalTest{
			a: 100,
			b: "an hundred",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Memory{
				Data:         tt.fields.Data,
				TimeToForget: tt.fields.TimeToForget,
			}
			if got := m.GetMemory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Memory.GetMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Memory
	}{
		{"init test", &Memory{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemory_IsForget(t *testing.T) {
	type fields struct {
		Data         interface{}
		TimeToForget time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"memorize", fields{
			Data:         "abc",
			TimeToForget: 100,
		}, false},
		{"forget", fields{
			Data:         nil,
			TimeToForget: 100,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Memory{
				Data:         tt.fields.Data,
				TimeToForget: tt.fields.TimeToForget,
			}
			if got := m.IsForget(); got != tt.want {
				t.Errorf("Memory.IsForget() = %v, want %v", got, tt.want)
			}
		})
	}
}
