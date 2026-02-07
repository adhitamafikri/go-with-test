package go_sync

import "sync"

// type Counter struct {
// 	num int
// }
type Counter struct {
	mu  sync.Mutex
	num int
}

// func (c *Counter) Inc() {
// 	c.num++
// }

// This means that any goroutine calling the `Inc` will acquire the lock on the `Counter`
// All other goroutines have to wait for the `Counter` to be unlocked before getting access
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num++
}

func (c *Counter) Value() int {
	return c.num
}

func NewCounter() *Counter {
	return &Counter{}
}
