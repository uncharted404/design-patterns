package hungry

import (
	"design-patterns/creational_patterns/singleton"
	"sync/atomic"
	"unsafe"
)

// var instance atomic.Pointer[]
var instanceByCas unsafe.Pointer

func GetInstanceByCas() *singleton.Singleton {
	if instance := atomic.LoadPointer(&instanceByCas); instance != unsafe.Pointer(nil) {
		return (*singleton.Singleton)(instance)
	}

	newSingleton := singleton.NewSingleton("cas")
	if swapped := atomic.CompareAndSwapPointer(&instanceByCas, unsafe.Pointer(nil), unsafe.Pointer(newSingleton)); swapped {
		return newSingleton
	}

	return (*singleton.Singleton)(atomic.LoadPointer(&instanceByCas))
}
