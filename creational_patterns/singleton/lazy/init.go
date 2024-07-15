package lazy

import (
	"design-patterns/creational_patterns/singleton"
)

var instanceByInit *singleton.Singleton

func init() {
	instanceByInit = singleton.NewSingleton("init")
}

func GetInstanceByInit() *singleton.Singleton {
	return instanceByInit
}
