package reconciler

import (
	"errors"
	"sync"
)

// used to avoid having the same object multiple times in the queue
type objectLocks struct {
	m        sync.Mutex
	capacity int
	objects  map[string]bool
}

var alreadyPresent = errors.New("item already present")
var queueOverflow = errors.New("queue is at capacity")

func newObjectLocks(capacity int) objectLocks {
	return objectLocks{
		capacity: capacity,
		objects:  make(map[string]bool, capacity),
	}
}

func (o *objectLocks) Take(id string) error {
	o.m.Lock()
	defer o.m.Unlock()
	if len(o.objects) == o.capacity {
		return queueOverflow
	}
	if _, ok := o.objects[id]; !ok {
		o.objects[id] = true
		return nil
	}
	return alreadyPresent
}

func (o *objectLocks) Free(id string) {
	o.m.Lock()
	defer o.m.Unlock()
	delete(o.objects, id)
}
