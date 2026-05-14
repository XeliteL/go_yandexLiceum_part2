package step4

import "sync"

type Counter struct {
	value int
	mu    sync.RWMutex
}

type Сount interface {
	Increment()
	GetValue() int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
