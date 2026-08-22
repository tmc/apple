//go:build darwin

package coreml_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/coreml"
)

func ExampleGetMLModelConfigurationClass() {
	cls := coreml.GetMLModelConfigurationClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: MLModelConfiguration
}

func ExampleGetCoreMLPlatformInfoClass() {
	cls := coreml.GetCoreMLPlatformInfoClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	// Output:
	// class name: CoreML.PlatformInfo
}
