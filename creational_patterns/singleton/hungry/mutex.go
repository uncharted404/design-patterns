package hungry

import (
	"design-patterns/creational_patterns/singleton"
	"sync"
)

var (
	mu              sync.Mutex
	instanceByMutex *singleton.Singleton
)

func GetInstanceByMutex() *singleton.Singleton {
	if instanceByMutex == nil {
		mu.Lock()
		defer mu.Unlock()
		if instanceByMutex == nil {
			instanceByMutex = singleton.NewSingleton("mutex")
		}
	}
	return instanceByMutex
}
