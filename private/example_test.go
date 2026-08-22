//go:build darwin

package private_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

func Example() {
	// Package private serves as the root container package for Apple private framework wrappers.
	cls := appleneuralengine.GetANEDeviceInfoClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("container class:", name)
	// Output:
	// container class: _ANEDeviceInfo
}
