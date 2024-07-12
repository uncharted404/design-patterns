package hungry

import (
	"design-patterns/singleton"
	"sync"
)

var (
	once           sync.Once
	instanceByOnce *singleton.Singleton
)

func GetInstanceByOnce() *singleton.Singleton {
	// 底层由原子操作+锁实现
	// tips：如果Do方法只能执行一次，如果执行的函数出错，会导致无法重新实例化
	once.Do(func() {
		instanceByOnce = singleton.NewSingleton()
	})
	return instanceByOnce
}
