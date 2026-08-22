//go:build darwin

package espresso_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/espresso"
)

func ExampleGetEspressoContextClass() {
	cls := espresso.GetEspressoContextClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: EspressoContext
}

func ExampleGetEspressoNetworkClass() {
	cls := espresso.GetEspressoNetworkClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: EspressoNetwork
}
