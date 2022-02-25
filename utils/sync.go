package utils

import "sync"

type Mutex struct {
	mutex          sync.RWMutex
	count          int64
	MaxInteraction int64
}

func (u *Mutex) Wait() {
	for u.Status() >= u.MaxInteraction {
	}
}

func (u *Mutex) Status() int64 {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	return u.count
}

func (u *Mutex) WaitDone() {
	for u.Status() != 0 {
	}
}

func (u *Mutex) Dec() {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.count--
}

func (u *Mutex) Add(i int) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.count += int64(i)
}

func (u *Mutex) Done() {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.count = 0
}
