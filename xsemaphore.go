package xsemaphore

import (
	"sync"

	"golang.org/x/sync/semaphore"
)

var semaphoreMap sync.Map

// Set semaphore to map
func Set(key any, sem *semaphore.Weighted) {
	semaphoreMap.Store(key, sem)
}

// Get semaphore from map
func Get(key any, n int64) *semaphore.Weighted {
	sem := loadSemaphore(key)
	if sem != nil {
		return sem
	}

	newSem := semaphore.NewWeighted(n)
	Set(key, newSem)
	return newSem
}

func loadSemaphore(key any) *semaphore.Weighted {
	v, found := semaphoreMap.Load(key)
	if !found {
		return nil
	}

	sem, ok := v.(*semaphore.Weighted)
	if !ok {
		return nil
	}

	return sem
}
