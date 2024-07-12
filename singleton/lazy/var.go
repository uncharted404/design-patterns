package lazy

import "design-patterns/singleton"

var instanceByVar = singleton.NewSingleton()

func GetInstanceByVar() *singleton.Singleton {
	return instanceByVar
}
