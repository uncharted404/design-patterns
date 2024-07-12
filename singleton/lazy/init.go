package lazy

import "design-patterns/singleton"

var instanceByInit *singleton.Singleton

func init() {
	instanceByInit = singleton.NewSingleton()
}

func GetInstanceByInit() *singleton.Singleton {
	return instanceByInit
}
