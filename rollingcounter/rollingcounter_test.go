package rollingcounter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const duration = time.Millisecond * 50

func TestRollingCounterAdd(t *testing.T) {
	const size = 3
	r := NewRollingCounter(size, duration)
	listBuckets := func() []int {
		var buckets []int
		r.Reduce(func(b int) {
			buckets = append(buckets, b)
		})
		return buckets
	}
	assert.Equal(t, []int{0, 0, 0}, listBuckets())
	r.Add(1)
	assert.Equal(t, []int{0, 0, 1}, listBuckets())
	elapse()
	r.Add(2)
	r.Add(3)
	assert.Equal(t, []int{0, 1, 5}, listBuckets())
	elapse()
	r.Add(4)
	r.Add(5)
	r.Add(6)
	assert.Equal(t, []int{1, 5, 15}, listBuckets())
	elapse()
	r.Add(7)
	assert.Equal(t, []int{5, 15, 7}, listBuckets())
	elapse()
	elapse()
	elapse()
	r.Add(7)
	assert.Equal(t, []int{0, 0, 7}, listBuckets())
}

func TestRollingCounterReduce(t *testing.T) {
	const size = 4
	r := NewRollingCounter(size, duration)

	for x := 0; x < size; x++ {
		for i := 0; i <= x; i++ {
			r.Add(i)
		}
		if x < size-1 {
			elapse()
		}
	}
	var result int
	r.Reduce(func(b int) {
		result += b
	})
	assert.Equal(t, 10, result)
}

func BenchmarkRollingCounterAdd(b *testing.B) {
	const size = 3

	r := NewRollingCounter(size, duration)

	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		r.Add(1)
	}
}

func BenchmarkRollingCounterReduce(b *testing.B) {
	const size = 3

	r := NewRollingCounter(size, duration)

	for i := 0; i <= 100; i++ {
		r.Add(1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result int
		r.Reduce(func(b int) {
			result += b
		})
	}

}

func elapse() {
	time.Sleep(duration)
}
