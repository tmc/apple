//go:build darwin

package hiservices_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/hiservices"
)

func ExampleGetHIRunLoopSemaphoreClass() {
	cls := hiservices.GetHIRunLoopSemaphoreClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: HIRunLoopSemaphore
}

func ExampleGetHIRunLoopUtilitiesClass() {
	cls := hiservices.GetHIRunLoopUtilitiesClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: HIRunLoopUtilities
}
