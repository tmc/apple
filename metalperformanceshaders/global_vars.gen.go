// Code generated from Apple documentation. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// MPSFunctionConstant values.
const (

	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunctionConstantNone
	MPSFunctionConstantNone MPSFunctionConstant = -1
)

var (
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRectNoClip
	MPSRectNoClip metal.MTLRegion
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MPSRectNoClip"); err == nil && ptr != 0 {
		MPSRectNoClip = objc.ValueAt[metal.MTLRegion](ptr)
	}

}
