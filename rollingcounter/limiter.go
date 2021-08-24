package rollingcounter

import (
	"time"
)

type Limiter struct {
	quota int
	*RollingCounter
}

func NewLimiter(quota, size int, interval time.Duration) *Limiter {
	return &Limiter{
		quota:          quota,
		RollingCounter: NewRollingCounter(size, interval),
	}
}

func (l *Limiter) Allow() bool {
	count := 0
	l.Reduce(func(b int) {
		count += b
	})

	if count <= l.quota {
		l.Add(1)
	}

	return count <= l.quota
}
