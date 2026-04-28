package counter

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increase() {
	c.value++
}

func (c *Counter) GetValue() int {
	return c.value
}

var counter *Counter
var once sync.Once

func GetInstance() *Counter {
	once.Do(func() {
		counter = &Counter{value: 0}
	})
	return counter
}
