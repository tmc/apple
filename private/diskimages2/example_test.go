//go:build darwin

package diskimages2_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/diskimages2"
)

func ExampleGetDIAttachedDeviceInfoClass() {
	cls := diskimages2.GetDIAttachedDeviceInfoClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: DIAttachedDeviceInfo
}

func ExampleGetDIConvertParamsClass() {
	cls := diskimages2.GetDIConvertParamsClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: DIConvertParams
}
