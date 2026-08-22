//go:build darwin

package appleneuralengine_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

func ExampleGetANEDeviceInfoClass() {
	cls := appleneuralengine.GetANEDeviceInfoClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: _ANEDeviceInfo
}

func ExampleANEDeviceInfoClass_HasANE() {
	cls := appleneuralengine.GetANEDeviceInfoClass()
	_ = cls.HasANE()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("queried class:", name)
	// Output:
	// queried class: _ANEDeviceInfo
}
