package rollingcounter

import (
	"sync"
	"time"
)

type RollingCounter struct {
	sync.RWMutex

	// size 桶大小
	size int
	// buckets 桶，一个环形数组，offset % size 算出 index
	buckets []int
	// interval 每个桶的时间间隔
	interval time.Duration
	// offset 上一次 add 时的偏移量
	offset int
	// lastTime 上一次 add 的时间
	lastTime time.Time
}

func NewRollingCounter(size int, interval time.Duration) *RollingCounter {
	return &RollingCounter{
		size:     size,
		buckets:  make([]int, size),
		interval: interval,
		lastTime: time.Now(),
	}
}

func (rw *RollingCounter) Add(v int) {
	rw.Lock()
	defer rw.Unlock()

	rw.updateOffset()
	rw.add(v)
}

func (rw *RollingCounter) Reduce(fn func(b int)) {
	rw.RLock()
	defer rw.RUnlock()

	// 只有在 add 时候会重置过期的桶
	// reduce 时要排除掉过期的桶
	span := rw.span()
	// count 表示没过期桶的数量
	count := rw.size - span
	if count > 0 {
		offset := (rw.offset + span + 1) % rw.size
		rw.reduce(offset, count, fn)
	}
}

// span 返回距离上次 add 过了几个 interval，也就是过期桶的数量
func (rw *RollingCounter) span() int {
	offset := int(time.Since(rw.lastTime) / rw.interval)
	if offset < rw.size {
		return offset
	}
	// offset > size 时表示桶全部过期
	return rw.size
}

func (rw *RollingCounter) updateOffset() {
	span := rw.span()
	// 等于0表示没有过期，不需要重置
	if span <= 0 {
		return
	}

	// offset + 1 是最旧数据的位置
	// 从 offset + 1 开始，重置 span 个桶
	offset := rw.offset
	for i := 0; i < span; i++ {
		rw.reset((offset + i + 1) % rw.size)
	}

	// 计算出当前偏移量，并更新时间
	rw.offset = (offset + span) % rw.size
	rw.lastTime = rw.lastTime.Add(time.Duration(span * int(rw.interval)))
}

func (rw *RollingCounter) add(v int) {
	rw.buckets[rw.offset] += v
}

func (rw *RollingCounter) reset(offset int) {
	rw.buckets[offset] = 0
}

func (rw *RollingCounter) reduce(start, count int, fn func(b int)) {
	for i := 0; i < count; i++ {
		fn(rw.buckets[(start+i)%rw.size])
	}
}
