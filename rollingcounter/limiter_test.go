package rollingcounter

import (
	"testing"
	"time"
)

func TestLimiterAllow(t *testing.T) {
	limiter := NewLimiter(10, 10, 100*time.Millisecond)

	for i := 0; i < 10; i++ {
		t.Log(limiter.Allow())
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(500 * time.Millisecond)

	for i := 0; i < 10; i++ {
		t.Log(limiter.Allow())
	}
}
