package lazy

import (
	"design-patterns/creational_patterns/singleton"
)

var instanceByVar = singleton.NewSingleton("var")

func GetInstanceByVar() *singleton.Singleton {
	return instanceByVar
}
