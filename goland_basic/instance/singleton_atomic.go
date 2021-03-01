package goland_basic

import (
	"sync"
	"sync/atomic"
)

type Singleton struct {
}

var (
	instance   *Singleton
	initalized uint32
	locker     sync.Mutex
)

func GetInstance() *Singleton {
	if atomic.LoadUint32(&initalized) == 1 {
		return instance
	}

	locker.Lock()
	defer locker.Unlock()

	if initalized == 0 {
		instance = &Singleton{}
		atomic.StoreUint32(&initalized, 1)
	}

	return instance

}
