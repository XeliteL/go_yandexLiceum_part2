package step4

import "sync"

type ConcurrentQueue struct {
	queue []interface{}
	mutex sync.Mutex
}

type Queue interface {
	Enqueue(element interface{})
	Dequeue() interface{}
}

func (c *ConcurrentQueue) Enqueue(element interface{}) {
	c.mutex.Lock()
	c.queue = append(c.queue, element)
	c.mutex.Unlock()
}

func (c *ConcurrentQueue) Dequeue() interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	firstInterface := c.queue[0]
	c.queue = c.queue[1:]
	return firstInterface
}
