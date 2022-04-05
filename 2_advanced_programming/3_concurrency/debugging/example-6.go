package main

import (
	"fmt"
	"sync"
)

type coordinator struct {
	lock   sync.RWMutex
	leader string
}

func newCoordinator(leader string) *coordinator {
	return &coordinator{
		lock:   sync.RWMutex{},
		leader: leader,
	}
}

// Assumes c.lock is already acquired
func (c *coordinator) logState() {
	fmt.Printf("leader = %q\n", c.leader)
}

func (c *coordinator) logStateWithLock() {
	c.lock.RLock()
	defer c.lock.RUnlock()

	c.logState()
}

func (c *coordinator) setLeader(leader string, shouldLog bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.leader = leader

	if shouldLog {
		c.logState()
	}
}

func main() {
	// A deadlock occurred when setLeader called logState because
	// setLeader was holding a lock and logState also tried to acquire a lock
	c := newCoordinator("us-east")
	c.logStateWithLock()
	c.setLeader("us-west", true)
}
