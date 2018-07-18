package sweetmemories

import "time"

type (
	Memory struct {
		Data         interface{}
		TimeToForget time.Duration
	}

	internalTest struct {
		a int
		b string
	}
)
